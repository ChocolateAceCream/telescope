import React, { useState } from 'react'
import { Avatar, Popover, MenuItem, Box, Typography } from '@mui/material'
import userStore from '@/store/user' // Zustand auth store
import { useNavigate } from 'react-router-dom'
import { postLogout } from '@/api/auth'

const AvatarPopover: React.FC = () => {
  const [anchorEl, setAnchorEl] = useState<HTMLElement | null>(null)
  const logout = userStore((state) => state.logout)
  const navigate = useNavigate()

  // Function moved to the top
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

  // Avatar URL (uses the function to generate avatar if no picture is available)
  const avatarUrl = userStore(
    (state) =>
      state.user.picture || generateCanvasAvatar(state.user.username || 'Guest')
  )

  // Open popover when clicking the avatar
  const handleOpen = (event: React.MouseEvent<HTMLElement>) => {
    setAnchorEl(event.currentTarget)
  }

  // Close popover
  const handleClose = () => {
    setAnchorEl(null)
  }

  // Handle menu actions
  const handleMenuClick = (option: string) => {
    console.log(option) // Replace with actual actions (e.g., navigation, logout)
    handleClose()
  }

  const handleLogout = async () => {
    console.log('Logging out')
    await postLogout()
    handleClose()
    logout()
    navigate('/login') // Redirect to login page after logout
  }

  return (
    <Box>
      {/* ✅ Avatar that triggers the popover */}
      <Avatar
        alt="User Avatar"
        src={avatarUrl} // Example image
        sx={{ cursor: 'pointer' }}
        onClick={handleOpen}
      />

      {/* ✅ Popover with Menu Items */}
      <Popover
        open={Boolean(anchorEl)}
        anchorEl={anchorEl}
        onClose={handleClose}
        anchorOrigin={{
          vertical: 'bottom',
          horizontal: 'right',
        }}
        transformOrigin={{
          vertical: 'top',
          horizontal: 'right',
        }}
      >
        <Box p={1} sx={{ minWidth: 150 }}>
          <MenuItem onClick={() => handleMenuClick('Profile')}>
            Profile
          </MenuItem>
          <MenuItem onClick={() => handleMenuClick('Settings')}>
            Settings
          </MenuItem>
          <MenuItem onClick={() => handleLogout()}>Logout</MenuItem>
        </Box>
      </Popover>
    </Box>
  )
}

export default AvatarPopover
