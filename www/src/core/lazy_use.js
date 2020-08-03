import Vue from 'vue'
import VueStorage from 'vue-ls'
import config from '@/config/defaultSettings'

// base library
import '@/core/lazy_lib/components_use'

// ext library
import VueClipboard from 'vue-clipboard2'
import PageLoading from '@/components/PageLoading'
import PermissionHelper from '@/utils/helper/permission'
import './directives/action'
import VueQrcode from '@chenfengyuan/vue-qrcode'
import 'viewerjs/dist/viewer.css'
import Viewer from 'v-viewer'
import VeRing from 'v-charts/lib/ring.common'
import VePie from 'v-charts/lib/pie.common'
import VeHistogram from 'v-charts/lib/histogram.common'
import VeLine from 'v-charts/lib/line.common'

// 文本长度过长的时候用这个组件，默认长度15个字，多的隐藏显示省略号
import Modules from '@/views/modules/index'

import Print from '@/utils/print'

VueClipboard.config.autoSetContainer = true

Vue.use(PageLoading)
Vue.use(VueStorage, config.storageOptions)
Vue.use(VueClipboard)
Vue.use(PermissionHelper)
Vue.component(VueQrcode.name, VueQrcode)
Vue.use(Modules)
Vue.use(Viewer)
Vue.use(Print)
Vue.component(VeRing.name, VeRing) // 按需引入
Vue.component(VePie.name, VePie) // 按需引入
Vue.component(VeHistogram.name, VeHistogram) // 按需引入
Vue.component(VeLine.name, VeLine) // 按需引入
