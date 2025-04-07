import apiAxios from '../utils/apiAxios'

export const getUploadPresignedUrl = (...args) => apiAxios.post('/aws/generate-presigned-url', ...args)
export const putS3Upload = (url, file) => apiAxios.put(url, file, {
  headers: {
    'Content-Type': "image/png", // Set the file's MIME type
  },
  baseURL: '', // Ensure baseURL isn't applied at all
})
export const postClassify = (...args) => apiAxios.post('/aws/classify', ...args)

// get files from s3
export const getDownloadPresignedUrl = (filename) => apiAxios.get('/aws/download?file_name=' + filename)
