import { create } from 'zustand'
import { persist } from 'zustand/middleware'

interface User {
  name: string
  email: string
  role: string
  language: string
  isAuthed: boolean
}

interface UserStore {
  user: User
  updateUser: (newUser: Partial<User>) => void
  logout: () => void
}

const userStore = create<UserStore>()(
  persist(
    (set, get) => ({
      user: {
        name: '',
        email: '',
        role: '',
        language: 'cn',
        isAuthed: false,
      },
      updateUser: (newUser: Partial<User>) => {
        set((state) => ({
          user: {
            ...state.user, // Keep old values
            ...newUser, // Override with new values
          },
        }))
      },
      logout: () => {
        // 🔹 Delete the specific `telescope_session` cookie
        console.log('Logging out...')
        document.cookie.split(';').forEach((cookie) => {
          const name = cookie.split('=')[0].trim()
          document.cookie = `${name}=; Path=/; Domain=${window.location.hostname}; Max-Age=0`
        })

        set((state) => ({
          user: {
            name: '',
            email: '',
            role: '',
            language: state.user.language,
            isAuthed: false,
          },
        }))
      },
    }),
    {
      name: 'user-storage',
    }
  )
)

export default userStore
