/*
 * @fileName index.tsx
 * @author Di Sheng
 * @date 2025/03/11 16:55:53
 * @description Description:  TextField component with validator option
 */

import React, { useState, useImperativeHandle, forwardRef } from 'react'
import { TextField, TextFieldProps } from '@mui/material'

export interface ValidatedTextFieldRef {
  validate: () => boolean
}

type ValidatedTextFieldProps = TextFieldProps & {
  validator: (value: string) => string
  defaultHelperText?: string
}

const ValidatedTextField = forwardRef<
  ValidatedTextFieldRef,
  ValidatedTextFieldProps
>(({ validator, defaultHelperText = '', children, ...props }, ref) => {
  const [value, setValue] = useState('')
  const [error, setError] = useState('')
  const [isTouched, setIsTouched] = useState(false)

  const validate = () => {
    setIsTouched(true)
    const errorMessage = validator(value)
    setError(errorMessage)
    return !Boolean(errorMessage)
  }

  useImperativeHandle(ref, () => ({
    validate,
  }))
  const handleBlur = () => {
    validate()
    // setIsTouched(true)
    // const errorMessage = validator(value)
    // setError(errorMessage)
    // console.log('!Boolean(errorMessage):', !Boolean(errorMessage))
    // resultCallback(!Boolean(errorMessage))
  }
  const handleChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setValue(event.target.value)
  }
  return React.createElement(
    TextField,
    {
      sx: { mt: 1 },
      onChange: handleChange,
      error: isTouched && Boolean(error),
      helperText: isTouched ? error : defaultHelperText,
      onBlur: handleBlur,
      ...props,
    },
    children
  )
})

export default ValidatedTextField
