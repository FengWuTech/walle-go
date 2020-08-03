import { axios } from '@/utils/request'
import qs from 'qs'
export function postProject (data) {
  return axios({
    url: '/api/project/',
    method: 'post',
    data: qs.stringify(data)
  })
}
export function getProjects (params) {
  return axios({
    url: '/api/project/',
    method: 'get',
    params
  })
}
export function getProject (params) {
  return axios({
    url: '/api/project/' + params.id,
    method: 'get'
  })
}
export function copyProject (data) {
  return axios({
    url: `/api/project/${data.id}/copy`,
    method: 'post'
  })
}
export function detectionProject (data) {
  return axios({
    url: `/api/project/${data.id}/detection`,
    method: 'post'
  })
}

export function putProject (data) {
  return axios({
    url: '/api/project/' + data.id,
    method: 'put',
    data: qs.stringify(data)
  })
}

export function delProject (data) {
  return axios({
    url: '/api/project/' + data.id,
    method: 'delete'
  })
}
