import React, { useEffect } from 'react'
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
import { postLogin, getUserInfo } from '@/api/auth'
import userStore from '@/store/user'
import { useNavigate } from 'react-router-dom'
import Icon from '@/components/icon'

interface FormData {
  email: string
  password: string
}

const LoginPage = () => {
  const navigate = useNavigate()
  const updateUser = userStore((state) => state.updateUser)
  const [formData, setFormData] = useState<FormData>({
    email: '',
    password: '',
  } as FormData)

  useEffect(() => {
    const handleOAuthCallback = async (event: MessageEvent) => {
      console.log('----event--', event)
      if (event.origin !== window.location.origin) {
        console.warn('untrusted origin', event.origin)
        return
      }

      const { status } = event.data
      if (status === 'success') {
        console.log('login success')
        const resp = await getUserInfo()
        console.log('resp:', resp)
        console.log('resp.data.data:', resp.data.data)
        updateUser({
          isAuthed: true,
          ...resp.data.data,
        })
        navigate('/')
      } else {
        console.error('login failed')
      }
    }

    window.addEventListener('message', handleOAuthCallback)

    return () => {
      window.removeEventListener('message', handleOAuthCallback, false)
    }
  }, [])

  const handlePasswordLogin = async (
    event: React.FormEvent<HTMLDivElement>
  ) => {
    console.log('Form data submitted:', formData, sha256(formData.password))
    await postLogin({
      ...formData,
      password: sha256(formData.password),
    })

    const resp = await getUserInfo()
    console.log('resp:', resp)
    console.log('resp.data.data:', resp.data.data)
    updateUser({
      isAuthed: true,
      ...resp.data.data,
    })
    navigate('/')
  }

  const handleGoogleLogin = async () => {
    console.log('Google login')
    const googleClientId = import.meta.env.VITE_APP_GOOGLE_CLIENT_ID
    console.log('googleClientId:', googleClientId)

    const redirectUri = import.meta.env.VITE_APP_GOOGLE_OAUTH_REDIRECT_URI
    const scope = encodeURIComponent('profile email')
    const authUrl = `https://accounts.google.com/o/oauth2/v2/auth?client_id=${googleClientId}&redirect_uri=${redirectUri}&response_type=code&scope=${scope}`
    window.open(authUrl, '_blank', 'width=500,height=600')
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
            onSubmit={handlePasswordLogin}
          >
            <TextField
              margin="normal"
              required
              fullWidth
              id="email"
              label="Email"
              name="email"
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

          <MyButton
            fullWidth
            variant="contained"
            color="primary"
            className="py-2"
            onClick={handleGoogleLogin}
          >
            <Icon name="google" className="w-4 h-4 pr-1" />
            Google Login
          </MyButton>
        </CardContent>
      </Card>
    </div>
  )
}

export default LoginPage
