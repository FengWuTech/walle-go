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

export function createUserRole (data) {
  return axios({
    url: '/cas/user/addRole',
    method: 'post',
    data
  })
}

export function deleteRoleByRefId (data) {
  return axios({
    url: '/cas/user/deleteRoleByRefId',
    method: 'post',
    data
  })
}

export function createUser (data) {
  return axios({
    url: '/cas/user/role/add',
    method: 'post',
    data
  })
}

export function updateUser (data) {
  return axios({
    url: '/cas/user/role/update',
    method: 'post',
    data
  })
}

export function deleteUser (data) {
  return axios({
    url: '/cas/user/role/delete',
    method: 'post',
    data
  })
}

export function fetchUser (params) {
  return axios({
    url: '/cas/user/role/get',
    method: 'get',
    params
  })
}

export function fetchUsers (params) {
  return axios({
    url: '/cas/user/role/list',
    method: 'get',
    params
  })
}
