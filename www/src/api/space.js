import { axios } from '@/utils/request'
import qs from 'qs'
export function createResource (data) {
  return axios({
    url: '/cas/resource/add',
    headers: {
      appid: data.appid
    },
    method: 'post',
    data
  })
}

export function updateResource (data) {
  return axios({
    url: '/cas/resource/update',
    headers: {
      appid: data.appid
    },
    method: 'post',
    data
  })
}

export function deleteResource (data) {
  return axios({
    url: '/cas/resource/delete',
    headers: {
      appid: data.appid
    },
    method: 'post',
    data
  })
}

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

export function fetchResources (params) {
  return axios({
    url: '/cas/resource/list',
    method: 'get',
    params
  })
}

export function createResourceActionAdd (data) {
  return axios({
    url: '/cas/resource/action/add',
    method: 'post',
    data
  })
}

export function fetchResourceAction (params) {
  return axios({
    url: '/cas/resource/action/list',
    headers: {
      appid: params.appid
    },
    method: 'get',
    params
  })
}

export function deleteResourceAction (data) {
  return axios({
    url: '/cas/resource/action/delete',
    headers: {
      appid: data.appid
    },
    method: 'post',
    data
  })
}

export function fetchCasResourceTree (params) {
  return axios({
    url: '/cas/resource/tree',
    method: 'get',
    params
  })
}
