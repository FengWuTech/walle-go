<template>
  <a-config-provider :locale="locale">
    <div id="app">
      <fw-watermark v-if="NODE_ENV" :input-text="` ${userInfo.name || ''}`" />
      <router-view />
    </div>
  </a-config-provider>
</template>

<script>
import { mapGetters } from 'vuex'
import FwWatermark from '@/components/FwWatermark'
import { AppDeviceEnquire } from '@/utils/mixin'
import { getConfig } from '@/api/common'
import zhCN from 'ant-design-vue/lib/locale-provider/zh_CN'
export default {
  components: {
    FwWatermark
  },
  mixins: [AppDeviceEnquire],
  data () {
    return {
      locale: zhCN,
      NODE_ENV: process.env.NODE_ENV === 'production'
    }
  },
  created () {
    // this.getCommonConfig()
  },
  computed: {
    ...mapGetters([
      'userInfo'
    ])
  },
  methods: {
    getPopupContainer (el, dialogContext) {
      if (dialogContext) {
        return dialogContext.getDialogWrap()
      } else {
        return document.body
      }
    },
    getCommonConfig () {
      getConfig({}).then(res => {
        if (res.code === 200) {
          this.$store.commit('common/SET_APP_CONFIG', res.data)
        } else {
          this.$message.error(res.msg)
        }
      })
    }
  }
}
</script>
<style lang="less">
  #app {
    height: 100%;
  }
</style>
