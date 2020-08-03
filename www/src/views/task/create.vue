<template>
  <a-card :body-style="{padding: '24px 32px'}" :bordered="false">
    <a slot="title" >
      选择项目
    </a>
    <a slot="extra" >
      <a-button @click="$router.go(-1)">返回</a-button>
    </a>
    <a-collapse accordion>
      <a-collapse-panel v-for="item in environments" :key="item.id" :header="item.env_name">
        <a-button type="primary" v-for="project in filProjects(item.id)" @click="$router.push(`/task/create/${project.id}`)" :key="project.id" style="margin:0 32px 18px 0">
          {{ project.name }}
        </a-button>
      </a-collapse-panel>
    </a-collapse>
  </a-card>
</template>

<script>

import { getUsers } from '@/api/user'
// import { updateResource } from '@/api/resource'
import { fetchPostSpace, fetchDelSpaceId } from '@/api/space'
// import { apps } from '@/mixin/apps'
import { getEnvironments } from '@/api/environment'
import { getProjects } from '@/api/project'

export default {
  name: 'ResourceAddOrEdit',
  // mixins: [apps],
  data () {
    return {
      form: this.$form.createForm(this),
      editId: this.$route.params.id || 'add',
      isAdd: true,
      pidOptions: [],
      projects: [],
      environments: [],
      text: 'sdsd',
      data: [],
      saveLoading: false
    }
  },
  created () {
    this.getProjects()
    this.getEnvironments()
  },
  methods: {
    getEnvironments () {
      getEnvironments().then(res => {
        console.log(res)
        this.environments = res.data.list
      })
    },
    getProjects () {
      getProjects({ page: 1, size: 1000 }).then(res => {
        this.projects = res.data.list
      })
    },
    filProjects (id) {
      return this.projects.filter(item => item.environment_id === id)
    },
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
