import { axios } from '@/utils/request'
import qs from 'qs'
export function postTask (data) {
  return axios({
    url: '/api/task/',
    method: 'post',
    data: qs.stringify(data)
  })
}
export function getTasks (params) {
  return axios({
    url: '/api/task/',
    method: 'get',
    params
  })
}
export function getTask (params) {
  return axios({
    url: '/api/task/' + params.id,
    method: 'get'
  })
}

export function putTask (data) {
  return axios({
    url: '/api/task/' + data.id,
    method: 'put',
    data: qs.stringify(data)
  })
}

export function delTask (data) {
  return axios({
    url: '/api/task/' + data.id,
    method: 'delete'
  })
}

export function auditRejectTask (data) {
  return axios({
    url: `/api/task/${data.id}/reject`,
    method: 'put'
  })
}

export function rejectTask (data) {
  return axios({
    url: `/api/task/${data.id}/reject`,
    method: 'put'
  })
}
export function auditTask (data) {
  return axios({
    url: `/api/task/${data.id}/audit`,
    method: 'put'
  })
}

export function rollbackTask (data) {
  return axios({
    url: `/api/task/${data.id}/rollback`,
    method: 'put'
  })
}
