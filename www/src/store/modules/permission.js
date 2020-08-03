import { asyncRouterMap, constantRouterMap } from '@/config/router.config'

/**
 * 过滤账户是否拥有某一个权限，并将菜单从加载列表移除
 *
 * @param permission
 * @param route
 * @returns {boolean}
 */
function hasPermission (permission, route, menuCodeList) {
  // console.log(route)
  // return menuCodeList.filter(item => route.path === item.url).length > 0
  return true
  // if (route.meta && route.meta.permission) {
  //   if (typeof (route.meta.permission) === 'string') {
  //     if (menuCodeList) {
  //       return menuCodeList.includes(route.meta.permission)
  //     }
  //     return false
  //   }
  //   let flag = false
  //   for (let i = 0, len = permission.length; i < len; i++) {
  //     flag = route.meta.permission.includes(permission[i])
  //     if (flag) {
  //       return true
  //     }
  //   }
  //   return false
  // }
  // return true
}

/**
 * 单账户多角色时，使用该方法可过滤角色不存在的菜单
 *
 * @param roles
 * @param route
 * @returns {*}
 */
// eslint-disable-next-line
function hasRole(roles, route) {
  if (route.meta && route.meta.roles) {
    return route.meta.roles.includes(roles.id)
  } else {
    return true
  }
}

function filterAsyncRouter (routerMap, roles, menuCodeList) {
  const accessedRouters = routerMap.filter(route => {
    if (hasPermission(roles.permissionList, route, menuCodeList)) {
      if (route.children && route.children.length) {
        route.children = filterAsyncRouter(route.children, roles, menuCodeList)
      }
      return true
    }
    return false
  })
  return accessedRouters
}

const permission = {
  state: {
    routers: constantRouterMap,
    addRouters: []
  },
  mutations: {
    SET_ROUTERS: (state, routers) => {
      state.addRouters = routers
      state.routers = constantRouterMap.concat(routers)
    }
  },
  actions: {
    GenerateRoutes ({ commit }, data) {
      return new Promise(resolve => {
        const { roles, menuCodeList } = data
        console.log(menuCodeList)
        const accessedRouters = filterAsyncRouter(asyncRouterMap, roles, menuCodeList)
        console.log(accessedRouters)
        commit('SET_ROUTERS', accessedRouters)
        resolve(accessedRouters)
      })
    }
  }
}

export default permission
