const getters = {
  device: state => state.app.device,
  isLogin: state => state.user.isLogin,
  theme: state => state.app.theme,
  color: state => state.app.color,
  token: state => state.user.token,
  avatar: state => state.user.avatar,
  nickname: state => state.user.name,
  welcome: state => state.user.welcome,
  roles: state => state.user.roles,
  operationCodeList: state => state.user.operationCodeList,
  menuCodeList: state => state.user.menuCodeList,
  userInfo: state => state.user.info,
  appConfig: state => state.common.appConfig,
  systemMessageStatus: state => state.common.systemMessageStatus,
  addRouters: state => state.permission.addRouters,
  multiTab: state => state.app.multiTab,
  lang: state => state.i18n.lang
}

export default getters
