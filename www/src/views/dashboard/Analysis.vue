<template>
  <page-view :avatar="avatar" :title="false">
    <div slot="headerContent">
      <div class="title"></div>
      <div>
        <!--        微组团 企业版| 到期时间：2019-12-05-->
        <!--        <a class="user-link" href="/data/xiaoqu/xiaoqu-renew" style="margin-left: 10px;">续费</a>-->
        <!--        <a class="user-link" href="/data/xiaoqu/xiaoqu-renew" style="margin-left: 10px;">购买组团</a>-->
      </div>
    </div>
    <div slot="extra" style="min-width: 242px;display: flex;">
      <head-info title="可添加组团" content="0" :center="true" :bordered="true"/>
      <head-info title="剩余短信" content="0" :center="true" :bordered="true"/>
      <head-info title="剩余人脸核身资源" content="0" :center="true" :bordered="true" />
      <head-info title="微豆" content="0" :center="true" />
    </div>
    <a-row :gutter="24">
      <a-col :xl="18" :lg="24" :md="24" :sm="24" :xs="24">
        <a-card
          class="project-list"
          :loading="loading"
          style="margin-bottom: 24px;"
          :bordered="false"
          :body-style="{ padding: 0 }">
          <span slot="title">
            基础数据
            <a-tooltip placement="right" v-if="baseInfo.update_time">
              <template slot="title">
                <span>每日更新，最后更新时间：{{ moment(baseInfo.update_time).format('YYYY-MM-DD HH:mm:ss') }}</span>
              </template>
              <a-icon type="question-circle" />
            </a-tooltip>
          </span>
          <div>
            <a-card-grid :id="item.id||''" class="project-card-grid" style="width:25%;padding: 16px 24px;" :key="i" v-for="(item, i) in projects">
              <a-card :bordered="false" :body-style="{ padding: 0 }">
                <a-card-meta>
                  <div slot="title" class="card-title">
                    <a-avatar size="small" :src="item.cover"/>
                    <div style="display: flex;justify-content: space-between;flex: 1;">
                      <router-link :to="{name: item.routerName}" >{{ item.title }}</router-link>
                      <router-link :to="{name: item.routerName}" >{{ item.extra }}</router-link>
                    </div>
                  </div>
                </a-card-meta>
                <div class="project-item">
                  <router-link :to="{name: link.routerName}" :key="idx" v-for="(link, idx) in item.links">{{ link.label }}</router-link>
                </div>
              </a-card>
            </a-card-grid>
          </div>
        </a-card>

        <!-- <a-card :loading="loading" title="应收已收欠费统计" :bordered="false" style="margin-bottom: 24px;">
          <div slot="extra">
            <a-range-picker :style="{width: '256px'}" />
          </div>
          <div>
            <a-col
              :xl="11"
              :lg="11"
              :md="24"
              :sm="24"
              :xs="24"
              style="text-align: center;"
            >
              <ve-ring :data="chartData" height="300px" :settings="chartSettings"></ve-ring>
            </a-col>
            <a-col :xl="13" :lg="13" :md="24" :sm="24" :xs="24">
              <a-table
                style="padding-top: 0;"
                :pagination="false"
                row-key="id"
                :columns="columns"
                :data-source="feeData"
              >
              </a-table>
            </a-col>
          </div>
        </a-card> -->
        <!-- <a-card :loading="loading" title="短信消耗统计" :bordered="false" style="margin-bottom: 24px;">
          <div>
            <a-col
              :xs="24"
              style="text-align: center;"
            >

              <ve-line :legend-visible="true" height="300px" :data="chartLineData" :grid="grid" :settings="chartLineSettings"></ve-line>
            </a-col>
          </div>
        </a-card> -->

        <a-card :loading="loading" title="操作日志" :bordered="false" :body-style="{padding: 0}">
          <a slot="extra" href="javascript:;" @click="$router.push('/system/log')">查看更多</a>
          <div style="padding: 0 24px 8px;">
            <a-list itemLayout="horizontal" size="large" :dataSource="logs" rowKey="log_id">
              <a-list-item
                slot="renderItem"
                slot-scope="item">
                <a-list-item-meta>
                  <template slot="title">

                    <div>
                      <a-row :gutter="8">
                        <a-col :span="8">操作时间：{{ item.create_time | moment }}</a-col>
                        <a-col :span="4">操作员：{{ item.name }}</a-col>
                        <a-col
                          :span="12"
                          style="overflow: hidden;
    text-overflow: ellipsis;
    white-space: nowrap;">操作内容：{{ item.remark }}</a-col>
                      </a-row>
                      <!-- <span class="m-l-sm"></span>
                      <span class="m-l-md"></span>
                      <span
                        class="m-l-md"
                        style="overflow: hidden;
    text-overflow: ellipsis;
    white-space: nowrap;
    width: 60%;"></span> -->
                    </div>
                  </template>
                </a-list-item-meta>
              </a-list-item>
            </a-list>
          </div>
        </a-card>
      </a-col>
      <a-col
        :xl="6"
        :lg="24"
        :md="24"
        :sm="24"
        :xs="24"
        style="padding: 0 12px">
        <a-card title="快速开始" style="margin-bottom: 24px" :bordered="false" :body-style="{padding: '0 10px 10px 0'}">
          <div class="item-group">
            <a v-for="(item, index) in startLinks" :key="index" href="javascript:;" @click="$router.push({name: item.name})">{{ item.title }}</a>
            <a-button size="small" type="primary" ghost icon="plus" @click="$refs.linkModal.add()">添加</a-button>
          </div>
        </a-card>
        <a-card title="今日提醒" style="margin-bottom: 24px" :bordered="false" :body-style="{padding: '15px'}">
          <li class="ant-list-item list_item" v-for="(item, index) in todayReminds" :key="index" @click="$router.push(item.url)">
            <div class="ant-list-item-main">{{ item.label }}</div>
            <div class="ant-list-item-extra">
              <a-badge slot="extra" v-if="item.tag && ~~item.tag !== 0" :count="item.tag" ></a-badge>
              <span v-else class="ant-badge ant-badge-not-a-wrapper" slot="extra">
                <sup title="3" class="ant-scroll-number ant-badge-count" data-show="true" style="background-color: #47c479!important;">
                  <span class="ant-scroll-number-only" >
                    <p class="current">{{ item.tag }}</p>
                  </span></sup>
              </span>

            </div>
          </li>
        </a-card>
        <!-- <a-card id="contact_dev" title="技术支持" style="margin-bottom: 24px" :bordered="false" :body-style="{padding: '15px'}">
          <div class="textList">
            <label>服务热线:</label>
            <span>{{ hotLine }}</span>
          </div>
          <div class="textList">
            <label>意见反馈</label>
            <a href="javascript:;" @click="$router.push({name: 'WorkOrder'})">立即反馈</a>
          </div>
        </a-card> -->
        <!-- <a-card :loading="loading" title="移动管理平台" :bordered="false" :body-style="{padding: '10px 10px 10px 18px'}">
          <div class="ant-list ant-list-vertical ant-list-split">
            <div class="ant-spin-nested-loading">
              <div class="ant-spin-container">
                <li class="ant-list-item ant-list-item-no-flex" style="padding: 0;">
                  <div class="oa_qrcode">
                    <img src="http://pms-static.gmtech.top/202016/1578279464036220078.png">
                    <div>
                      <a target="_blank" href="https://work.weixin.qq.com/api/doc#90001/90143/90578">申请企业微信，与微信互联<br>体验更加强大的移动管理功能</a>
                    </div>
                  </div>
                </li>
              </div>
            </div>
          </div>
        </a-card> -->
      </a-col>
    </a-row>

    <edit-link ref="linkModal" @ok="getStartLinks"/>
  </page-view>
