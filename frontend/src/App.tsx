/*
 * @fileName App.tsx
 * @author Di Sheng
 * @date 2025/02/20 17:00:17
 * @description Description UI Logic (Routing, Zustand session check)
 */

import React, { useEffect } from 'react'
import { RouterProvider } from 'react-router-dom'
import router from '@/router'
import userStore from '@/store/user' // Zustand auth store

export default function App() {
  const { renewSession } = userStore()
  const isAuthed = userStore.getState().user.isAuthed
  console.log('isAuthed:', isAuthed)
  useEffect(() => {
    if (isAuthed) {
      console.log('---------isAuthed--------', isAuthed)
      renewSession()
    }
  }, [isAuthed])
  return <RouterProvider router={router} />
}
