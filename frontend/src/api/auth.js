import apiAxios from '../utils/apiAxios'
export const postLogin = (...data) =>
  apiAxios.post('/public/auth/login', ...data)
export const postRefreshToken = () => apiAxios.post('/public/auth/refresh-token')

export const postRenewSession = () => apiAxios.post('/public/auth/renew-session')

export const getUserInfo = () => apiAxios.get('/user/info')
export const postLogout = () => apiAxios.post('/user/logout')

export const postSendCode = (...data) => apiAxios.post('/public/auth/send-code', ...data)
export const postRegister = (...data) => apiAxios.post('/public/auth/register', ...data)