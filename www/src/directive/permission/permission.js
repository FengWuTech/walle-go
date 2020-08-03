import store from '@/store'

export default {
  inserted (el, binding, vnode) {
    const { value } = binding
    const operationCodeList = store.getters && store.getters.operationCodeList

    if (value && value instanceof Array && value.length > 0) {
      const permissionRoles = value

      const hasPermission = operationCodeList.some(role => {
        return permissionRoles.includes(role)
      })

      if (!hasPermission) {
        el.parentNode && el.parentNode.removeChild(el)
      }
    } else {
      throw new Error(`need roles! Like v-permission="['admin','editor']"`)
    }
  }
}
