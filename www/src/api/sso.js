import { axios } from '@/utils/request'

export function fetchCasSsoCompanyList (params) {
  return axios({
    url: '/cas/sso/company/list',
    method: 'get',
    params
  })
}

export function fetchCasSsoStaffList (params) {
  return axios({
    url: '/cas/sso/staff/list',
    method: 'get',
    params
  })
}