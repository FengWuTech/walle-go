<template>
  <transition name="showHeader">
    <div v-if="visible" class="header-animat">
      <a-layout-header
        v-if="visible"
        :class="[fixedHeader && 'ant-header-fixedHeader', sidebarOpened ? 'ant-header-side-opened' : 'ant-header-side-closed', ]"
        :style="{ padding: '0' }">
        <div v-if="mode === 'sidemenu'" class="header">
          <a-icon v-if="device==='mobile'" class="trigger" :type="collapsed ? 'menu-fold' : 'menu-unfold'" @click="toggle"/>
          <a-icon v-else class="trigger" :type="collapsed ? 'menu-unfold' : 'menu-fold'" @click="toggle"/>
          <!-- <a-select
            style="width:240px"
            @change="handleChange"
            v-model="appid"
          >
            <a-select-option v-for="item in apps" :value="item.id" :key="item.id">{{ item.name }}</a-select-option>
          </a-select> -->
          <a-button color="#fff" type="link" @click="$router.push('/')">切换</a-button>
          <user-menu></user-menu>
        </div>

        <div v-if="mode === 'blank'" class="header">
          <logo className="top-nav-header-blank" style="display: inline-block;" :show-title="device !== 'mobile'"/>
          <user-menu></user-menu>
        </div>
        <div v-if="mode !== 'blank' && mode !== 'sidemenu'" :class="['top-nav-header-index', theme]">
          <div class="header-index-wide">
            <div class="header-index-left">
              <logo class="top-nav-header" :show-title="device !== 'mobile'"/>
              <s-menu v-if="device !== 'mobile'" mode="horizontal" :menu="menus" :theme="theme" />
              <a-icon v-else class="trigger" :type="collapsed ? 'menu-fold' : 'menu-unfold'" @click="toggle" />
            </div>
            <user-menu class="header-index-right"></user-menu>
          </div>
        </div>
      </a-layout-header>
    </div>
  </transition>
</template>

<script>
import UserMenu from '../tools/UserMenu'
import SMenu from '../Menu/'
import Logo from '../tools/Logo'
import { mixin } from '@/utils/mixin'
import { apps } from '@/mixin/apps'
import { setAppIdToken, getAppIdToken } from '@/utils/auth' // get token from cookie

export default {
  name: 'GlobalHeader',
  components: {
    UserMenu,
    SMenu,
    Logo
  },
  mixins: [mixin, apps],
  props: {
    mode: {
      type: String,
      // sidemenu, topmenu
      default: 'sidemenu'
    },
    menus: {
      type: Array,
      required: true
    },
    theme: {
      type: String,
      required: false,
      default: 'dark'
    },
    collapsed: {
      type: Boolean,
      required: false,
      default: false
    },
    device: {
      type: String,
      required: false,
      default: 'desktop'
    }
  },
  data () {
    return {
      visible: true,
      oldScrollTop: 0,
      appid: ~~getAppIdToken()
    }
  },
  mounted () {
    document.addEventListener('scroll', this.handleScroll, { passive: true })
  },
  methods: {
    handleScroll () {
      if (!this.autoHideHeader) {
        return
      }

      const scrollTop = document.body.scrollTop + document.documentElement.scrollTop
      if (!this.ticking) {
        this.ticking = true
        requestAnimationFrame(() => {
          if (this.oldScrollTop > scrollTop) {
            this.visible = true
          } else if (scrollTop > 300 && this.visible) {
            this.visible = false
          } else if (scrollTop < 300 && !this.visible) {
            this.visible = true
          }
          this.oldScrollTop = scrollTop
          this.ticking = false
        })
      }
    },
    toggle () {
      this.$emit('toggle')
    },
    handleChange (value) {
      const that = this
      this.$confirm({
        title: '确认要切换关联应用吗?',
        content: '切换成功后，将要更新所有页面',
        okText: '确定',
        okType: 'danger',
        cancelText: '取消',
        onOk () {
          location.reload()
          setAppIdToken(value)
        },
        onCancel () {
          that.appid = ~~getAppIdToken()
        }
      })
    }
  },
  beforeDestroy () {
    document.body.removeEventListener('scroll', this.handleScroll, true)
  }
}
</script>

<style lang="less">
@import '../index.less';

.header-animat{
  position: relative;
  z-index: @ant-global-header-zindex;
}
.showHeader-enter-active {
  transition: all 0.25s ease;
}
.showHeader-leave-active {
  transition: all 0.5s ease;
}
.showHeader-enter, .showHeader-leave-to {
  opacity: 0;
}
</style>
