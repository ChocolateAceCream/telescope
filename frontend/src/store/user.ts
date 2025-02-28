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
  picture?: string
}

interface UserStore {
  user: User
  updateUser: (newUser: Partial<User>) => void
  logout: () => void
  getAvatar: () => string
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
        picture: '',
      },
      getAvatar: () => {
        let avatarUrl = get().user.picture
        if (!avatarUrl) {
          return generateCanvasAvatar(get().user.username || 'Guest')
        }
        return avatarUrl
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

const generateCanvasAvatar = (name: string, size = 100) => {
  const canvas = document.createElement('canvas')
  canvas.width = size
  canvas.height = size
  const ctx = canvas.getContext('2d')
  if (!ctx) throw new Error('Canvas not supported')

  // Background color
  ctx.fillStyle = '#ffb74d'
  ctx.fillRect(0, 0, size, size)

  // Text styling
  ctx.fillStyle = '#ffffff'
  ctx.font = `${size / 2.5}px Arial`
  ctx.textAlign = 'center'
  ctx.textBaseline = 'middle'

  // Extract initials
  const initials = name
    .split(' ')
    .map((n) => n[0])
    .join('')
    .toUpperCase()

  // Draw initials
  ctx.fillText(initials, size / 2, size / 2)

  // Convert canvas to image
  return canvas.toDataURL()
}

export default userStore
