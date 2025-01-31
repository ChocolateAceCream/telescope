import React from 'react'
import {
  TextField,
  Button,
  Card,
  CardContent,
  Typography,
  Checkbox,
  FormControlLabel,
} from '@mui/material'
// import { FcGoogle } from 'react-icons/fc'
import MyForm from '@/components/form'
import { useState } from 'react'
import MyButton from '@/components/button'
import { sha256 } from '@/utils/encryption'
import { postLogin } from '@/api/auth'
import userStore from '@/store/user'
import { useNavigate } from 'react-router-dom'
interface FormData {
  username: string
  password: string
}

const LoginPage = () => {
  const navigate = useNavigate()
  const updateUser = userStore((state) => state.updateUser)
  const [formData, setFormData] = useState<FormData>({
    username: '',
    password: '',
  } as FormData)

  const handleSubmit = async (event: React.FormEvent<HTMLDivElement>) => {
    console.log('Form data submitted:', formData, sha256(formData.password))
    const { data: userInfo } = await postLogin({
      ...formData,
      password: sha256(formData.password),
    })
    console.log('Login result:', userInfo)

    updateUser({
      isAuthed: true,
      ...userInfo,
    })

    navigate('/')
  }
  return (
    <div className="flex items-center justify-center min-h-screen">
      <Card className="max-w-sm p-8 shadow-lg rounded-3xl">
        <CardContent>
          <Typography
            variant="h4"
            component="h2"
            className="mb-2 text-center font-bold text-gray-800"
          >
            WELCOME BACK
          </Typography>
          <Typography
            variant="body1"
            className="mb-6 text-center text-gray-500"
          >
            Welcome back! Please enter your details.
          </Typography>
          <MyForm
            formData={formData}
            setFormData={setFormData}
            onSubmit={handleSubmit}
          >
            <TextField
              margin="normal"
              required
              fullWidth
              id="username"
              label="Username"
              name="username"
              autoFocus
              className="mb-4"
            />
            <TextField
              margin="normal"
              required
              fullWidth
              name="password"
              label="Password"
              type="password"
              id="password"
              autoComplete="current-password"
              className="mb-6"
            />
            <MyButton
              type="submit"
              fullWidth
              variant="contained"
              color="primary"
              className="py-2"
            >
              Sign In
            </MyButton>
          </MyForm>
          <Typography
            variant="body2"
            className="mt-6 text-center text-gray-600"
          >
            Don&#39;t have an account?{' '}
            <a href="#" className="text-red-500 hover:underline">
              Sign up for free!
            </a>
          </Typography>
        </CardContent>
      </Card>
    </div>
  )
}

export default LoginPage
