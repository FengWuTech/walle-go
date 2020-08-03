// with polyfills
import 'core-js/stable'
import 'regenerator-runtime/runtime'

import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store/'
import { VueAxios } from './utils/request'
import { scrollTo } from '@/utils/scroll-to'
import driver from '@/components/Driver/index.js'

import VueCodemirror from '@/components/codemirror'

// import base style
import 'codemirror/lib/codemirror.css'

import permission from '@/directive/permission/index.js' // 权限判断指令

// mock
// WARNING: `mockjs` NOT SUPPORT `IE` PLEASE DO NOT USE IN `production` ENV.
// import './mock'

import bootstrap from './core/bootstrap'
import './core/lazy_use'
import './permission' // permission control
import './utils/filter' // global filter
import './components/global.less'
import './styles/customStyle.less'

import * as Sentry from '@sentry/browser'
import * as Integrations from '@sentry/integrations'

// import socketio from 'socket.io-client'

// export const SocketInstance = socketio()
// Vue.use(new VueSocketIO({
//   debug: true,
//   connection: socketio('/walle'),
//   options: {

//   } // Optional options
// }))
import VueSocketIO from 'vue-socket.io'
// import VueSocketIO from 'vue-socket.io'

// import more codemirror resource...

// you can set default global options and events when Vue.use
Vue.use(VueCodemirror /* {
  // options: { theme: 'base16-dark', ... },
  events: ['scroll', ...]
} */)
Vue.use(new VueSocketIO({
  debug: false,
  connection: '/walle',
  vuex: {
    store,
    actionPrefix: 'SOCKET_', // 为vuex设置的两个前缀
    mutationPrefix: 'SOCKET_'
  }
}))

// Vue.use(VueSocketIO, SocketInstance)

process.env.NODE_ENV === 'production' && Sentry.init({
  dsn: 'https://72d43d46759d46c6baafb3ef69a74fe5@sentry.io/5184387',
  integrations: [new Integrations.Vue({ Vue, attachProps: true })]
})

Vue.config.productionTip = false

// mount axios Vue.$http and this.$http
Vue.use(VueAxios)
Vue.use(permission)
Vue.prototype.$scrollTo = scrollTo
Vue.prototype.$driver = driver
new Vue({
  router,
  store,
  created: bootstrap,
  render: h => h(App)
}).$mount('#app')
