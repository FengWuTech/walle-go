<template>
  <div :style="!$route.meta.hiddenHeaderContent ? 'margin: -24px -24px 0px;' : null">
    <!-- pageHeader , route meta :true on hide -->
    <page-header v-if="!$route.meta.hiddenHeaderContent" :title="pageTitle" :logo="logo" :avatar="avatar" :title-tip="titleTip">
      <slot slot="action" name="action" />
      <slot v-if="headActions && headActions.length && headActions.length > 0" slot="action">
        <a-button style="margin-right: 8px;" v-for="(item, idx) in headActions" :key="idx" :type="item.type" @click="item.onClick">{{ item.label }}</a-button>
      </slot>
      <slot slot="content" name="headerContent" />
      <slot v-if="!this.$slots.headerContent && description" slot="content">
        <p style="font-size: 14px;color: rgba(0,0,0,.65)">{{ description }}</p>
        <div class="link">
          <template v-for="(link, index) in linkList">
            <a :key="index" :href="link.href">
              <a-icon :type="link.icon" />
              <span>{{ link.title }}</span>
            </a>
          </template>
        </div>
      </slot>
      <slot slot="extra" name="extra">
        <div class="extra-img">
          <img v-if="typeof extraImage !== 'undefined'" :src="extraImage">
        </div>
      </slot>
      <div slot="pageMenu">
        <div v-if="search" class="page-menu-search">
          <a-input-search
            style="width: 80%; max-width: 522px;"
            placeholder="请输入..."
            size="large"
            enter-button="搜索"
          />
        </div>
        <div v-if="tabs && tabs.items&&tabs.items.length" class="page-menu-tabs">
          <!-- @change="callback" :activeKey="activeKey" -->
          <a-tabs :tab-bar-style="{margin: 0}" :active-key="tabs.active()" @change="tabs.callback">
            <a-tab-pane v-for="item in getTabs(tabs.items)" :key="item.key" :tab="item.badge?undefined:item.title">
              <span slot="tab" v-if="item.badge">
                <a-badge :count="item.count" :title="item.title ">
                  <div style="padding-right:12px;padding-top:6px">{{ item.title }}</div>
                </a-badge>
              </span>
            </a-tab-pane>
          </a-tabs>
        </div>
      </div>
      <slot v-if="extend" slot="extend">
        <div style="width: 100%;">
          <description-list :title="extend.title || ''" :col="4">
            <description-list-item v-for="(item, index) in extend.list" :key="index" :term="item.term">
              <div v-html="item.value" />
            </description-list-item>
          </description-list>
        </div>
      </slot>
    </page-header>
    <div class="content">
      <div class="page-header-index-wide">
        <slot>
          <!-- keep-alive  -->
          <keep-alive v-if="multiTab">
            <router-view ref="content" />
          </keep-alive>
          <router-view v-else ref="content" />
        </slot>
      </div>
    </div>
  </div>
</template>

<script>
import { mapState } from 'vuex'
import PageHeader from '@/components/PageHeader'
import ARow from 'ant-design-vue/es/grid/Row'
import { DescriptionList } from '@/components/'
const DescriptionListItem = DescriptionList.Item

export default {
  name: 'PageView',
  components: {
    ARow,
    PageHeader,
    DescriptionList,
    DescriptionListItem
  },
  props: {
    avatar: {
      type: String,
      default: null
    },
    title: {
      type: [String, Boolean],
      default: true
    },
    logo: {
      type: String,
      default: null
    },
    directTabs: {
      type: Object,
      default: null
    }
  },
  data () {
    return {
      pageTitle: null,
      description: null,
      extend: null,
      headActions: [],
      linkList: [],
      extraImage: '',
      search: false,
      tabs: {},
      titleTip: null
    }
  },
  computed: {
    ...mapState({
      multiTab: state => state.app.multiTab
    })
  },
  mounted () {
    this.tabs = this.directTabs
    this.getPageMeta()
  },
  // 不确定为什么要加这样的处理，会导致在子组建无法操作父组建的标题
  updated () {
    this.getPageMeta()
  },
  methods: {
    getTabs (data) {
      let hasPath = false
      const newData = data.filter(item => {
        if (this.$store.getters.operationCodeList && item.permission && item.path) {
          if (!hasPath) {
            hasPath = item.path === this.$route.path
          }
          return this.$store.getters.operationCodeList.includes(item.permission)
        }
        return true
      })
      const hasItem = newData.filter(element => {
        return element.path === this.$route.path
      })
      if (hasPath && !hasItem.length) {
        if (newData.length) {
          this.$router.replace(newData[0].path)
        } else {
          this.$router.replace('/403')
        }
      }
      return newData
    },
    getPageMeta () {
      // eslint-disable-next-line
      this.pageTitle = (typeof(this.title) === 'string' || !this.title) ? this.title : this.$route.meta.title

      const content = this.$refs.content
      if (content) {
        if (content.pageMeta) {
          Object.assign(this, content.pageMeta)
        } else {
          this.extend = content.extend
          this.headActions = content.headActions
          this.titleTip = content.titleTip
          this.description = content.description
          this.linkList = content.linkList
          this.extraImage = content.extraImage
          this.search = content.search === true
          this.tabs = content.tabs
        }
      }
    }
  }
}
</script>

<style lang="less" scoped>
  .content {
    margin: 24px 24px 0;
    .link {
      margin-top: 16px;
      &:not(:empty) {
        margin-bottom: 16px;
      }
      a {
        margin-right: 32px;
        height: 24px;
        line-height: 24px;
        display: inline-block;
        i {
          font-size: 24px;
          margin-right: 8px;
          vertical-align: middle;
        }
        span {
          height: 24px;
          line-height: 24px;
          display: inline-block;
          vertical-align: middle;
        }
      }
    }
  }
  .page-menu-search {
    text-align: center;
    margin-bottom: 16px;
  }
  // .page-menu-tabs {
  //   margin-top: 48px;
  // }

  .extra-img {
    margin-top: -60px;
    text-align: center;
    width: 195px;

    img {
      width: 100%;
    }
  }

  .mobile {
    .extra-img{
      margin-top: 0;
      text-align: center;
      width: 96px;

      img{
        width: 100%;
      }
    }
  }
</style>
