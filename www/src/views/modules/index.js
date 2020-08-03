// import LongText from './LongText.vue'
// import FwCascader from './FwCascader.vue'
// import LinkBaseText from './LinkBaseText.vue'
// import FwTreeDrawer from './FwTreeDrawer.vue'
import InputNumberWithAddon from './InputNumberWithAddon.vue'

const Modules = {}

// 系统常用的组件可以在这里注册成全局组件，避免页面内再引的麻烦
Modules.install = Vue => {
  // Vue.component('LongText', LongText)
  // Vue.component('FwCascader', FwCascader)
  // Vue.component('LinkBaseText', LinkBaseText)
  // Vue.component('FwTreeDrawer', FwTreeDrawer)
  Vue.component('InputNumberWithAddon', InputNumberWithAddon)
}

export default Modules
