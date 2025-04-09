// src/components/snackbarAlert.tsx

import { Snackbar, Alert } from '@mui/material'
import { SyntheticEvent } from 'react'

interface SnackbarAlertProps {
  open: boolean
  message: string
  severity?: 'error' | 'warning' | 'info' | 'success'
  duration?: number
  onClose: (event?: SyntheticEvent | Event, reason?: string) => void
}

const SnackbarAlert = ({
  open,
  message,
  severity = 'info',
  duration = 3000,
  onClose,
}: SnackbarAlertProps) => {
  return (
    <Snackbar
      open={open}
      autoHideDuration={duration}
      onClose={onClose}
      anchorOrigin={{ vertical: 'top', horizontal: 'center' }}
    >
      <Alert onClose={onClose} severity={severity} sx={{ width: '100%' }}>
        {message}
      </Alert>
    </Snackbar>
  )
}

export default SnackbarAlert
