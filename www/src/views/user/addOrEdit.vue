<template>
  <a-card :body-style="{padding: '24px 32px'}" :bordered="false">
    <a-form
      style="margin: 0 auto;"
      @submit="handleSubmit"
      :form="form"
      :labelCol="{lg: {span: 7}, sm: {span: 7}}"
      :wrapperCol="{lg: {span: 10}, sm: {span: 17} }">
      <a-form-item label="用户名">
        <a-input
          v-decorator="['username', {rules: [{required: true, message: '请输入用户名'}, {max: 255, message: '用户名超出长度限制'}]}]"
          placeholder="请输入用户名"></a-input>
      </a-form-item>
      <a-form-item label="邮箱">
        <a-input
          :disabled="this.editId !== 'add'"
          v-decorator="['email', {rules: [{required: true, message: '请输入邮箱'}, {type: 'email', message: '邮箱格式不正确'}]}]"
          placeholder="请输入邮箱"></a-input>
      </a-form-item>
      <a-form-item label="密码">
        <a-input
          type="password"
          v-decorator="['password', {rules: [{required: true, message: '请输入密码'},
                                             {pattern: /^(?=.*[a-z])(?=.*\d)[a-zA-Z\d]{6,}$/,message:'密码至少6个字符，至少1个大写字母，1个小写字母，1个数字'}]}]"
          placeholder="请输入密码"></a-input>
      </a-form-item>
      <a-form-item :wrapper-col="{ span: 12, offset: 7 }">
        <a-button :loading="saveLoading" htmlType="submit" type="primary">保存</a-button>
        <a-button style="margin-left: 16px" @click="cancelInfo">返回</a-button>
      </a-form-item>
    </a-form>
  </a-card>
</template>

<script>
import { fetchCasSsoCompanyList, fetchCasSsoStaffList } from '@/api/sso'
import { getPostUsers, fetchUserGet, fetchCasUserRoleSet } from '@/api/user'

export default {
  name: 'UserAddRoleOrEdit',
  // mixins: [apps],
  data () {
    this.lastFetchId = 0
    return {
      form: this.$form.createForm(this),
      editId: this.$route.params.id || 'add',
      isAdd: true,
      roles: [],
      companys: [],
      users: [],
      data: [],
      value: [],
      fetching: false,
      company_id: null,
      saveLoading: false
    }
  },
  created () {
    // editId如果是add则认为是添加页面，否则认为是编辑页面
    if (this.editId === 'add') {
      this.isAdd = true
    } else {
      this.isAdd = false
      this.getItemDetail()
    }
  },
  methods: {
    /**
     * 编辑页需要获取当前编辑的资源详细信息，这里需要请
     * 求接口获取数据，并回填到表单里边
     */
    getItemDetail () {
      fetchUserGet({ id: this.editId }).then(res => {
        if (res.code === 0) {
          const data = res.data
          this.form.setFieldsValue({
            username: data.username,
            email: data.email
          })

          // this.fetchUser(data.user_info.name)
        } else {
          this.$message.error(res.message)
        }
      })
    },

    /**
     * 检查表单内容并提交动作
     * @param e
     */
    handleSubmit (e) {
      e.preventDefault()
      this.form.validateFieldsAndScroll((err, values) => {
        if (err) {
          return
        }
        const data = {
          ...values
        }
        this.saveLoading = true
        if (this.isAdd) {
          getPostUsers(data).then(res => {
            this.saveLoading = false
            if (res.code === 0) {
              this.$message.success(`添加成功`)
              this.$router.back()
              this.handleReset()
            } else {
              this.$message.error(res.message)
            }
          }).catch(() => {
            this.saveLoading = false
          })
        } else {
          const params = {
            user_id: ~~this.editId,
            role_id_list: data.role_id_list
          }
          fetchCasUserRoleSet(params).then(res => {
            this.saveLoading = false
            if (res.code === 0) {
              this.$message.success('更新成功')
              this.$router.back()
            } else {
              this.$message.error(res.message)
            }
          }).catch(() => {
            this.saveLoading = false
          })
        }
      })
    },
    companyChange (value) {
      this.company_id = value
      this.fetchUser()
    },
    // handleSearch (value) {
    //   this.company_name = value
    //   this.fetchCompany()
    // },
    // getCompany () {
    //   const params = { page: 1, page_size: 20, name: this.company_name }
    //   fetchCasSsoCompanyList(params).then(res => {
    //     this.companys = res.data.list
    //     console.log(this.companys)
    //   })
    // },
    fetchCompany (value) {
      this.lastFetchId += 1
      const fetchId = this.lastFetchId
      this.data = []
      this.fetching = true
      const params = {
        page: 1,
        page_size: 100,
        // app_id: getAppIdToken(),
        name: value
      }
      fetchCasSsoCompanyList(params).then(body => {
        if (fetchId !== this.lastFetchId) {
          return
        }
        const data = body.data.list
        this.companys = data
        this.fetching = false
      })
    },
    fetchUser (value) {
      console.log(value)
      this.lastFetchId += 1
      const fetchId = this.lastFetchId
      this.data = []
      this.fetching = true
      const params = {
        company_id: this.company_id,
        page: 1,
        page_size: 100,
        name: value
      }
      fetchCasSsoStaffList(params).then(body => {
        if (fetchId !== this.lastFetchId) {
          return
        }
        const data = body.data.list
        this.data = data
        this.fetching = false
      })
    },
    handleCompanyChange (companyId) {
      Object.assign(this, {
        company_id: companyId,
        fetching: false
      })
      this.fetchUser()
    },
    handleChange (value) {
      Object.assign(this, {
        value,
        data: [],
        fetching: false
      })
    },
    /**
     * 重置操作
     */
    handleReset () {
      this.form.resetFields(['name'])
    },

    /**
     * 取消按钮点击事件
     */
    cancelInfo () {
      this.$router.back()
    }
  }
}
</script>
