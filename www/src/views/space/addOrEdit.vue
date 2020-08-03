<template>
  <a-card :body-style="{padding: '24px 32px'}" :bordered="false">
    <a-form
      style="margin: 0 auto;"
      @submit="handleSubmit"
      :form="form"
      :labelCol="{lg: {span: 7}, sm: {span: 7}}"
      :wrapperCol="{lg: {span: 10}, sm: {span: 17} }">

      <a-form-item label="空间名称">
        <a-input
          v-decorator="['name', {rules: [{required: true, message: '请输入空间名称'}, {max: 255, message: '空间名称超出长度限制'}]}]"
          placeholder="请输入空间名称"></a-input>
      </a-form-item>

      <a-form-item label="所属人">
        <a-select
          show-search
          placeholder="请分配用户"
          v-decorator="['user_id', {rules: [{required: true, message: '请分配用户'}]}]"
          labelInValue
          @search="handleSearch"
        >
          <a-select-option v-for="d in data" :key="d.user_id">
            {{ d.username }}
          </a-select-option>
        </a-select>
        <!-- <a-input
          :disabled="this.editId !== 'add'"
          v-decorator="['code', {rules: [{required: true, message: '请输入所属人'}, {max: 255, message: '所属人超出长度限制'}]}]"
          placeholder="请输入所属人"></a-input> -->
      </a-form-item>

      <a-form-item :wrapper-col="{ span: 12, offset: 7 }">
        <a-button :loading="saveLoading" htmlType="submit" type="primary">保存</a-button>
        <a-button style="margin-left: 16px" @click="cancelInfo">返回</a-button>
      </a-form-item>
    </a-form>
  </a-card>
</template>

<script>

import { getUsers } from '@/api/user'
// import { updateResource } from '@/api/resource'
import { fetchPostSpace, fetchDelSpaceId } from '@/api/space'
// import { apps } from '@/mixin/apps'

export default {
  name: 'ResourceAddOrEdit',
  // mixins: [apps],
  data () {
    return {
      form: this.$form.createForm(this),
      editId: this.$route.params.id || 'add',
      isAdd: true,
      pidOptions: [],
      data: [],
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
    handleSearch (value) {
      if (!value) return
      const params = {
        kw: value
      }
      getUsers(params).then(res => {
        console.log(res)
        this.data = res.data.list
      })
    },
    handleChange (selectedItems) {
      console.log(selectedItems)
      this.selectedItems = selectedItems
    },
    /**
     * 编辑页需要获取当前编辑的资源详细信息，这里需要请
     * 求接口获取数据，并回填到表单里边
     */
    getItemDetail () {
      fetchDelSpaceId({ id: this.editId }).then(res => {
        if (res.code === 0) {
          const data = res.data
          this.data = data.members
          this.form.setFieldsValue({
            name: data.name,
            user_id: {
              key: data.user_id,
              label: data.members[0].user_name
            }
          })
        } else {
          this.$message.error(res.msg)
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
          ...values,
          user_id: values.user_id.key,
          user_name: values.user_id.label
        }
        this.saveLoading = true
        if (this.isAdd) {
          try {
            fetchPostSpace(data).then(res => {
              this.saveLoading = false
              if (res.code === 0) {
                this.$message.success(`添加成功`)
                this.$router.back()
              } else {
                this.$message.error(res.msg)
              }
            }).catch(() => {
              this.saveLoading = false
            })
          } catch (error) {
            console.log(error)
          }
        } else {
          if (data.pid.length) {
            data.pid = data.pid[data.pid.length - 1]
          }
          console.log(data)

          // updateResource({ ...data, id: ~~this.editId }).then(res => {
          //   this.saveLoading = false
          //   if (res.code === 0) {
          //     this.$message.success('更新成功')
          //     this.$router.back()
          //   } else {
          //     this.$message.error(res.msg)
          //   }
          // }).catch(() => {
          //   this.saveLoading = false
          // })
        }
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
