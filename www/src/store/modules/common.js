const state = {
  appConfig: {},
  systemMessageStatus: false
}

const mutations = {
  SET_APP_CONFIG: (state, config) => {
    state.appConfig = config
  },
  /**
   * systemMessageStatus 用来表示系统消息发生了变化
   * 右上角消息通知监听该变化重新获取数据实时改变通知数
   */
  toggle_SystemMessageStatus: (state) => {
    state.systemMessageStatus = !state.systemMessageStatus
  }
}

const actions = {
}

export default {
  namespaced: true,
  state,
  mutations,
  actions
}
