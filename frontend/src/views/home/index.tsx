/*
 * @fileName index.jsx
 * @author Di Sheng
 * @date 2024/08/13 21:42:45
 * @description home page
 */
import { useState, useRef } from 'react'
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
  const [isCameraActive, setIsCameraActive] = useState(false)

  const handleImageCapture = (imageUrl: string, blob: Blob) => {
    setCapturedImage(imageUrl)
    setIsCameraActive(false)
    S3Uploader(imageUrl, blob)
  }

  const handleRetakePhoto = () => {
    setCapturedImage(null) // ✅ Reset image
    setIsCameraActive(true) // ✅ Reopen camera
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
      <Box className="flex flex-col items-center space-y-4 sm:w-[25rem] w-[20rem]  px-2 mx-auto mt-4">
        {/* Show Camera if no image is captured */}
        {!capturedImage ? (
          <Camera onCapture={handleImageCapture} isOpen={isCameraActive} />
        ) : (
          <Box className="w-full">
            {/* Display Captured Image */}
            <img
              src={capturedImage}
              alt="Captured"
              className="w-full h-[15rem] overflow-hidden border rounded-lg relative"
            />

            {/* 🔹 Name & Confidence Score */}
            {className && confidence !== null && (
              <p className="mt-2 text-lg font-semibold text-gray-700">
                Detected: {className} ({confidence.toFixed(1)}%)
              </p>
            )}

            {/* 🔹 Retake Button */}
            <button
              onClick={handleRetakePhoto} // Reset image
              className="mt-3 px-4 py-2 bg-red-500 text-white rounded-lg hover:bg-red-600 transition"
            >
              Retake Photo
            </button>
          </Box>
        )}
      </Box>
    </>
  )
}

export default Home
