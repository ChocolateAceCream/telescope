import React, { useState } from 'react'
import { Avatar, Popover, MenuItem, Box, Typography } from '@mui/material'
import userStore from '@/store/user' // Zustand auth store
import { useNavigate } from 'react-router-dom'
import { postLogout } from '@/api/auth'

const AvatarPopover: React.FC = () => {
  const [anchorEl, setAnchorEl] = useState<HTMLElement | null>(null)
  const logout = userStore((state) => state.logout)
  const navigate = useNavigate()

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
  const avatarUrl = userStore((state) => state.getAvatar)

  return (
    <Box>
      {/* ✅ Avatar that triggers the popover */}
      <Avatar
        alt="User Avatar"
        src={avatarUrl()} // Example image
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
