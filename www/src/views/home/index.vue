<template>
  <page-view :avatar="avatar" :title="false">
    <div slot="headerContent">
      <div class="title">{{ timeFix }}，{{ user.name }}
        <!-- <span class="welcome-text">，{{ welcome }}</span> -->
      </div>
    </div>
    <div slot="extra">
      <a-row class="more-info">
        <a-col :span="8">
          <head-info title="项目数" :content="projectCount" :center="false" :bordered="false"/>
        </a-col>
        <a-col :span="8">
          <head-info title="用员数" :content="teamsCount" :center="false" />
        </a-col>
        <a-col :span="8">
          <head-info title="部署数" :content="taskCount" :center="false" :bordered="false"/>
        </a-col>

      </a-row>
    </div>

    <div>
      <a-row :gutter="24">
        <a-col :xl="24" :lg="24" :md="24" :sm="24" :xs="24">
          <a-card
            class="project-list"
            :loading="loading"
            style="margin-bottom: 24px;"
            :bordered="false"
            title="进行中的项目"
            :body-style="{ padding: 0 }">
            <router-link slot="extra" :to="`project/index`">全部项目</router-link>
            <div>
              <a-card-grid class="project-card-grid" :key="i" v-for="(item, i) in projects">
                <a-card :bordered="false" :body-style="{ padding: 0 }">
                  <a-card-meta>
                    <div slot="title" class="card-title">
                      <a-avatar size="small" :src="item.cover"/>
                      <router-link :to="`project/${item.id}`">{{ item.name }}</router-link>
                    </div>
                    <div slot="description" class="card-description">
                      {{ item.description }}
                    </div>
                  </a-card-meta>
                  <div class="project-item">
                    <router-link :to="`project/${item.id}`">{{ item.environment_name }}</router-link>
                    <span class="datetime">{{ item.updated_at }}</span>
                  </div>
                </a-card>
              </a-card-grid>
            </div>
          </a-card>

          <a-card :loading="loading" title="动态" :bordered="false">
            <a-list>
              <a-list-item :key="index" v-for="(item, index) in activities" @click="$router.push(`/task/deploy/${item.id}`)" :style="{cursor:'pointer'}">
                <a-list-item-meta>
                  <a-avatar slot="avatar" :src="item.avatar" />
                  <div slot="title">
                    <a href="#">{{ item.user_name }}</a>
                   &nbsp; &nbsp;

                    <span>{{ statusMemo[item.status] }}</span>
                   &nbsp; &nbsp;
                    <span>{{ item.name }}</span>
                    &nbsp;
                    &nbsp;
                    &nbsp;
                    <a href="#">{{ item.event }}</a>
                  </div>
                  <div slot="description">{{ item.updated_at }}</div>
                </a-list-item-meta>
              </a-list-item>
            </a-list>
          </a-card>
        </a-col>
      </a-row>
    </div>
  </page-view>
</template>

<script>
import { timeFix } from '@/utils/util'
import { mapState } from 'vuex'
import { PageView } from '@/layouts'
import HeadInfo from '@/components/tools/HeadInfo'
import { getProjects } from '@/api/project'
import { getTasks } from '@/api/task'
import { getUsers } from '@/api/user'

// import { getRoleList, getServiceList } from '@/api/manage'
// const DataSet = require('@antv/data-set')
export default {
  name: 'Workplace',
  components: {
    PageView,
    HeadInfo
  },
  data () {
    const statusMemo = {
      0: '新建提交',
      1: '审核通过',
      2: '审核拒绝',
      3: '上线中',
      4: '上线完成',
      5: '上线失败'
    }
    return {
      statusMemo,
      timeFix: timeFix(),
      avatar: '',
      user: {},
      projects: [],
      projectCount: '0',
      taskCount: '0',
      teamsCount: '0',
      loading: true,
      radarLoading: true,
      activities: [],
      teams: [],
      // data
      axis1Opts: {
        dataKey: 'item',
        line: null,
        tickLine: null,
        grid: {
          lineStyle: {
            lineDash: null
          },
          hideFirstLine: false
        }
      },
      axis2Opts: {
        dataKey: 'score',
        line: null,
        tickLine: null,
        grid: {
          type: 'polygon',
          lineStyle: {
            lineDash: null
          }
        }
      },
      scale: [{
        dataKey: 'score',
        min: 0,
        max: 80
      }],
      axisData: [
        { item: '引用', a: 70, b: 30, c: 40 },
        { item: '口碑', a: 60, b: 70, c: 40 },
        { item: '产量', a: 50, b: 60, c: 40 },
        { item: '贡献', a: 40, b: 50, c: 40 },
        { item: '热度', a: 60, b: 70, c: 40 },
        { item: '引用', a: 70, b: 50, c: 40 }
      ],
      radarData: []
    }
  },
  computed: {
    ...mapState({
      nickname: (state) => state.user.nickname,
      welcome: (state) => state.user.welcome
    }),
    userInfo () {
      return this.$store.getters.userInfo
    }
  },
  created () {
    this.user = this.userInfo
    this.avatar = this.userInfo.avatar
    // getRoleList().then(res => {
    //   // console.log('workplace -> call getRoleList()', res)
    // })
    // getServiceList().then(res => {
    //   // console.log('workplace -> call getServiceList()', res)
    // })
  },
  mounted () {
    this.getProjects()
    this.getActivity()
    this.getTeams()
    // this.initRadar()
  },
  methods: {
    getProjects () {
      getProjects({ page: 1, size: 6 }).then(res => {
        this.projects = res.data && res.data.list
        this.projectCount = res.data.count.toString()
        this.loading = false
      })
    },
    getActivity () {
      getTasks({ size: 10, page: 1 }).then(res => {
        this.taskCount = res.data.count.toString()
        this.activities = res.data.list.filter(item => [3, 4, 5].includes(item.status))
      })
    },
    getTeams () {
      getUsers({ page: 1, size: 1 }).then(res => {
        this.teamsCount = res.data.count.toString()
      })
    },
    initRadar () {
      this.radarLoading = false
      // this.$http.get('/workplace/radar')
      //   .then(res => {
      //     const dv = new DataSet.View().source(res.result)
      //     dv.transform({
      //       type: 'fold',
      //       fields: ['个人', '团队', '部门'],
      //       key: 'user',
      //       value: 'score'
      //     })
      //     this.radarData = dv.rows
      //     this.radarLoading = false
      //   })
    }
  }
}
</script>

<style lang="less" scoped>
  .project-list {
    .card-title {
      font-size: 0;
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
      margin-top: 8px;
      overflow: hidden;
      font-size: 12px;
      height: 20px;
      line-height: 20px;
      a {
        color: rgba(0, 0, 0, 0.45);
        display: inline-block;
        flex: 1 1 0;
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
      width: 25%;
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
</style>
