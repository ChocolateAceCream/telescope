import React, { useEffect, useRef } from 'react'
import {
  TextField,
  Button,
  Card,
  CardContent,
  Typography,
  Checkbox,
  Divider,
  FormControlLabel,
} from '@mui/material'
// import { FcGoogle } from 'react-icons/fc'
import MyForm from '@/components/form'
import { useState } from 'react'
import MyButton from '@/components/button'
import { sha256 } from '@/utils/encryption'
import { postLogin, getUserInfo } from '@/api/auth'
import userStore from '@/store/user'
import { useNavigate } from 'react-router-dom'
import Icon from '@/components/icon'
import { Link } from 'react-router-dom'
import showNotification from '@/components/notification'
import { postSendCode, postRegister } from '@/api/auth'

import ValidatedTextField, {
  ValidatedTextFieldRef,
} from '@/components/textField'
import { validatePassword, validateEmail } from '@/utils/validators'

interface FormData {
  email: string
  password: string
  code: string
  username: string
}

const SignupPage = () => {
  const navigate = useNavigate()
  const updateUser = userStore((state) => state.updateUser)
  const [formData, setFormData] = useState<FormData>({
    email: '',
    password: '',
    code: '',
  } as FormData)

  const [emailCodeCountdown, setEmailCodeCountdown] = useState(60)
  const [isEmailCodeButtonDisabled, setIsEmailCodeButtonDisabled] =
    useState(false)

  const fieldRefs = useRef<{ [key: string]: ValidatedTextFieldRef | null }>({
    email: null,
    password: null,
    passwordConfirm: null,
  })

  const formRef = useRef<HTMLFormElement | null>(null) // ✅ Ref to form

  const onSendCode = async () => {
    showNotification({
      type: 'success',
      message: 'Verification code has been sent to email, please check',
      duration: 3000, // Duration is in milliseconds
    })
    const resp = await postSendCode({
      email: formData.email,
    })
    console.log('resp:', resp)
  }
  const handleSendCode = () => {
    console.log(formRef.current)
    let isValid = formRef.current?.reportValidity()
    isValid =
      Object.values(fieldRefs.current).every(
        (ref) => ref?.validate() ?? false
      ) && isValid

    if (isValid && !isEmailCodeButtonDisabled) {
      onSendCode()
      setIsEmailCodeButtonDisabled(true)

      let countdown = 60
      const timer = setInterval(() => {
        countdown -= 1
        setEmailCodeCountdown(countdown)
        if (countdown === 0) {
          clearInterval(timer)
          setIsEmailCodeButtonDisabled(false)
          setEmailCodeCountdown(60)
        }
      }, 1000)
    }
  }

  const handleSignup = async (e: React.FormEvent<HTMLDivElement>) => {
    console.log('Form data submitted:', formData, sha256(formData.password))
    e.preventDefault()
    let isValid = formRef.current?.reportValidity()
    isValid =
      Object.values(fieldRefs.current).every(
        (ref) => ref?.validate() ?? false
      ) && isValid
    console.log('validated?', isValid)

    const payload = {
      password: sha256(formData.password),
      code: formData.code,
      email: formData.email,
      username: formData.username,
    }
    await postRegister(payload)

    const { data: resp } = await getUserInfo()
    updateUser({
      isAuthed: true,
      ...resp.data,
    })
    navigate('/')
  }

  const passwordConfirmValidator = (value: string) => {
    if (value !== formData.password) {
      return 'Passwords do not match'
    }
    return ''
  }

  return (
    <div className="flex items-center justify-center min-h-screen">
      <Card className="max-w-sm m-4 p-4 shadow-lg rounded-2xl !important ">
        <CardContent>
          <Typography
            variant="h4"
            component="h2"
            className="mb-2 text-center font-bold text-gray-800"
          >
            Register New Account
          </Typography>
          <Typography
            variant="body1"
            className="mb-6 text-center text-gray-500"
          >
            Please enter your details.
          </Typography>
          <MyForm
            ref={formRef}
            formData={formData}
            setFormData={setFormData}
            onSubmit={handleSignup}
          >
            <ValidatedTextField
              ref={(el) => (fieldRefs.current.email = el)}
              validator={validateEmail}
              margin="normal"
              required
              fullWidth
              id="email"
              label="Email"
              name="email"
              className="mb-6"
            />
            <TextField
              margin="normal"
              required
              fullWidth
              id="username"
              label="Username"
              name="username"
              className="mb-6"
            />
            <ValidatedTextField
              ref={(el) => (fieldRefs.current.password = el)}
              validator={validatePassword}
              margin="normal"
              required
              fullWidth
              name="password"
              label="Password"
              type="password"
              id="password"
              defaultHelperText="Password must include letters, numbers, and special characters, 8-20 characters long"
              autoComplete="current-password"
              className="mb-6"
            />

            <ValidatedTextField
              ref={(el) => (fieldRefs.current.passwordConfirm = el)}
              validator={passwordConfirmValidator}
              margin="normal"
              required
              fullWidth
              name="passwordConfirm"
              label="Confirm Password"
              type="password"
              id="passwordConfirm"
              autoComplete="current-password"
              className="mb-6"
            />

            <div className="flex w-full space-x-4 items-center mb-6">
              <TextField
                className="w-[60%]"
                label="Email Verification Code"
                variant="outlined"
                id="code"
                name="code"
              />
              <Button
                className="w-[40%] "
                variant="contained"
                disabled={isEmailCodeButtonDisabled}
                onClick={handleSendCode}
              >
                {isEmailCodeButtonDisabled
                  ? `${emailCodeCountdown} `
                  : 'Get Code'}
              </Button>
            </div>

            <MyButton
              type="submit"
              fullWidth
              variant="contained"
              color="primary"
              className="my-6"
            >
              Register
            </MyButton>
          </MyForm>
          <Divider className="p-2 px-2 text-gray-500">OR</Divider>
          <MyButton
            type="submit"
            fullWidth
            variant="contained"
            color="secondary"
            className="!bg-blue-500 !hover:bg-green-700 !text-white"
          >
            <Link className="w-full" to="/login">
              Sign in to an existing account
            </Link>
          </MyButton>
        </CardContent>
      </Card>
    </div>
  )
}

export default SignupPage
