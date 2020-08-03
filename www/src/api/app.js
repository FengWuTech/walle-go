import { axios } from '@/utils/request'

export function getGeneralMenu (params) {
  return axios({
    url: '/api/general/menu',
    method: 'get',
    params
  })
}
export function getGeneralInfo (params) {
  return axios({
    url: '/api/general/info',
    method: 'get',
    params
  })
}

export function createApp (data) {
  return axios({
    url: '/cas/app/add',
    method: 'post',
    data
  })
}

export function updateApp (data) {
  return axios({
    url: '/cas/app/update',
    method: 'post',
    data
  })
}

export function fetchApp (params) {
  return axios({
    url: '/cas/app/get',
    method: 'get',
    params
  })
}

export function fetchApps (params) {
  return axios({
    url: '/cas/app/list',
    method: 'get',
    params
  })
}
