import { axios } from '@/utils/request'

/**
 * login func
 * parameter: {
 *     username: '',
 *     password: '',
 *     remember_me: true,
 *     captcha: '12345'
 * }
 * @param parameter
 * @returns {*}
 */
export function login (data) {
  return axios({
    url: '/api/passport/login',
    method: 'post',
    headers: {
      'Content-Type': 'application/json;charset=UTF-8'
    },
    data
  })
}

export function logout () {
  return axios({
    url: '/api/passport/logout',
    method: 'post'
  })
}

export function getInfo () {
  return axios({
    url: '/sso/login/staff/info',
    method: 'get',
    headers: {
      'Content-Type': 'application/json;charset=UTF-8'
    }
  })
}

export function getCurrentUserNav (token) {
  return axios({
    url: '/user/nav',
    method: 'get'
  })
}

export function ssoTempTokenExchange (parameter) {
  return axios({
    url: '/sso/temp/token/exchange',
    method: 'post',
    data: parameter
  })
}
