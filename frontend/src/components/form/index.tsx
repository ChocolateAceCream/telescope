/*
 * @fileName index.tsx
 * @author Di Sheng
 * @date 2024/09/09 14:49:57
 * @description MyForm with ref support
 */
import { Box, BoxProps } from '@mui/material'
import React, { FormEvent, ChangeEvent, forwardRef } from 'react'

interface MyFormProps<T> extends BoxProps {
  formData: T // Accepts any form data structure
  setFormData: React.Dispatch<React.SetStateAction<T>> // Function to update form data
  onSubmit: (event: FormEvent<HTMLDivElement>) => void // Optional onSubmit handler
}

// ✅ Use `forwardRef` to expose the form reference
const MyForm = forwardRef<HTMLFormElement, MyFormProps<any>>(
  ({ formData, setFormData, onSubmit, children, ...props }, ref) => {
    const handleSubmit = (event: FormEvent<HTMLDivElement>) => {
      event.preventDefault()
      onSubmit(event)
    }

    const handleChange = (event: ChangeEvent<HTMLInputElement>) => {
      const { name, value } = event.target
      setFormData((prevData: typeof formData) => ({
        ...prevData,
        [name]: value,
      }))
    }

    return (
      <Box
        ref={ref} // ✅ Attach the ref to `<Box component="form">`
        sx={{ mt: 1 }}
        component="form"
        onSubmit={handleSubmit}
        onChange={handleChange}
        {...props}
      >
        {children}
      </Box>
    )
  }
)

export default MyForm
