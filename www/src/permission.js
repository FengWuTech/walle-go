// import Vue from 'vue'
import router from './router'
import store from './store'
import { ssoTempTokenExchange } from '@/api/login'

import NProgress from 'nprogress' // progress bar
import '@/components/NProgress/nprogress.less' // progress bar custom style
// import notification from 'ant-design-vue/es/notification'
import { setDocumentTitle, domTitle } from '@/utils/domUtil'
// import { ACCESS_TOKEN } from '@/store/mutation-types'
import { getToken, setToken } from '@/utils/auth' // get token from cookie

NProgress.configure({ showSpinner: false }) // NProgress Configuration

const whiteList = ['login', '/login', 'register', 'registerResult'] // no redirect whitelist

// const loginUrl = () => {
//   // const redirectUri = encodeURIComponent(`${process.env.VUE_APP_API_BASE_PROTOCOL}${window.location.host}${window.location.pathname}`)
//   // return `${process.env.VUE_APP_API_BASE_SSO_URL}?from=${redirectUri}`
//   return '/login'
// }
const loginRoutePath = '/login'

function hasPermission (permissions) {
  if (!permissions) return true
  // if (store.getters.menuCodeList.length === 0) return false
  if (typeof (permissions) === 'string') {
    if (store.getters.menuCodeList) {
      return store.getters.menuCodeList.includes(permissions)
    }
    return false
  }
  const permissionList = store.getters.roles && store.getters.roles.permissionList
  return permissionList.some(role => permissions.indexOf(role) >= 0)
}

let firstPermissionedRoute = ''
// 获取第一个有权限的路由页面
function getFirstPermissionedRoute (addRouters) {
  console.log(addRouters)
  if (firstPermissionedRoute) {
    return
  }
  for (var i = 0; i < addRouters.length; i++) {
    if (firstPermissionedRoute) {
      break
    }
    const item = addRouters[i]
    if (item.children && item.children.length > 0) {
      getFirstPermissionedRoute(item.children)
    } else {
      if (item.meta && item.meta.permission && hasPermission(item.meta.permission)) {
        firstPermissionedRoute = item.path
      }
    }
  }
}

router.beforeEach(async (to, from, next) => {
  NProgress.start() // start progress bar

  var hasToken = getToken()
  if (to.query.token) {
    hasToken = to.query.token
    delete to.query.token
    delete to.query.from
    try {
      const tempToken = await ssoTempTokenExchange({ token: hasToken })
      setToken(tempToken.data.token)
    } catch (error) {
      console.error(error)
    }
  }
  to.meta && (typeof to.meta.title !== 'undefined' && setDocumentTitle(`${to.meta.title} - ${domTitle}`))
  if (hasToken) {
    /* has token */
    if (whiteList.indexOf(to.name) !== -1) {
      getFirstPermissionedRoute(store.getters.addRouters)
      router.push({ path: firstPermissionedRoute })
      NProgress.done()
    } else {
      // 有nickname代表已经获取到了用户信息
      const hasGotInfo = store.getters.isLogin
      if (hasGotInfo) {
        if (hasPermission(to.meta.permission)) {
          // console.log(to.path)
          // if (!getAppIdToken() && (to.path).indexOf('/apps/') === -1) {
          //   next({ path: '/apps/list', replace: true })
          // } else {
          next()
          // }
        } else {
          next({ path: '/403', replace: true, query: { noGoBack: true } })
        }
      } else {
        try {
          const { role: roles, menuCodeList } = await store.dispatch('getMenu')
          const accessRoutes = await store.dispatch('GenerateRoutes', { roles, menuCodeList })
          if (accessRoutes.length === 2 && accessRoutes[0].children.length === 0) {
            next({ path: '/403' })
            return
          }
          router.addRoutes(accessRoutes)
          const redirect = decodeURIComponent(from.query.redirect || to.path)
          if (redirect === '/') {
            getFirstPermissionedRoute(store.getters.addRouters)
            router.push({ path: firstPermissionedRoute })
          } else {
            next({ path: redirect, replace: true, query: to.query, params: to.params })
          }
        } catch (e) {
          // notification.error({
          //   message: '错误',
          //   description: '请求用户信息失败，请重试'
          // })
          console.log(e)
          // store.dispatch('Logout').then(() => {
          //   location.href = loginUrl()
          // })
        }
      }
    }
  } else {
    // location.href = loginUrl()
    console.log(to)
    if (whiteList.includes(to.name)) {
      // 在免登录白名单，直接进入
      next()
    } else {
      next({ path: loginRoutePath, query: { redirect: to.fullPath } })
      NProgress.done() // if current page is login will not trigger afterEach hook, so manually handle it
    }
  }
})

router.afterEach(() => {
  NProgress.done() // finish progress bar
})
