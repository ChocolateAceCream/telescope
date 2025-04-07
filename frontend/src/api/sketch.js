import apiAxios from '../utils/apiAxios'

export const postSketchUpload = (formData) => apiAxios.post(
  '/sketch/upload',
  formData,
  {
    headers: {
      'Content-Type': 'multipart/form-data'
    },
    transformRequest: (data) => data,
  }
)