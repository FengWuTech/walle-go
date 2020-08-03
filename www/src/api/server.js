import { axios } from '@/utils/request'
import qs from 'qs'
export function postServer (data) {
  return axios({
    url: '/api/server/',
    method: 'post',
    data: qs.stringify(data)
  })
}
export function getServers (params) {
  return axios({
    url: '/api/server/',
    method: 'get',
    params
  })
}
export function getServer (params) {
  return axios({
    url: '/api/server/' + params.id,
    method: 'get'
  })
}

export function putServer (data) {
  return axios({
    url: '/api/server/' + data.id,
    method: 'put',
    data: qs.stringify(data)
  })
}

export function delServer (data) {
  return axios({
    url: '/api/server/' + data.id,
    method: 'delete'
  })
}
