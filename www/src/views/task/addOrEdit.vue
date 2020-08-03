<template>
  <a-card :body-style="{padding: '24px 32px'}" :bordered="false">

    <a-form
      style="margin: 0 auto;"
      @submit="handleSubmit"
      :form="form"
      :labelCol="{lg: {span: 7}, sm: {span: 7}}"
      :wrapperCol="{lg: {span: 10}, sm: {span: 17} }">

      <a-form-item label="上线单标题">
        <a-input
          v-decorator="['name', {rules: [{required: true, message: '请输入上线单标题'}, {max: 255, message: '上线单标题超出长度限制'}]}]"
          placeholder="请输入上线单标题"></a-input>
      </a-form-item>

      <a-form-item label="选取分支">
        <a-select
          placeholder="请选择分支"
          :loading="branchLoading"
          :disabled="branchLoading"
          @change="handleBranch"
          v-decorator="['branch', {rules: [{required: true, message: '请选择分支'}]}]"
        >
          <a-select-option v-for="item in branchList" :key="item" :value="item">
            {{ item }}
          </a-select-option>
        </a-select>
      </a-form-item>
      <a-form-item label="选取版本">
        <a-select
          placeholder="请选择版本"
          :loading="commitLoading"
          :disabled="commitLoading"
          v-decorator="['commit_id', {rules: [{required: true, message: '请选择版本'}]}]"
        >
          <a-select-option v-for="item in commitList" :key="item.id" :value="item.id">
            {{ item.id }}#{{ item.message }}
          </a-select-option>
        </a-select>
      </a-form-item>
      <a-form-item label="选取服务器">
        <a-radio-group v-decorator="['server',{initialValue:'all'}]" @change="handleChange">
          <a-radio value="all">
            全量服务器上线
          </a-radio>
          <a-radio value="diy">
            自定义服务器上线
          </a-radio>
        </a-radio-group>
        <div v-if="form.getFieldValue('server') === 'diy'" class="server-box">
          <p v-for="(item,index) in servers" :key="item.id" style="height:30px;margin:0">
            <a-tag :closable="servers.length>1" @close="closeServer(index)" style="marginBottom:6px">
              {{ item.name }}
            </a-tag>
          </p>

        </div>
      </a-form-item>
      <a-form-item :wrapper-col="{ span: 12, offset: 7 }">
        <a-button :loading="saveLoading" htmlType="submit" type="primary">保存</a-button>
        <a-button style="margin-left: 16px" @click="cancelInfo">返回</a-button>
      </a-form-item>
    </a-form>
  </a-card>
</template>

<script>
import { postTask, getTask, putTask } from '@/api/task'
import { getProject } from '@/api/project'
export default {
  name: 'ResourceAddOrEdit',
  // mixins: [apps],
  data () {
    return {
      form: this.$form.createForm(this),
      editId: this.$route.params.id || 'add',
      pidOptions: [],
      data: [],
      servers: [],
      taskInfo: {},
      commitList: [],
      branchList: [],
      commitLoading: true,
      branchLoading: true,
      saveLoading: false
    }
  },

  created () {
    if (this.$route.name === 'taskEditCreate') {
      console.log('edit')
      getTask({ id: this.editId }).then(res => {
        const data = res.data
        this.projectInfo = data.project_info

        this.taskInfo = data
        this.$socket.emit('open', { 'project_id': data.project_id })
        this.$socket.emit('branches')
        this.servers = data.servers_info
        const server = data.servers_info.length === data.project_info.servers_info.length ? 'all' : 'diy'
        this.form.setFieldsValue({
          name: data.name,
          server
        })
      })
    } else {
      this.$nextTick(() => {
        this.getProjectInfo()
        this.$socket.emit('open', { 'project_id': ~~this.editId })
        this.$socket.emit('branches')
      })
    }
  },
  sockets: {
    connect () {
      console.log('socket connected')
    },
    branches (data) {
      this.branchLoading = false
      this.commitLoading = true
      this.branchList = data.data
      const branchId = this.taskInfo.branch ? this.taskInfo.branch : data.data[0]

      this.form.setFieldsValue({
        branch: branchId
      })

      this.$socket.emit('commits', { branch: branchId })
      // console.log('branches>>>>>', data)
    },
    commits (data) {
      this.commitList = data.data
      this.commitLoading = false
      const commitId = this.taskInfo.commit_id ? this.taskInfo.commit_id : data.data[0].id
      this.form.setFieldsValue({
        commit_id: commitId
      })
    }
  },
  methods: {
    handleChange () {
      this.servers = JSON.parse(JSON.stringify(this.projectInfo.servers_info))
    },
    closeServer (index) {
      this.servers.splice(index, 1)
    },
    handleBranch (value) {
      this.commitLoading = true
      this.$socket.emit('commits', { branch: value })
    },
    getProjectInfo () {
      getProject({ id: this.editId }).then(res => {
        this.projectInfo = res.data
        this.servers = JSON.parse(JSON.stringify(this.projectInfo.servers_info))
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

        data.file_transmission_mode = 0
        data.file_list = '*.log'
        data.project_id = this.editId
        if (data.server === 'all') {
          data.servers = this.projectInfo.server_ids
        } else {
          data.servers = this.servers.map(item => item.id).join(',')
        }
        delete data.server
        this.saveLoading = true
        try {
          if (this.$route.name === 'taskEditCreate') {
            data.project_id = this.taskInfo.project_id
            putTask({ ...data, id: this.taskInfo.id }).then(res => {
              if (res.code === 0) {
                this.$message.success(`编辑成功`)
                this.$router.push('/deploy/index')
              } else {
                this.$message.error(res.msg)
              }
            })
            return
          }
          data.project_id = this.editId
          postTask(data).then(res => {
            this.saveLoading = false
            if (res.code === 0) {
              this.$message.success(`添加成功`)
              this.$router.push('/deploy/index')
            } else {
              this.$message.error(res.msg)
            }
          }).catch(() => {
            this.saveLoading = false
          })
        } catch (error) {
          console.log(error)
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
<style  scoped>
.server-box{
    width: 300px;
    border: 1px dashed #ddd;
    border-radius: 4px;
    padding: 0 0 10px 12px;
    /* display: -ms-flexbox;
    display: flex;
    -ms-flex-direction: column;
    flex-direction: column; */
}
</style>
