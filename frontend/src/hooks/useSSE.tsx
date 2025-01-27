import React, {
  createContext,
  useState,
  useEffect,
  useContext,
  useCallback,
  ReactNode,
} from 'react'

// Define the structure of the SSE message
interface SSEMessage {
  type: string
  [key: string]: any // The message can have other dynamic properties
}

// Define the shape of the context
interface SSEContextType {
  subscribe: (
    type: string,
    callback: (message: SSEMessage) => void
  ) => () => void
  isConnected: boolean
  error: Event | null
}

// Initialize the context with a default value (null to be safe)
const SSEContext = createContext<SSEContextType | undefined>(undefined)

// Define the subscribers object
let eventSource: EventSource | null = null
const subscribers: Record<string, Array<(message: SSEMessage) => void>> = {}

// SSE Provider component
interface SSEProviderProps {
  url: string
  children: ReactNode // ReactNode type for children, which can be JSX
}

export const SSEProvider: React.FC<SSEProviderProps> = ({ url, children }) => {
  const [isConnected, setIsConnected] = useState(false)
  const [error, setError] = useState<Event | null>(null)

  useEffect(() => {
    eventSource = new EventSource(url)
    console.log('---url---', url)
    console.log('---eventSource---', eventSource)

    eventSource.onopen = () => {
      setIsConnected(true)
      console.log('---onopen---')
    }

    eventSource.onmessage = (event: MessageEvent) => {
      console.log('-------onmessage---', event)
      const message: SSEMessage = JSON.parse(event.data)
      const { type } = message

      // Notify all subscribers for the specific message type
      if (subscribers[type]) {
        subscribers[type].forEach((callback) => callback(message))
      }
    }

    eventSource.onerror = (err: Event) => {
      setError(err)
      setIsConnected(false)
      eventSource?.close() // Optionally close on error
    }

    return () => {
      eventSource?.close() // Clean up the connection when unmounted
    }
  }, [url])

  // Function to subscribe to a specific message type
  const subscribe = useCallback(
    (type: string, callback: (message: SSEMessage) => void) => {
      if (!subscribers[type]) {
        subscribers[type] = []
      }
      subscribers[type].push(callback)

      return () => {
        subscribers[type] = subscribers[type].filter((cb) => cb !== callback)
      }
    },
    []
  )

  return (
    <SSEContext.Provider value={{ subscribe, isConnected, error }}>
      {children} {/* Ensure children are passed and rendered */}
    </SSEContext.Provider>
  )
}

// Custom hook for subscribing to SSE messages
export const useSSE = (
  type: string,
  callback: (message: SSEMessage) => void
) => {
  const context = useContext(SSEContext)

  if (!context) {
    throw new Error('useSSE must be used within an SSEProvider')
  }

  const { subscribe } = context

  useEffect(() => {
    const unsubscribe = subscribe(type, callback)

    return () => {
      unsubscribe() // Clean up subscription on unmount
    }
  }, [type, callback, subscribe])
}
