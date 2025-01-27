/*
 * @fileName index.tsx
 * @author Di Sheng
 * @date 2024/09/09 14:49:57
 * @description MyForm
 */
import { Box, BoxProps } from '@mui/material'
import React, { FC, FormEvent, ChangeEvent } from 'react'

interface MyFormProps<T> extends BoxProps {
  formData: T // Accepts any form data structure
  setFormData: React.Dispatch<React.SetStateAction<T>> // Function to update form data
  onSubmit: (event: FormEvent<HTMLDivElement>) => void // Optional onSubmit handler
}

const MyForm = <T extends Record<string, any>>({
  formData,
  setFormData,
  onSubmit,
  children,
  ...props
}: MyFormProps<T>) => {
  const handleSubmit = (event: FormEvent<HTMLDivElement>) => {
    event.preventDefault()
    onSubmit(event)
  }
  const handleChange = (event: ChangeEvent<HTMLInputElement>) => {
    const { name, value } = event.target
    setFormData((prevData) => ({
      ...prevData,
      [name]: value,
    }))
  }

  return React.createElement(
    Box,
    {
      sx: { mt: 1 },
      component: 'form',
      onSubmit: handleSubmit,
      onChange: handleChange,
      ...props,
    },
    children
  )
}

export default MyForm
