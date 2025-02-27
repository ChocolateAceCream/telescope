import React, { useRef, useState, useEffect } from 'react'
import Box from '@mui/material/Box'
import MyButton from '@/components/button'
import { on } from 'events'
import Icon from '@/components/icon'
import { styled } from '@mui/material/styles'
interface CameraProps {
  onCapture: (imageUrl: string, blob: Blob) => void
  isOpen?: boolean
}

const CameraCapture: React.FC<CameraProps> = ({
  onCapture,
  isOpen = false,
}) => {
  const videoRef = useRef<HTMLVideoElement>(null)
  const canvasRef = useRef<HTMLCanvasElement>(null)
  const [isCameraOn, setIsCameraOn] = useState(isOpen)
  const [stream, setStream] = useState<MediaStream | null>(null)
  const [facingMode, setFacingMode] = useState<'user' | 'environment'>('user')

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

  const handleUpload = (event: React.ChangeEvent<HTMLInputElement>) => {
    console.log('handleUpload', event)
    const files = event.target.files
    if (!files || files.length === 0) {
      console.warn('No file selected.')
      return
    }

    const file = files[0] // Only handle the first file

    // Use FileReader to read the file as a data URL
    const reader = new FileReader()
    reader.onload = () => {
      const imageDataUrl = reader.result as string // Get the base64 image URL

      // Convert the File into a Blob
      const blob = new Blob([file], { type: file.type })

      console.log('Blob:', blob)
      console.log('ImageDataUrl:', imageDataUrl)

      // Call the same handler as the camera capture function
      onCapture(imageDataUrl, blob)
    }

    reader.readAsDataURL(file) // Read the file as a data URL
  }
  // Start the camera
  const startCamera = async (mode: 'user' | 'environment') => {
    console.log('-------mode------------', mode)
    try {
      const newStream = await navigator.mediaDevices.getUserMedia({
        video: { facingMode: mode },
      })
      if (videoRef.current) {
        videoRef.current.srcObject = newStream
      }
      setStream(newStream)
      setIsCameraOn(true)
    } catch (error) {
      console.error('Error accessing camera:', error)
    }
  }

  // Stop the camera
  const stopCamera = () => {
    if (stream) {
      stream.getTracks().forEach((track) => track.stop())
    }
    if (videoRef.current) {
      videoRef.current.srcObject = null
    }
  }

  // 🔹 Sync internal state with `defaultActive` changes
  useEffect(() => {
    console.log('isOpen:', isOpen)
    setIsCameraOn(isOpen)
  }, [isOpen])

  // 🔹 Sync internal state with `defaultActive` changes
  useEffect(() => {
    if (isCameraOn) {
      stopCamera()
      startCamera(facingMode)
    } else {
      stopCamera()
    }
  }, [isCameraOn])

  const toggleCamera = () => {
    setIsCameraOn((prev) => !prev)
  }

  const flipCamera = () => {
    setFacingMode((prevMode) => (prevMode === 'user' ? 'environment' : 'user'))
  }

  // 🔹 Use `useEffect` to restart the camera when `facingMode` changes
  useEffect(() => {
    if (isCameraOn) {
      stopCamera()
      startCamera(facingMode)
    }
  }, [facingMode])

  const captureImage = () => {
    const video = videoRef.current
    const canvas = canvasRef.current
    if (video && canvas) {
      const context = canvas.getContext('2d')
      if (context) {
        canvas.width = video.videoWidth
        canvas.height = video.videoHeight
        context.drawImage(video, 0, 0, canvas.width, canvas.height)

        const imageDataUrl = canvas.toDataURL('image/png')
        console.log('---------imageDataUrl----------', imageDataUrl)
        canvas.toBlob((blob) => {
          console.log('blob:', typeof blob)
          if (!blob) {
            return
          }
          onCapture(imageDataUrl, blob)
        })
      }
    }
  }

  return (
    <>
      {/* Video Box with Responsive Fixed Size */}
      <Box className="w-full h-[15rem] overflow-hidden border rounded-lg relative">
        <video
          ref={videoRef}
          autoPlay
          playsInline
          className="absolute inset-0 w-full h-full z-10"
          style={{ objectFit: 'cover', background: '#000' }}
        />
      </Box>

      {/* Hidden Canvas */}
      <canvas ref={canvasRef} className="hidden" />

      {/* Buttons */}
      <Box className="grid grid-cols-4 gap-4">
        <MyButton
          variant="contained"
          color={isCameraOn ? 'error' : 'primary'}
          onClick={toggleCamera}
        >
          {isCameraOn ? <Icon name="cameraClosed" /> : <Icon name="camera" />}
        </MyButton>
        <MyButton
          variant="contained"
          color="secondary"
          onClick={captureImage}
          disabled={!isCameraOn}
        >
          <Icon name="shot" />
        </MyButton>
        <MyButton variant="contained" color="secondary" onClick={flipCamera}>
          <Icon name="flip" />
        </MyButton>
        <MyButton variant="contained" color="secondary" component="label">
          <Icon name="upload" />
          <VisuallyHiddenInput type="file" onChange={handleUpload} multiple />
        </MyButton>
      </Box>
    </>
  )
}

export default CameraCapture
