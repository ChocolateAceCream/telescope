import apiAxios from '../utils/apiAxios'
export const postLogin = (...data) =>
  apiAxios.post('/public/auth/login', ...data)
