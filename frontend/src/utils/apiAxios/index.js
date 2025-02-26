import axios from 'axios'
import supportCancelToken from './cancelToken'
import { addSignature } from './signature'
import showNotification from '@/components/notification'
import NProgress from 'nprogress'
import qs from 'qs'
import { postRenewSession } from '@/api/auth'

import userStore from '@/store/user'
const SESSION_EXPIRED = 499

const logout = userStore.getState().logout; // Zustand logout function
const apiAxios = new Proxy(
  axios.create({
    // https://cn.vitejs.dev/guide/env-and-mode.html
    baseURL: import.meta.env.VITE_APP_API_BASE_URL || '/',
    timeout: 1000 * 60,
    paramsSerializer: {
      serialize: function(params) {
        return qs.stringify(params, { indices: false })
      },
    },
  }),
  {
    get(target, ...args) {
      return Reflect.get(target, ...args) || Reflect.get(axios, ...args)
    },
  }
)

apiAxios.defaults.meta = {
  retry: 0 /* times*/,
  retryDelay: 100 /* ms*/,
  curRetry: 0 /* times*/,
  // 断开相同请求，判断条件 如果!!cancelToken存在 则计算config.url+cancelToken的值作为唯一key值，key值相同，则断开之前请求
  cancelToken: '',
  withProgressBar: false,
}

apiAxios.defaults.headers.post['Content-Type'] =
  'application/json;charset=UTF-8'
// axios.defaults.headers.post['Content-Type'] = 'application/x-www-form-urlencoded'

supportCancelToken(apiAxios)

let isRefreshing = false
let refreshSubscribers = []

let activeRequest = 0
let logoutFlag = false
// request interceptors
apiAxios.interceptors.request.use(
  (config) => {
    activeRequest++

    const { language } = userStore.getState().user
    config.headers['X-Language'] = language
    if (config['Content-Type']) {
      config.headers['Content-Type'] = config['Content-Type']
    } else {
      config.headers['Content-Type'] = 'application/json;charset=UTF-8'
    }
    if (config.meta?.withProgressBar) {
      NProgress.start()
    }
    if (!Object.prototype.hasOwnProperty.call(config.params || {}, 'sign')) {
      addSignature(config)
    }

    // params encoding for get request
    if (config.method === 'get' && config.params) {
      const params = encodeURIComponent(JSON.stringify(config.params))
      // const serviceName = config.url.split('/')[1]
      config.url = config.url + '?params=' + params
      // console.log('---------config.url-------', config.url)
      config.params = {}
    }

    return config
  },
  (error) => {
    activeRequest--
    return Promise.reject(error)
  }
)

// 响应拦截
apiAxios.interceptors.response.use(
  async (res) => {
    if (res.config.meta?.withProgressBar) {
      NProgress.done()
    }
    // 请求成功
    activeRequest--
    // ✅ Detect if this is an S3 pre-signed URL upload
    if (res.config.url.includes("telescope-develop.s3.us-east-1.amazonaws.com") ){
      console.log("S3 Upload Success:", res.config.url);
      return Promise.resolve(res);
    }
    if (res.data.error_code !== 0) {
      showNotification({
        type: 'error',
        message: res.data.msg,
        duration: 3000, // Duration is in milliseconds
      })
      // unauthorized, logout current user
      if (res.data.error_code === 401) {
        if (!logoutFlag) {
          logoutFlag = true
          logout()
          window.location.href = "/login";
        }
        if (activeRequest === 0) {
          // last request
          logoutFlag = false
        }
      }

      // session expired, refresh token
      if (res.data.error_code === SESSION_EXPIRED && res.config.meta.curRetry === 0) {
        if (!isRefreshing) {
          isRefreshing = true
          await postRenewSession()
          apiAxios(res.config)
          refreshSubscribers.forEach((r) => apiAxios(r))
          isRefreshing = false
        } else {
          refreshSubscribers.push(res.config)
        }
      }

      return Promise.reject(res.data)
    }
    return Promise.resolve(res)
  },
  (error) => {
    // 请求失败
    activeRequest--
    if (axios.isCancel(error)) {
      console.error('cancel by client')
    } else {
      const config = error.config
      if (config?.meta && config.meta.curRetry !== config.meta.retry) {
        config.meta.curRetry++
        return new Promise((resolve) => {
          setTimeout(
            () => {
              console.warn(`${config.url},retry: ${config.meta.curRetry} times`)
              resolve(apiAxios(config))
            },
            config.meta.retryDelay,
            1000
          )
        })
      }
    }
    return Promise.reject(error)
  }
)
export default apiAxios
