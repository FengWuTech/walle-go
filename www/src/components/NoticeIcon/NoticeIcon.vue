<template>
  <a-popover
    v-model="visible"
    trigger="click"
    placement="bottomRight"
    overlayClassName="header-notice-wrapper"
    :getPopupContainer="() => $refs.noticeRef.parentElement"
    :autoAdjustOverflow="true"
    :arrowPointAtCenter="true"
    :overlayStyle="{ width: '300px', top: '50px' }"
  >
    <template slot="content">
      <div class="message-header">
        <div class="title">未读消息</div>
        <div class="extra"><a class="user-link" href="javascript:;" @click="showMore">更多</a></div>
      </div>
      <a-list>
        <a-list-item v-for="(item, idx) in messageList" :key="idx">
          <a-list-item-meta :title="`${item.content.length>13?item.content.slice(0, 13)+'...':item.content}`" :description="item.create_time|moment">
            <a-avatar slot="avatar" style="backgroundColor:#87d068" icon="mail"/>
          </a-list-item-meta>
        </a-list-item>
      </a-list>
    </template>
    <span @click="fetchNotice" class="header-notice" ref="noticeRef">
      <a-badge :count="count">
        <a-icon style="font-size: 16px; padding: 4px" type="bell" />
      </a-badge>
    </span>
  </a-popover>
</template>

<script>
// import { getSysMessage } from '@/api/system/system'
import { mapGetters } from 'vuex'

export default {
  name: 'HeaderNotice',
  data () {
    return {
      visible: false,
      count: 0,
      messageList: []
    }
  },
  computed: {
    ...mapGetters([
      'systemMessageStatus'
    ])
  },
  watch: {
    systemMessageStatus () {
      // this.fetchNotice()
    }
  },
  created () {
    // this.fetchNotice()
  },
  methods: {
    // fetchNotice () {
    //   const data = {
    //     read_status: 1,
    //     page: 1,
    //     page_size: 1000
    //   }
    //   getSysMessage(data).then(res => {
    //     if (res.code === 200) {
    //       this.messageList = res.data.list
    //       this.count = res.data.total
    //     } else {
    //       this.$message.error(res.msg)
    //     }
    //   })
    // },

    showMore () {
      this.$router.push({ name: 'SystemMessage' })
      this.visible = false
    }
  }
}
</script>

<style lang="css">
  .header-notice-wrapper {
    top: 50px !important;
  }
</style>
<style lang="less" scoped>
  .header-notice{
    display: inline-block;
    transition: all 0.3s;

    span {
      vertical-align: initial;
    }
  }
  .message-header {
    display: flex;
    justify-content: space-between;

    .title {
      float: left;
    }

    .extra {
      float: right;
    }
  }
</style>
