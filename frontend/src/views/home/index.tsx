/*
 * @fileName index.jsx
 * @author Di Sheng
 * @date 2024/08/13 21:42:45
 * @description home page
 */
import { useState, useRef } from 'react'
import Icon from '@/components/icon'
import showNotification from '@/components/notification'
import { postLogin } from '@/api/auth'
import './index.css'
import Camera from '@/components/camera'
import Box from '@mui/material/Box'
import { getUploadPresignedUrl, putS3Upload, postClassify } from '@/api/service'

// import { useSSE } from '@/hooks/useSSE'

const Home = () => {
  const [capturedImage, setCapturedImage] = useState<string | null>(null)
  const [className, setClassName] = useState<string | null>(null)
  const [confidence, setConfidence] = useState<number | null>(null)

  const handleImageCapture = (imageUrl: string, blob: Blob) => {
    setCapturedImage(imageUrl)
    S3Uploader(imageUrl, blob)
  }

  const S3Uploader = async (url: string, blob: Blob) => {
    const payload = {
      file_name: 'test.png',
    }
    const { data: res } = await getUploadPresignedUrl(payload)
    console.log('get url result:', res)

    const { data: res2 } = await putS3Upload(res.data.presigned_url, blob)
    const { data: res3 } = await postClassify({
      image_url:
        'https://telescope-develop.s3.us-east-1.amazonaws.com/admin/test.png',
    })
    setClassName(res3.data.class_name)
    setConfidence(res3.data.confidence)
  }

  return (
    <>
      <Box className="flex flex-col items-center space-y-4 p-4">
        <Camera onCapture={handleImageCapture} />
        {/* Display Captured Image */}
        {capturedImage && (
          <Box className="mt-4">
            <img
              src={capturedImage}
              alt="Captured"
              className="w-[375px] md:w-[400px] h-[250px] border rounded-lg object-cover"
            />

            {/* 🔹 Name & Confidence Score */}
            {className && confidence !== null && (
              <p className="mt-2 text-lg font-semibold text-gray-700">
                Detected: {className} ({confidence.toFixed(1)}%)
              </p>
            )}
          </Box>
        )}
      </Box>
    </>
  )
}

export default Home