</template>

<script>
import Vue from 'vue'
import moment from 'moment'
import { mapState } from 'vuex'
import { PageView } from '@/layouts'
import editLink from './modules/editLink'
import { arr2obj } from '@/utils/index'
import { startLinks } from '@/utils/const'
// import { getSystemLog } from '@/api/system/system'
import HeadInfo from '@/components/tools/HeadInfo'
import { dashboardInfo, homepageBaseinfo } from '@/api/system/dashboard'
import steps from './steps'

export default {
  name: 'Workplace',
  components: {
    PageView,
    HeadInfo,
    editLink
  },
  data () {
    // this.chartLineSettings = {
    //   legendName: {
    //     'num': '消耗数量'
    //   }
    // }
    // this.grid = {
    //   show: true,
    //   top: 35,
    //   bottom: 10,
    //   left: 10,
    //   right: 10
    // }
    return {
      avatar: '',
      hotLine: '',
      todayNotice: {},
      user: {},
      baseInfo: {
        project_cnt: '-',
        group_cnt: '-',
        building_cnt: '-',
        room_cnt: '-',
        user_cnt: '-',
        wechat_cnt: '-',
        parking_cnt: '-',
        employee_cnt: '-',
        update_time: null
      },
      loading: true,
      logs: [], // 操作日志
      startLinks: [], // 快速开始
      todayReminds: [
        { label: '今日投诉:', tag: 0, url: '/repair/complaint/list?date=today' },
        { label: '今日工单:', tag: 0, url: '/repair/ticket/list?date=today' },
        { label: '超时工单:', tag: 0, url: '/repair/ticket/list?isovertime=1' }
      ]
      // chartData: {
      //   columns: ['item', 'count'],
      //   rows: [
      //     { 'item': '已收', 'count': 12832, 'percent': 0.3536 },
      //     { 'item': '欠费', 'count': 23452, 'percent': 0.6464 }
      //   ]
      // },
      // chartLineData: {
      //   columns: ['日期', '消耗短信'],
      //   rows: []
      // },
      // chartSettings: {
      //   radius: [75, 90],
      //   offsetY: 180
      // },
      // columns: [
      //   {
      //     title: '收费项目',
      //     dataIndex: 'fee_name'
      //   },
      //   {
      //     title: '应收',
      //     dataIndex: 'money'
      //   },
      //   {
      //     title: '实收',
      //     dataIndex: 'payed'
      //   },
      //   {
      //     title: '欠费',
      //     dataIndex: 'lack'
      //   }
      // ],
      // feeData: [
      //   { id: 1, 'fee_name': '房租', 'money': 8836.48, 'payed': 36.48, 'lack': 8800 },
      //   { id: 2, 'fee_name': '车位管理费', 'money': 8836.48, 'payed': 36.48, 'lack': 8800 },
      //   { id: 3, 'fee_name': '物业费', 'money': 8836.48, 'payed': 36.48, 'lack': 8800 },
      //   { id: 4, 'fee_name': '住宿费', 'money': 8836.48, 'payed': 36.48, 'lack': 8800 },
      //   { id: 5, 'fee_name': '维修费', 'money': 8836.48, 'payed': 36.48, 'lack': 8800 }
      // ]
    }
  },
  computed: {
    ...mapState({
      nickname: (state) => state.user.nickname,
      welcome: (state) => state.user.welcome
    }),
    userInfo () {
      return this.$store.getters.userInfo
    },
    projects () {
      return [
        {
          cover: '/images/project.png',
          title: '项目',
          id: 'add_project',
          extra: `${this.baseInfo.project_cnt}个`,
          routerName: 'BaseProject',
          links: [
            // {
            //   label: '添加项目',
            //   routerName: 'BaseProjectOperationAdd'
            // }
          ]
        },
        {
          cover: '/images/group.png',
          title: '组团',
          id: 'add_group',
          extra: `${this.baseInfo.group_cnt}个`,
          routerName: 'BaseGroup',
          links: [
            // {
            //   label: '添加组团',
            //   routerName: 'BaseGroupOperationAdd'
            // }
          ]
        },
        {
          cover: '/images/building.png',
          title: '楼宇',
          id: 'add_building',
          extra: `${this.baseInfo.building_cnt}个`,
          routerName: 'BaseBuilding',
          links: [
            // {
            //   label: '添加楼宇',
            //   routerName: 'BaseBuildingOperationAdd'
            // }
          ]
        },
        {
          cover: '/images/room.png',
          title: '房屋',
          id: 'add_room',
          extra: `${this.baseInfo.room_cnt}个`,
          routerName: 'BaseRoom',
          links: [
            // {
            //   label: '添加房屋',
            //   routerName: 'BaseRoomOperationAdd'
            // },
            // {
            //   label: '导入房屋',
            //   routerName: 'BaseRoomDataImport'
            // }
          ]
        },
        {
          cover: '/images/person.png',
          title: '住户',
          extra: `${this.baseInfo.user_cnt}个`,
          routerName: 'BasePersonList',
          links: [
            // {
            //   label: '添加住户',
            //   routerName: 'BasePersonAdd'
            // },
            // {
            //   label: '导入住户',
            //   routerName: 'BasePersonDataImport'
            // }
          ]
        },
        {
          cover: '/images/wechat.png',
          title: '公众号',
          extra: `${this.baseInfo.wechat_cnt}个`,
          routerName: '',
          links: [
            // {
            //   label: '添加公众号',
            //   routerName: 'WechatMpSetting'
            // }
          ]
        },
        {
          cover: '/images/parking.png',
          title: '车位',
          extra: `${this.baseInfo.parking_cnt}个`,
          routerName: 'BaseParking',
          links: [
            // {
            //   label: '添加车位',
            //   routerName: 'BaseParkingOperationAdd'
            // },
            // {
            //   label: '导入车位',
            //   routerName: 'BaseParkingDataImport'
            // }
          ]
        },
        {
          cover: '/images/staff.png',
          title: '公司',
          extra: `${this.baseInfo.employee_cnt}个`,
          routerName: '',
          links: [
            {
              label: '添加公司',
              routerName: 'companyCreate'
            }
            // ,
            // {
            //   label: '权限管理',
            //   routerName: 'MobileCas'
            // }
          ]
        }
      ]
    }
  },
  created () {
    this.getStartLinks()
    this.getBaseInfo()
    this.user = this.userInfo
    this.avatar = this.userInfo.avatar

    this.getHotLine()
  },
  mounted () {
    // this.getLog()
    if (!Vue.ls.get('analysis_guide')) {
      this.startSteps()
    }
  },
  // 这个必须有 否则分步引导的浮层在没关闭的情况下会一直存在
  beforeRouteLeave (to, from, next) {
    this.$driver.reset()
    next()
  },
  methods: {
    startSteps () {
      setTimeout(() => {
        if (document.getElementById('add_project')) {
          this.$driver.defineSteps(steps)
          this.$driver.start()
          this.initEvent()
        } else {
          this.startSteps()
        }
      }, 500)
    },
    initEvent () {
      document.addEventListener('click', (event) => {
        if (event.target.className === 'driver-close-btn') {
          Vue.ls.set('analysis_guide', true)
        }
      })
    },
    moment,
    getStartLinks (value) {
      let data = Vue.ls.get('startLink')
      // 编辑后的数据
      if (value) { data = value }
      data = data ? this.formatLinks(data) : startLinks.filter((item, index) => index < 7)

      this.startLinks = data
      Vue.ls.set('startLink', data)
    },
    formatLinks (list) {
      const obj = arr2obj(startLinks, 'title', 'name')

      return list.map(item => {
        item.name = obj[item.title]
        return item
      })
    },
    // getLog () {
    //   getSystemLog({ page_size: 10 }).then(res => {
    //     this.loading = false
    //     this.logs = res.data.list || []
    //   })
    // },
    getBaseInfo () {
      homepageBaseinfo().then(res => {
        if (res.code === 200) {
          this.baseInfo = res.data || {}
        }
      })
    },
    async getHotLine () {
      const res = await dashboardInfo()
      if (res && res.code === 200) {
        this.hotLine = res.data.hotline || ''
        const todayNotice = res.data.today_notice
        // let smsStatics = []
        // if (res.data.sms_statics && res.data.sms_statics.length) {
        //   smsStatics = res.data.sms_statics.map(item => {
        //     return { '日期': item.dt.toString(), '消耗短信': item.num }
        //   })
        // }
        this.todayReminds = [
          { label: '今日投诉:', tag: todayNotice.complaint_order_num, url: '/repair/complaint/list?date=today' },
          { label: '今日工单:', tag: todayNotice.dispatch_order_num, url: '/repair/ticket/list?date=today' },
          { label: '超时工单:', tag: todayNotice.dispatch_order_overtime_num, url: '/repair/ticket/list?isovertime=1' }
        ]
        // this.chartLineData.rows = smsStatics
      }
    }
  }
}
</script>

