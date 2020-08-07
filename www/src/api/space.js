import { axios } from '@/utils/request'
import qs from 'qs'

export function fetchSpace (params) {
  return axios({
    url: '/api/space/',
    method: 'get',
    params
  })
}
export function fetchPostSpace (data) {
  return axios({
    url: '/api/space/',
    method: 'post',
    data: qs.stringify(data)
  })
}
export function fetchDelSpace (data) {
  return axios({
    url: '/api/space/' + data.id,
    method: 'delete'
  })
}
export function fetchDelSpaceId (data) {
  return axios({
    url: '/api/space/' + data.id,
    method: 'get'
  })
}