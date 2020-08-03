import { axios } from '@/utils/request'
import qs from 'qs'
export function postEnvironment (data) {
  return axios({
    url: '/api/environment/',
    method: 'post',
    data: qs.stringify(data)
  })
}
export function getEnvironments (params) {
  return axios({
    url: '/api/environment/',
    method: 'get',
    params
  })
}
export function getEnvironment (params) {
  return axios({
    url: '/api/environment/' + params.id,
    method: 'get'
  })
}

export function putEnvironment (data) {
  return axios({
    url: '/api/environment/' + data.id,
    method: 'put',
    data: qs.stringify(data)
  })
}

export function delEnvironment (data) {
  return axios({
    url: '/api/environment/' + data.id,
    method: 'delete'
  })
}
