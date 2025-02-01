/*
 * @fileName index.jsx
 * @author Di Sheng
 * @date 2024/08/13 21:42:45
 * @description home page
 */
import { useState } from 'react'
import Icon from '@/components/icon'
import showNotification from '@/components/notification'
import { postLogin } from '@/api/auth'
import './index.css'
// import { useSSE } from '@/hooks/useSSE'

// Define the structure of the SSE message
interface SSEMessage {
  type: string
  [key: string]: any // The message can have other dynamic properties
}

const Home = () => {
  const handleClick = () => {
    showNotification({
      type: 'success',
      message: 'This is a success notification!',
      duration: 3000, // Duration is in milliseconds
    })
  }
  const handleLogin = async () => {
    const res = await postLogin({ username: 'admin', password: 'admin' })
    console.log(res)
  }

  return (
    <>
      <div>
        <a href="https://vitejs.dev" target="_blank">
          <Icon name="vite" className="logo" />
        </a>
        <a href="https://react.dev" target="_blank">
          <Icon name="react" className="logo" />
        </a>
      </div>
      <h1>Vite + React</h1>
      <div className="card">
        <button onClick={handleLogin}>login</button>
        <p>
          Edit <code>src/App.jsx</code> and save to test HMR
        </p>
      </div>
      <h1 className="font-sans text-shadow-lg ">
        This text uses the extended{' '}
      </h1>
      <div className="bg-red">asdfsadf</div>
      <h1 className="font-inter text-shadow-xl ">
        This text uses the extended{' '}
      </h1>
      <h1 className="font-chewy hover:text-shadow-xl">
        This text uses the extended{' '}
      </h1>
      <button onClick={handleClick}>Show Success Notification</button>
    </>
  )
}

export default Home
