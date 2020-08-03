<template>
  <div>
    <a-card :bordered="false" :bodyStyle="{ padding: '0 32px', height: '100%' }" :headStyle="{'border-bottom': 'none'}">
      <div slot="title">
        <a-button type="primary" icon="plus" @click="$router.push('/task/create')">添加上线单</a-button>
      </div>
      <div slot="extra">
        <!-- <h4 style="display: inline-block;"> 应用：</h4>
        <a-select v-model="queryParam.appid" placeholder="请选择应用" style="width: 200px;margin:0 24px 0 0" >
          <a-select-option value="all">全部</a-select-option>
          <a-select-option v-for="item in apps" :key="item.id" :value="item.id">{{ item.name }}</a-select-option>
        </a-select> -->
        <a-input-search v-model="queryParam.kw" placeholder="请输入空间名称" style="width: 260px" enter-button @search="handleEdit" />
        <!--
        <a-input-search v-model="queryParam.kw" placeholder="请输入关键字" style="width: 260px" enter-button @search="handleEdit" />
        <a-tooltip placement="top">
          <template slot="title">
            <span>高级搜索</span>
          </template>
          <a-button type="primary" shape="circle" style="margin-left:18px;border-radius: 4px;" @click="open"><a-icon type="filter" :style="{ fontSize: '16px' }" /></a-button>
        </a-tooltip> -->
      </div>
      <a-table
        style="padding-top: 0;"
        :pagination="pagination"
        @change="onPageChanged"
        row-key="id"
        :loading="loading"
        :columns="columns"
        :data-source="data"
      >
        <span slot="create_time" slot-scope="text">
          {{ text | moment }}
        </span>
        <span slot="status" slot-scope="text">
          <span>{{ statusMemo[text] }}</span>
          <!-- <a-badge :status="text | statusTypeFilter" :text="" /> -->
        </span>
        <template slot="action" slot-scope="text, record">
          <a-dropdown>
            <a class="user-link" href="javascript:;">操作<a-icon type="down" /> </a>
            <a-menu slot="overlay">
              <a-menu-item key="view" v-if="[3,4,5].includes(record.status)" @click="$router.push(`/task/deploy/${record.id}`)">查看</a-menu-item>
              <a-menu-item key="edit" v-if="[0,2,5].includes(record.status)" @click="$router.push(`/task/edit/${record.id}`)">编辑</a-menu-item>
              <a-menu-item v-if="[0].includes(record.status) && record.enable_audit" key="audit" @click="auditItem(record)">审核</a-menu-item>
              <a-menu-item v-if="[1].includes(record.status)" key="deploy" @click="$router.push(`/task/deploy/${record.id}`)">上线</a-menu-item>
              <a-menu-item v-if="[0,1,2].includes(record.status)" key="disable" @click="delItem(record)">删除</a-menu-item>
              <a-menu-item v-if="[4].includes(record.status)" key="disable" @click="rollbackItem(record)">回滚</a-menu-item>
            </a-menu>
          </a-dropdown>
        </template>
      </a-table>
    </a-card>
    <!-- :confirm-loading="confirmLoading" -->

    <a-modal
      title="上线单审核"
      :visible="visible"
      @cancel="handleCancel"
    >
      <h3>审核是否通过？</h3>
      <div slot="footer">
        <a-button @click="rejectTask" type="danger" :style="{marginRight:'6px'}">
          驳回
        </a-button>
        <a-button @click="auditTask" type="primary" >
          通过
        </a-button>
      </div>

    </a-modal>
  </div>
</template>

<script>
// import SearchForm from './modules/SearchForm'
import { pagination } from '@/utils'
import { getTasks, delTask, rejectTask, auditTask, rollbackTask } from '@/api/task'
export default {
  name: 'ResourceList',
  // mixins: [apps],
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
      selectedRowKeys: [],
      mdl: {},
      // 高级搜索 展开/关闭
      advanced: false,
      pagination,
      loading: false,
      // 查询参数
      queryParam: {
      },
      // 表头
      columns: [
        {
          title: 'ID',
          dataIndex: 'id'
        },
        {
          title: '上线单',
          width: 200,
          dataIndex: 'name'
        },
        {
          title: '项目',
          dataIndex: 'project_name'
        },
        {
          title: '用户名',
          dataIndex: 'user_name'
        },
        {
          title: '版本',
          dataIndex: 'commit_id'
        },
        {
          title: '环境',
          dataIndex: 'environment_name'
        },
        {
          title: '更新时间',
          dataIndex: 'updated_at'
        },
        {
          title: '状态',
          dataIndex: 'status',
          width: 100,
          scopedSlots: { customRender: 'status' }

        },
        {
          title: '操作',
          dataIndex: 'action',
          width: 100,
          scopedSlots: { customRender: 'action' }
        }

      ],
      data: [],
      options: [],
      labelCol: {
        xs: { span: 24 },
        sm: { span: 7 }
      },
      wrapperCol: {
        xs: { span: 24 },
        sm: { span: 13 }
      },
      itemInfo: {},
      visible: false,
      confirmLoading: false,

      form: this.$form.createForm(this)
      // 加载数据方法 必须为 Promise 对象
    }
  },
  created () {
    this.loadDatas()
  },
  methods: {
    open () {
      this.visible = true
    },
    loadDatas () {
      this.loading = true
      const params = {
        ...this.queryParam,
        page: this.pagination.current,
        size: this.pagination.pageSize
      }
      getTasks(params).then(res => {
        this.data = res.data.list
        this.loading = false

        this.visible = false
        this.confirmLoading = false
        this.pagination.total = res.data.count
      })
    },
    rollbackItem (item) {
      const that = this
      this.$confirm({
        title: '确定回滚吗??',
        content: '',
        okText: '确定',
        cancelText: '取消',
        onOk () {
          rollbackTask({ id: item.id }).then(res => {
            if (res.code === 0) {
              that.$message.success('回滚成功')
              that.loadDatas()
            }
          })
        }
      })
    },
    rejectTask () {
      rejectTask({ id: this.itemInfo.id }).then(res => {
        if (res.code === 0) {
          this.$message.success('驳回成功')
          this.loadDatas()
          this.visible = false
        }
      })
    },
    auditTask () {
      auditTask({ id: this.itemInfo.id }).then(res => {
        if (res.code === 0) {
          this.$message.success('审核通过')
          this.loadDatas()
          this.visible = false
        }
      })
    },
    auditItem (item) {
      this.visible = true
      this.itemInfo = item
    },
    delItem (item) {
      const that = this
      this.$confirm({
        title: '确认要删除吗?',
        content: '删除后将无法恢复',
        okText: '确定',
        okType: 'danger',
        cancelText: '取消',
        onOk () {
          delTask({ id: item.id }).then(res => {
            if (res.code === 0) {
              that.$message.success('删除成功')
              that.loadDatas()
            }
          })
        },
        onCancel () {
        }
      })
    },
    // 当分页发生变化的时候，我们需要对pagination重新赋值并重新获取列表
    onPageChanged (pagination) {
      this.pagination = pagination
      this.$scrollTo(0, 800)
      this.loadDatas()
    },
    handleEdit (record) {
      this.pagination.current = 1
      this.loadDatas()
    },
    handleCancel (e) {
      this.visible = false
    }
  }
}
</script>
