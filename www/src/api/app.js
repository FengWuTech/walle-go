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