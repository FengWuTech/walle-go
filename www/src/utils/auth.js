import Cookies from 'js-cookie'

const TokenKey = 'wgsession'

let domain = {}
if (~document.domain.indexOf('gmtech.top')) {
  domain = { path: '/', domain: 'gmtech.top' }
}
export function getToken () {
  return Cookies.get(TokenKey, domain)
}

export function setToken (token) {
  return Cookies.set(TokenKey, token, { ...domain, expires: 30 })
}
export function removeToken () {
  return Cookies.remove(TokenKey, domain)
}

export function getAppIdToken () {
  return Cookies.get('appid', domain)
}

export function setAppIdToken (token) {
  return Cookies.set('appid', token, { ...domain, expires: 30 })
}

export function removeAppIdToken () {
  return Cookies.remove('appid', domain)
}

export function removeTokenAll () {
  Cookies.remove(TokenKey, domain)
  return Cookies.remove('session', domain)
}
