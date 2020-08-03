import { axios } from '@/utils/request'

export function getConfig (params) {
  return axios({
    url: '/pubsrv/config/get',
    method: 'get',
    params
  })
}

export function getUploadToken (data) {
  return axios({
    url: '/pubsrv/upload/token',
    method: 'post',
    data: data
  })
}

/**
 * 获取上传token
 * @returns {*}
 */
export function getTecentUploadToken (data) {
  return axios({
    url: '/pubsrv/upload/cos_token',
    method: 'post',
    data: data
  })
}
