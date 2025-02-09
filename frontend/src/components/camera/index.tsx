import React, { useRef, useState, useEffect } from 'react'
import Box from '@mui/material/Box'
import MyButton from '@/components/button'
import { on } from 'events'

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

  // Start the camera
  const startCamera = async () => {
    try {
      const newStream = await navigator.mediaDevices.getUserMedia({
        video: true,
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
    setIsCameraOn(false)
  }

  // 🔹 Sync internal state with `defaultActive` changes
  useEffect(() => {
    if (isOpen) {
      startCamera()
    } else {
      stopCamera()
    }
  }, [isOpen])

  const toggleCamera = () => {
    if (isCameraOn) {
      stopCamera()
    } else {
      startCamera()
    }
  }

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
          className="absolute inset-0 w-full h-full"
          style={{ objectFit: 'cover', background: '#000' }}
        />
      </Box>

      {/* Hidden Canvas */}
      <canvas ref={canvasRef} className="hidden" />

      {/* Buttons */}
      <Box className="flex space-x-4">
        <MyButton
          variant="contained"
          color={isCameraOn ? 'error' : 'primary'}
          onClick={toggleCamera}
        >
          {isCameraOn ? 'Close Camera' : 'Open Camera'}
        </MyButton>
        <MyButton
          variant="contained"
          color="secondary"
          onClick={captureImage}
          disabled={!isCameraOn}
        >
          Capture
        </MyButton>
      </Box>
    </>
  )
}

export default CameraCapture
