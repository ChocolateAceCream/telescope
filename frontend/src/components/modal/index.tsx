import React from 'react'
import {
  Dialog,
  DialogTitle,
  DialogContent,
  DialogActions,
  Button,
} from '@mui/material'

type ModalProps = {
  open: boolean
  onClose: () => void
  onSubmit: () => void
  title: string
  children: React.ReactNode
  submitButtonText?: string
  cancelButtonText?: string
  isLoading?: boolean
}

const Modal: React.FC<ModalProps> = ({
  open,
  onClose,
  onSubmit,
  title,
  children,
  submitButtonText = 'Submit',
  cancelButtonText = 'Cancel',
  isLoading = false,
}) => {
  return (
    <Dialog
      open={open}
      onClose={(_, reason) => {
        // Only allow closing via cancel button
        if (reason === 'backdropClick') return
      }}
      fullWidth
      maxWidth="sm"
      className="rounded-lg"
    >
      {/* Header */}
      <DialogTitle className="flex items-start text-left p-4 border-b border-gray-200">
        <span className="text-lg font-semibold text-gray-800">{title}</span>
      </DialogTitle>

      {/* Body */}
      <DialogContent className="p-4">
        <div className="mt-2">{children}</div>
      </DialogContent>

      {/* Footer */}
      <DialogActions className="p-4 border-t border-gray-200">
        <Button
          onClick={onClose}
          className="px-4 py-2 text-gray-700 hover:bg-gray-100 rounded-md transition-colors"
        >
          {cancelButtonText}
        </Button>
        <Button
          onClick={onSubmit}
          variant="contained"
          loading={isLoading}
          className="px-4 py-2 bg-blue-600 text-white hover:bg-blue-700 rounded-md transition-colors"
        >
          {submitButtonText}
        </Button>
      </DialogActions>
    </Dialog>
  )
}

export default Modal
