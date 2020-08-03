<template>
  <div>
    <a-card :bordered="false" :bodyStyle="{ padding: '0 32px', height: '100%' }" :headStyle="{'border-bottom': 'none'}">
      <div slot="title">
        <a-button type="primary" icon="plus" @click="$router.push('/user/add')">添加用户</a-button>
      </div>
      <div slot="extra">
        <a-input-search v-model="queryParam.kw" placeholder="请输入用户名称" style="width: 260px" enter-button @search="handleEdit" />
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

        <span slot="roleList" slot-scope="text">
          <a-tag v-for="item in text" :key="item.id" color="#2db7f5" style="margin:3px">{{ item.name }}</a-tag>
        </span>
        <span slot="status" slot-scope="text">
          <a-badge :status="text | statusTypeFilter" :text="text | statusFilter" />
        </span>
        <template slot="action" slot-scope="text, record">
          <a-dropdown>
            <a class="user-link" href="javascript:;">操作<a-icon type="down" /> </a>
            <a-menu slot="overlay">
              <a-menu-item key="edit" @click="$router.push(`/user/${record.user_id}`)">编辑</a-menu-item>
              <a-menu-item key="disable" @click="delItem(record)">删除</a-menu-item>
            </a-menu>
          </a-dropdown>
        </template>
      </a-table>
      <!-- <search-form ref="searchModal" @ok="handleOk" /> -->
      <a-modal
        title="高级搜索"
        :width="520"
        :visible="visible"
        :confirm-loading="confirmLoading"
        cancel-text="关闭"
        @ok="handleSubmit"
        @cancel="handleCancel"
      >
        <a-spin :spinning="confirmLoading">
          <a-form :form="form">
            <a-form-item
              label="所属组团"
              :label-col="labelCol"
              :wrapper-col="wrapperCol"
            >
              <FwCascader :max-level="2" :input-value="form.getFieldValue('group_id')" v-decorator="['group_id']"></FwCascader>
            </a-form-item>
            <a-form-item
              label="时间"
              :label-col="labelCol"
              :wrapper-col="wrapperCol"
            >
              <a-range-picker v-decorator="['search_time']" @change="onChange" />
            </a-form-item>
            <a-form-item
              label="关键字"
              :label-col="labelCol"
              :wrapper-col="wrapperCol"
            >
              <a-input
                v-decorator="[
                  'name',
                ]"
                placeholder="小区/房号/收费项目/收费标准"
              />
            </a-form-item>
          </a-form>
        </a-spin>
        <template slot="footer">
          <a-button key="back" @click="handleCancel">取消</a-button>
          <a-button key="reset" @click="handleReset">重置</a-button>
          <a-button key="submit" type="primary" :loading="loading" @click="handleSubmit">
            搜索
          </a-button>
        </template>
      </a-modal>
    </a-card>
  </div>
</template>

<script>
// import SearchForm from './modules/SearchForm'
import { pagination } from '@/utils'
import { getUsers, delUser } from '@/api/user'
// import { apps } from '@/mixin/apps'

export default {
  name: 'UserList',
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
      // 表头
      columns: [
        {
          title: '用户名',
          dataIndex: 'username'
        },
        {
          title: '邮箱',
          dataIndex: 'email'
        },
        {
          title: '状态',
          dataIndex: 'status'
        },
        {
          title: '角色',
          dataIndex: 'role_list',
          scopedSlots: { customRender: 'roleList' }
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

      getUsers(params).then(res => {
        this.data = res.data.list
        this.loading = false

        this.visible = false
        this.confirmLoading = false
        this.pagination.total = res.data.count
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
          delUser({ id: item.user_id }).then(res => {
            if (res.code === 0) {
              that.$message.success('删除成功')
              that.loadDatas()
            } else {
              that.$message.success(res.message)
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
      console.log(11)
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
