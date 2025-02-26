import { create } from 'zustand'
import { persist } from 'zustand/middleware'
import { postRenewSession } from '@/api/auth'
import { decodeBase64 } from '@/utils/decoder'
interface User {
  username: string
  email: string
  role: string
  info: string
  language: string
  isAuthed: boolean
}

interface UserStore {
  user: User
  updateUser: (newUser: Partial<User>) => void
  logout: () => void
  renewSession: () => void
}

const userStore = create<UserStore>()(
  persist(
    (set, get) => ({
      user: {
        username: '',
        email: '',
        role: '',
        info: '',
        language: 'cn',
        isAuthed: false,
      },
      updateUser: (newUser: Partial<User>) => {
        const info = decodeBase64(newUser.info)
        set((state) => ({
          user: {
            ...state.user, // Keep old values
            ...newUser, // Override with new values
            ...info,
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
            username: '',
            email: '',
            role: '',
            info: '',
            language: state.user.language,
            isAuthed: false,
          },
        }))
      },
      renewSession: async () => {
        await postRenewSession()
      },
    }),
    {
      name: 'user-storage',
    }
  )
)

export default userStore
