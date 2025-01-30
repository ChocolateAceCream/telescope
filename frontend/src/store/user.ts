import { create } from 'zustand'
import { persist } from 'zustand/middleware'

interface User {
  name: string
  email: string
  role: string
  language: string
}

interface UserStore {
  user: User
  updateUser: (newUser: Partial<User>) => void
}

const userStore = create<UserStore>()(
  persist(
    (set, get) => ({
      user: {
        name: '',
        email: '',
        role: '',
        language: 'cn',
      },
      updateUser: (newUser: Partial<User>) =>
        set((state) => ({
          user: {
            ...state.user, // Keep old values
            ...newUser, // Override with new values
          },
        })),
    }),
    {
      name: 'user-storage',
    }
  )
)

export default userStore
