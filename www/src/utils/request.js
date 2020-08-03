// import Vue from 'vue'
import axios from 'axios'
import store from '@/store'
import notification from 'ant-design-vue/es/notification'
import { VueAxios } from './axios'
// import { ACCESS_TOKEN } from '@/store/mutation-types'
import { getToken, getAppIdToken, removeToken } from '@/utils/auth' // get token from cookie

// 创建 axios 实例
const service = axios.create({
  baseURL: process.env.VUE_APP_API_BASE_URL, // api base_url
  // Content-Type: 'application/x-www-form-urlencoded',
  // headers: {
  //   'Content-type': 'application/x-www-form-urlencoded'
  // },
  timeout: 6000 // 请求超时时间
})
// axios.defaults.headers = {
//   'Content-type': 'application/x-www-form-urlencoded'
// }

const err = (error) => {
  if (error.response) {
    const data = error.response.data
    const token = getToken()
    if (error.response.status === 403) {
      notification.error({
        message: 'Forbidden',
        description: data.message
      })
    }
    if (error.response.status === 401 && !(data.result && data.result.isLogin)) {
      notification.error({
        message: 'Unauthorized',
        description: 'Authorization verification failed'
      })
      if (token) {
        store.dispatch('Logout').then(() => {
          setTimeout(() => {
            window.location.reload()
          }, 1500)
        })
      }
    }
  }
  return Promise.reject(error)
}

// request interceptor
service.interceptors.request.use(config => {
  const token = getToken()
  const appId = getAppIdToken()
  if (token) {
    config.headers['token'] = token // 让每个请求携带自定义 token 请根据实际情况自行修改

    config.headers['appid'] = appId // 让每个请求携带自定义 token 请根据实际情况自行修改
  }
  if (['/sso/staff/logout', '/sso/temp/token/exchange'].includes(config.url)) {
    config.url = `${process.env.VUE_APP_API_SSO_BASE_URL}${config.url}`
    delete config.headers
  } else {
    config.baseURL = process.env.VUE_APP_API_BASE_URL
  }
  return config
}, err)

// response interceptor
service.interceptors.response.use(async (response) => {
  if (response.data.code === 10702 || response.data.code === 10101) {
    notification.error({
      message: '重新登录',
      description: '登录状态失效，请重新登录'
    })
    removeToken()
    location.reload(true)
  }

  if (response.data.code === 1000) {
    location.href = '/login'
    return
  }
  if (response.data.code === 10703) {
    notification.destroy()
    notification.error({
      message: '非管理员',
      description: '请核对登录账号，确保账号有管理权限'
    })
    setTimeout(async () => {
      try {
        await store.dispatch('Logout')
      } catch (error) {
        console.error(error)
      }
    }, 2000)
  } else if (response.data.code !== 0 && response.data.code !== 200) {
    notification.destroy()
    notification.error({
      message: '网络错误',
      description: response.data.msg
    })
    return Promise.reject(response.data)
  }

  return response.data
}, err)

const installer = {
  vm: {},
  install (Vue) {
    Vue.use(VueAxios, service)
  }
}

export {
  installer as VueAxios,
  service as axios
}
