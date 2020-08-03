import { axios } from '@/utils/request'
import qs from 'qs'

export function getInfo (token) {
  return axios({
    url: '/sso/login/staff/info',
    method: 'get',
    params: { token }
  })
}

export function getUsers (params) {
  return axios({
    url: '/api/user/',
    method: 'get',
    params
  })
}

export function getPostUsers (data) {
  return axios({
    url: '/api/user/',
    method: 'post',
    data: qs.stringify(data)
  })
}

export function fetchUserGet (data) {
  return axios({
    url: '/api/user/' + data.id,
    method: 'get'
  })
}

export function delUser (data) {
  return axios({
    url: '/api/user/' + data.id,
    method: 'delete'
  })
}

export function fetchCasUserRoleSet (data) {
  return axios({
    url: '/api/user/' + data.id,
    method: 'get'
  })
}

export function getProviderInfo () {
  return axios({
    url: '/sso/work/weixin/provider/info',
    method: 'get'
  })
}

export function getSsoWorkWeixinTokenExchange (params) {
  return axios({
    url: '/sso/work/weixin/token/exchange',
    method: 'get',
    params
  })
}

export function ssoVerifyLoginByToken (params) {
  return axios({
    url: '/sso/verify/login/by/token',
    method: 'get',
    params
  })
}

export function ssoLoginCompanyList (data) {
  return axios({
    url: '/sso/login/company/list',
    method: 'post',
    data
  })
}

export function ssoResetPasswordSmsSend (data) {
  return axios({
    url: '/sso/reset/password/sms/send',
    method: 'post',
    data
  })
}

export function ssoResetPassword (data) {
  return axios({
    url: '/sso/reset/password',
    method: 'post',
    data
  })
}

export function ssoLoginStaffResetPassword (data) {
  return axios({
    url: '/sso/login/staff/reset/password',
    method: 'post',
    data
  })
}

export function ssoLoginStaffInfoEdit (data) {
  return axios({
    url: '/sso/login/staff/info/edit',
    method: 'post',
    data
  })
}

export function ssoKancloudLogin (params) {
  return axios({
    url: '/sso/kancloud/login',
    method: 'get',
    params
  })
}

export function ssoTempTokenGenerate (data) {
  return axios({
    url: '/sso/temp/token/generate',
    method: 'post',
    data
  })
}