<style lang="less" scoped>
  .project-list {

    .card-title {
      font-size: 0;
      display: flex;

      a {
        color: rgba(0, 0, 0, 0.85);
        margin-left: 12px;
        line-height: 24px;
        height: 24px;
        display: inline-block;
        vertical-align: top;
        font-size: 14px;

        &:hover {
          color: #1890ff;
        }
      }
    }
    .card-description {
      color: rgba(0, 0, 0, 0.45);
      height: 44px;
      line-height: 22px;
      overflow: hidden;
    }
    .project-item {
      display: flex;
      justify-content: flex-end;
      margin-top: 18px;
      overflow: hidden;
      font-size: 12px;
      height: 20px;
      line-height: 20px;
      a {
        color: rgba(0, 0, 0, 0.45);
        display: inline-block;
        margin-right: 8px;
        overflow: hidden;
        text-overflow: ellipsis;
        word-break: break-all;
        white-space: nowrap;

        &:hover {
          color: #1890ff;
        }
      }
      .datetime {
        color: rgba(0, 0, 0, 0.25);
        flex: 0 0 auto;
        float: right;
      }
    }
    .ant-card-meta-description {
      color: rgba(0, 0, 0, 0.45);
      height: 44px;
      line-height: 22px;
      overflow: hidden;
    }
  }
  .item-group {
    padding: 20px 0 8px 24px;
    font-size: 0;
    a {
      color: rgba(0, 0, 0, 0.65);
      display: inline-block;
      font-size: 14px;
      margin-bottom: 13px;
      margin-right: 5px;
    }
  }
  .members {
    a {
      display: block;
      margin: 12px 0;
      line-height: 24px;
      height: 24px;
      .member {
        font-size: 14px;
        color: rgba(0, 0, 0, .65);
        line-height: 24px;
        max-width: 100px;
        vertical-align: top;
        margin-left: 12px;
        transition: all 0.3s;
        display: inline-block;
      }
      &:hover {
        span {
          color: #1890ff;
        }
      }
    }
  }
  .mobile {

    .project-list {

      .project-card-grid {
        width: 100%;
      }
    }

    .more-info {
      border: 0;
      padding-top: 16px;
      margin: 16px 0 16px;
    }

    .headerContent .title .welcome-text {
      display: none;
    }
  }
  .m-l-sm {
    margin-left: 8px!important;
  }
  .m-l-md {
    margin-left: 16px!important;
  }
  .list_item {
    border-bottom: none;
    padding: 5px 12px;
    cursor: pointer;
    margin: 0
  }
  .textList {
    padding: 0 10px 10px 12px;
    label {
      margin-right: 8px;
    }
  }
  .oa_qrcode {
    display: flex;
    align-items: center;
    img {
      display: inline-block;
      width: 100px;
      height: 100px;
      margin-right: 15px;
    }
    div {
      line-height: 30px;
      flex: 1;
    }
  }
</style>
