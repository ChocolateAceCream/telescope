// src/components/LoginPage.jsx

import { TextField, Box, Typography, Container, Button } from '@mui/material'
import MyButton from '@/components/button'
import MyForm from '@/components/form'
import { useState } from 'react'
import { postLogin } from '@/api/auth'
import {
  postUpload,
  getUploadPresignedUrl,
  putS3Upload,
  getDownloadPresignedUrl,
} from '@/api/service'
import { sha256 } from '@/utils/encryption'
import { styled } from '@mui/material/styles'

interface FormData {
  username: string
  password: string
}

const VisuallyHiddenInput = styled('input')({
  clip: 'rect(0 0 0 0)',
  clipPath: 'inset(50%)',
  height: 1,
  overflow: 'hidden',
  position: 'absolute',
  bottom: 0,
  left: 0,
  whiteSpace: 'nowrap',
  width: 1,
})

const LoginPage = () => {
  const [backgroundImageUrl, setBackgroundImageUrl] = useState<string | null>(
    null
  )
  const [formData, setFormData] = useState<FormData>({
    username: '',
    password: '',
  } as FormData)

  const handleSSE = async () => {
    const es = new EventSource('/backend/api/v1/sse/subscribe')
    // Whenever the connection is established between the server and the client we'll get notified
    es.onopen = (e) => console.log('>>> Connection opened!', e)
    // Made a mistake, or something bad happened on the server? We get notified here
    es.onerror = (e) => console.log('ERROR!', e)
    // This is where we get the messages. The event is an object and we're interested in its `data` property
    es.onmessage = (e) => {
      console.log('>>>', e.data)
    }
  }
  const handleSubmit = async (event: React.FormEvent<HTMLDivElement>) => {
    console.log('Form data submitted:', formData, sha256(formData.password))
    const { data: res } = await postLogin({
      ...formData,
      password: sha256(formData.password),
    })
    console.log('Login result:', res)
  }
  const handleUpload = async () => {
    const { data: res2 } = await postUpload()
    console.log('upload result:', res2)
  }

  const handleDownload = async () => {
    const { data: res } = await getDownloadPresignedUrl('logo.png')
    console.log('handleDownload result:', res)
    setBackgroundImageUrl(res.data.url)
  }

  const handleS3Upload = async (e: any) => {
    console.log(e.target.files[0].name)
    const payload = {
      file_name: e.target.files[0].name,
    }
    const { data: res } = await getUploadPresignedUrl(payload)
    console.log('get url result:', res)

    const { data: res2 } = await putS3Upload(
      res.data.presigned_url,
      e.target.files[0]
    )
    console.log('upload result:', res2)
  }
  return (
    <Box
      maxWidth="sm"
      className="p-8 bg-lite-orange shadow-lg rounded-lg m-auto "
      sx={{
        backgroundImage: backgroundImageUrl
          ? `url(${backgroundImageUrl})`
          : 'none',
      }}
    >
      <Button onClick={handleSSE}>start sse</Button>
      <Button onClick={handleUpload}>upload</Button>
      <Button onClick={handleDownload}>download</Button>
      <Button
        component="label"
        role={undefined}
        variant="contained"
        tabIndex={-1}
      >
        Upload files
        <VisuallyHiddenInput type="file" onChange={handleS3Upload} multiple />
      </Button>
      <Typography variant="h4" className="mb-4 font-bold text-gray-800">
        Login
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
    </Box>
  )
}

export default LoginPage
