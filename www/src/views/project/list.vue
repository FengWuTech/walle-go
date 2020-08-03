<template>
  <div>
    <a-card :bordered="false" :bodyStyle="{ padding: '0 32px', height: '100%' }" :headStyle="{'border-bottom': 'none'}">
      <div slot="title">
        <a-button type="primary" icon="plus" @click="$router.push('/project/add')">添加项目</a-button>
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
        row-key="id"
        :loading="loading"
        :columns="columns"
        :data-source="data"
        @change="onPageChanged"
      >
        <span slot="create_time" slot-scope="text">
          {{ text | moment }}
        </span>
        <span slot="status" slot-scope="text">
          {{ statusText[text] }}
        </span>
        <template slot="action" slot-scope="text, record">
          <a-dropdown>
            <a class="user-link" href="javascript:;">操作<a-icon type="down" /> </a>
            <a-menu slot="overlay">
              <a-menu-item key="edit" @click="$router.push(`/project/${record.id}`)">编辑</a-menu-item>
              <!-- <a-menu-item key="copy" @click="copyItem(record)">复制</a-menu-item> -->
              <!-- <a-menu-item key="detection" @click="$router.push(`/project/detection/${record.id}`)">检测</a-menu-item> -->
              <a-menu-item key="disable" @click="delItem(record)">删除</a-menu-item>
            </a-menu>
          </a-dropdown>
        </template>
      </a-table>
    </a-card>
  </div>
</template>

<script>
// import SearchForm from './modules/SearchForm'
import { pagination } from '@/utils'
import { getProjects, delProject, copyProject } from '@/api/project'

export default {
  name: 'ResourceList',
  // mixins: [apps],
  data () {
    return {
      selectedRowKeys: [],
      mdl: {},
      // 高级搜索 展开/关闭
      advanced: false,
      pagination,
      loading: false,
      // 查询参数
      queryParam: {
      },
      statusText: {
        1: '正常',
        0: '异常'
      },
      // 表头
      columns: [
        {
          title: '项目名',
          dataIndex: 'name'
        },
        {
          title: '状态',
          dataIndex: 'status',
          scopedSlots: { customRender: 'status' }
        },
        {
          title: '环境',
          dataIndex: 'environment_name'
        },
        {
          title: '空间',
          dataIndex: 'space_name'
        },
        {
          title: '操作',
          dataIndex: 'action',
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
      visible: false,
      confirmLoading: false,

      form: this.$form.createForm(this)
      // 加载数据方法 必须为 Promise 对象
    }
  },
  computed: {
    rowSelection () {
      return {
        onChange: (selectedRowKeys, selectedRows) => {
          console.log(`selectedRowKeys: ${selectedRowKeys}`, 'selectedRows: ', selectedRows)
        },
        getCheckboxProps: record => ({
          props: {
            disabled: record.name === 'Disabled User', // Column configuration not to be checked
            name: record.name
          }
        })
      }
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
      getProjects(params).then(res => {
        this.data = res.data.list
        this.loading = false

        this.visible = false
        this.confirmLoading = false
        this.pagination.total = res.data.count
      })
    },
    copyItem (item) {
      copyProject({ id: item.id }).then(res => {
        if (res.code === 0) {
          this.$message.success('删除成功')
          this.loadDatas()
        } else {
          this.$message.error(res.message)
        }
      })
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
          delProject({ id: item.id }).then(res => {
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
    onChange (value) {
      console.log(value)
    },
    handleSubmit () {
      const { form: { validateFields } } = this
      this.confirmLoading = true
      validateFields((errors, values) => {
        if (!errors) {
          this.queryParam = {
            name: values.name
          }
          if (Array.isArray(values.group_id)) {
            this.queryParam.group_id = values.group_id[values.group_id.length - 1]
          }
          if (Array.isArray(values.search_time)) {
            this.queryParam.start_time = values.search_time[0]
            this.queryParam.end_time = values.search_time[1]
          }
          this.loadDatas()
        } else {
          this.confirmLoading = false
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
      this.$emit('onEdit', record)
    },
    handleOk (e) {
      this.loading = true
      setTimeout(() => {
        this.visible = false
        this.loading = false
      }, 3000)
    },
    handleReset () {
      this.form.resetFields()
    },
    handleCancel (e) {
      this.visible = false
    }
  }
}
</script>
