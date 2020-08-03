import Vue from 'vue'
import { login, logout } from '@/api/login'
import { getGeneralMenu } from '@/api/app'
// import { ACCESS_TOKEN } from '@/store/mutation-types'
import { welcome } from '@/utils/util'
import { removeTokenAll, setToken } from '@/utils/auth' // get token from cookie
var index = 0
const user = {
  state: {
    token: '',
    isLogin: false,
    name: '',
    welcome: '',
    avatar: '',
    roles: [],
    operationCodeList: [],
    menuCodeList: [],
    info: {},
    projectItem: {},
    projectList: []
  },

  mutations: {
    SET_TOKEN: (state, token) => {
      state.token = token
    },
    SET_ISLOGIN: (state, status) => {
      state.isLogin = status
    },
    SET_NAME: (state, { name, welcome }) => {
      state.name = name
      state.welcome = welcome
    },
    SET_AVATAR: (state, avatar) => {
      state.avatar = avatar
    },
    SET_ROLES: (state, roles) => {
      state.roles = roles
    },
    SET_OPERATION_CODE_LIST: (state, operationCodeList) => {
      state.operationCodeList = operationCodeList
    },
    SET_MENU_CODE_LIST: (state, menuCodeList) => {
      state.menuCodeList = menuCodeList
    },
    SET_INFO: (state, info) => {
      state.info = info
    },
    SET_PROJECT_ITEM: (state, item) => {
      state.projectItem = item
      Vue.ls.set('PROJECT_ITEM', item)
    },
    SET_PROJECT_LIST: (state, list) => {
      state.projectList = list
    }
  },

  actions: {
    // 登录
    login ({ dispatch, commit }, userInfo) {
      return new Promise((resolve, reject) => {
        login(userInfo).then(response => {
          if (response.data.token) {
            commit('SET_TOKEN', response.data.token)
            setToken(response.data.token)
            dispatch('getInfo')
            resolve(response)
          } else {
            reject(response)
          }
        }).catch(error => {
          reject(error)
        })
      })
    },

    // 获取权限菜单信息
    getMenu ({ commit }) {
      return new Promise((resolve, reject) => {
        console.log(1212)
        getGeneralMenu().then(response => {
          const operationCodeList = response.data.operation_code_list || []
          // const menuCodeList = response.data.menu_code_list || []
          const projectList = response.data.project_list || []
          const menus = []
          const getPath = (data) => {
            data.forEach(element => {
              if (element.sub_menu && element.sub_menu.length) {
                getPath(element.sub_menu)
              } else {
                menus.push(element)
              }
            })
            return menus
          }
          const menuCodeList = getPath(response.data.menu)
          const result = {
            'name': response.data.user.username,
            // 操作里有跳转到菜单的地方，将菜单权限放到操作权限里，等于查看权限
            'operationCodeList': operationCodeList,
            'menuCodeList': menuCodeList,
            'projectList': projectList,
            'company_name': response.data.company_name,
            'avatar': response.data.user.avatar || '/avatar2.png',
            'role': {
              'permissions': [{
                'roleId': 'admin',
                'permissionId': 'dashboard',
                'permissionName': '仪表盘'
              }]
            }
          }
          commit('SET_ISLOGIN', true)
          if (result.role && result.role.permissions.length > 0) {
            const role = result.role
            role.permissions = result.role.permissions
            role.permissions.map(per => {
              if (per.actionEntitySet != null && per.actionEntitySet.length > 0) {
                const action = per.actionEntitySet.map(action => { return action.action })
                per.actionList = action
              }
            })
            role.permissionList = role.permissions.map(permission => { return permission.permissionId })
            commit('SET_ROLES', result.role)
            commit('SET_INFO', result)
            commit('SET_OPERATION_CODE_LIST', result.operationCodeList)
            commit('SET_MENU_CODE_LIST', result.menuCodeList)
          } else {
            reject(new Error('getInfo: roles must be a non-null array !'))
          }

          commit('SET_NAME', { name: result.name, welcome: welcome() })
          commit('SET_AVATAR', result.avatar)
          commit('SET_PROJECT_LIST', result.projectList)
          if (result.projectList.length > 0) {
            const PROJECT_ITEM = Vue.ls.get('PROJECT_ITEM')
            let setItem = result.projectList[0]
            if (PROJECT_ITEM) {
              const findItem = result.projectList.find(item => {
                return item.id === PROJECT_ITEM.id
              })
              if (findItem) {
                setItem = findItem
              }
            }
            commit('SET_PROJECT_ITEM', setItem)
          }
          // if (index > 10) {
          //   reject(new Error('getInfo: roles must be a non-null array !'))
          // }
          index = index + 1
          resolve(result)
        }).catch(error => {
          reject(error)
        })
      })
    },

    // 登出
    Logout ({ commit, state }) {
      return new Promise((resolve, reject) => {
        logout().then(res => {
          commit('SET_TOKEN', '')
          commit('SET_ROLES', [])
          removeTokenAll()
          // commit('SET_ISLOGIN', false)

          // Vue.ls.remove(ACCESS_TOKEN)
          resolve(res)
        }).catch(error => {
          reject(error)
        })
      })
    }

  }
}

export default user
