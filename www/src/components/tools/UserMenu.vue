<template>
  <div class="user-wrapper">
    <div class="content-box">
      <!--      <a href="https://pro.loacg.com/docs/getting-started" target="_blank">-->
      <!--        <span class="action">-->
      <!--          <a-icon type="question-circle-o"></a-icon>-->
      <!--        </span>-->
      <!--      </a>-->
      <a href="http://www.gmtech.top/" target="_blank">
        <span class="action">
          <a-tooltip title="公司官网">
            <a-icon style="font-size: 16px; padding: 4px" type="global" />
          </a-tooltip>
        </span>
      </a>
      <!-- <notice-icon class="action"/> -->
      <a-dropdown>
        <span class="action ant-dropdown-link user-dropdown-menu">
          <a-avatar class="avatar" size="small" :src="avatar"/>
          <span>{{ nickname }}</span>
        </span>
        <a-menu slot="overlay" class="user-dropdown-menu-wrapper">
          <!-- <a-menu-item key="0">
            <router-link :to="{ name: 'SystemUserSetting' }">
              <a-icon type="user"/>
              <span>账户资料</span>
            </router-link>
          </a-menu-item> -->
          <!-- <a-menu-item key="1">
            <router-link :to="{ name: 'InvoiceApply' }">
              <a-icon type="wallet" />
              <span>发票申请</span>
            </router-link>
          </a-menu-item> -->
          <!-- <a-menu-divider/> -->
          <a-menu-item key="2">
            <a href="javascript:;" @click="handleLogout">
              <a-icon type="logout"/>
              <span>退出登录</span>
            </a>
          </a-menu-item>
        </a-menu>
      </a-dropdown>
    </div>
  </div>
</template>

<script>
import NoticeIcon from '@/components/NoticeIcon'
import { mapActions, mapGetters } from 'vuex'

export default {
  name: 'UserMenu',
  components: {
    NoticeIcon
  },
  computed: {
    ...mapGetters(['nickname', 'avatar'])
  },
  methods: {
    ...mapActions(['Logout']),
    handleLogout () {
      this.$confirm({
        title: '提示',
        content: '真的要注销登录吗 ?',
        onOk: () => {
          return this.Logout({}).then(() => {
            window.location.reload()
          }).catch(err => {
            this.$message.error({
              title: '错误',
              description: err.message
            })
          })
        },
        onCancel () {
        }
      })
    }
  }
}
</script>
