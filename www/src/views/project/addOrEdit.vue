<template>
  <a-card :body-style="{padding: '24px 32px'}" :bordered="false" class="project">
    <a-form
      style="margin: 0 auto;"
      layout="vertical"
      @submit="handleSubmit"
      :form="form"
    >
      <a-card
        class="card"
        title="基本配置"
        :bordered="false"
        :headStyle="{padding:0}"
        size="small"
      >
        <a-row :gutter="100">
          <a-col :span="12">
            <a-form-item label="项目名称">
              <a-input
                v-decorator="['name', {rules: [{required: true, message: '请输入项目名称'}]}]"
                placeholder="请输入项目名称"></a-input>
            </a-form-item>
          </a-col>
          <a-col :span="12">
            <a-form-item label="环境">
              <a-select
                v-decorator="['environment_id', {rules: [{required: true, message: '请选择环境'}]}]"
                placeholder="请选择环境"
              >
                <a-select-option v-for="item in environments" :value="item.id" :key="item.id">
                  {{ item.env_name }}
                </a-select-option>
              </a-select>
            </a-form-item>
          </a-col>
          <a-col :span="12">
            <a-form-item label="Git Repo">
              <a-input
                v-decorator="['repo_url', {rules: [{required: true, message: '请输入项目Git Repo'}]}]"
                placeholder="请输入项目Git Repo"></a-input>
            </a-form-item>
          </a-col>
          <a-col :span="12">
            <a-form-item label="分支/tag">
              <a-select
                v-decorator="['repo_mode', {rules: [{required: true, message: '请选择环境'}]}]"
                placeholder="请选择环境"
              >
                <a-select-option value="tag"> tag</a-select-option>
                <a-select-option value="branch"> branch</a-select-option>
              </a-select>
            </a-form-item>
          </a-col>
        </a-row>
      </a-card>
      <a-card class="card" title="目标集群" :bordered="false" :headStyle="{padding:0}" size="small">
        <a-row :gutter="24">
          <a-col :span="24" :style="{margin:'12px 0 18px'}">
            <a-transfer
              :titles="['服务器', '目标集群']"
              :data-source="servers"
              :list-style="{
                width: '400px',
                height: '400px',
              }"
              :render="item => item.title"
              :target-keys="targetKeys"
              @change="handleChange"
            />
          </a-col>
          <a-col :span="8">
            <a-form-item>
              <span slot="label">
                <span>目标集群部署路径</span>
                <a-tooltip placement="topLeft" >
                  <template slot="title">
                    <span>一般为webroot，不能为已存在目录，walle会自行生成</span>
                  </template>
                  <a-icon type="exclamation-circle" theme="filled" :style="{cursor:'pointer',marginLeft:'6px'}"/>
                </a-tooltip>
              </span>

              <a-input
                v-decorator="['target_root', {rules: [{required: true, message: '一般为webroot，不能为已存在目录'}]}]"
                placeholder="一般为webroot，不能为已存在目录"></a-input>
            </a-form-item>
          </a-col>
          <a-col :span="8">
            <a-form-item>
              <span slot="label">
                <span>目标集群部署仓库</span>
                <a-tooltip placement="topLeft" >
                  <template slot="title">
                    <span>仓库存储每次发布的版本，用于储存、回滚等版本管理</span>
                  </template>
                  <a-icon type="exclamation-circle" theme="filled" :style="{cursor:'pointer',marginLeft:'6px'}"/>
                </a-tooltip>
              </span>
              <a-input
                v-decorator="['target_releases', {rules: [{required: false, message: '仓库存储每次发布的版本'}]}]"
                placeholder="仓库存储每次发布的版本"></a-input>
            </a-form-item>
          </a-col>
          <a-col :span="8">
            <a-form-item label="目标集群部署仓库版本保留数">
              <a-input
                v-decorator="['keep_version_num', {rules: [{required: false, message: ''}]}]"
                placeholder="版本保留是为了做历史版本回滚"></a-input>
            </a-form-item>
          </a-col>
        </a-row>
      </a-card>
      <a-card class="card" title="任务配置" :bordered="false" :headStyle="{padding:0}" size="small">
        <a-row :gutter="100">
          <a-col :span="12">
            <a-form-item>
              <span slot="label">
                <span>部署排除文件</span>
                <a-tooltip placement="topLeft" >
                  <template slot="title">
                    <span>排除不需要打包同步至服务器的文件或目录。一行一个，支持正则，如：*.log</span>
                  </template>
                  <a-icon type="exclamation-circle" theme="filled" :style="{cursor:'pointer',marginLeft:'6px'}"/>
                </a-tooltip>
              </span>
              <codemirror
                style="width:100%"
                class="codemirror"
                ref="cmEditor"
                v-decorator="['excludes', {rules: [{required: false, message: '请输入服务器ip'}]}]"
                :options="cmOptions"
                @ready="onCmReady"
                @focus="onCmFocus"
                @input="(code)=>onCmCodeChange(code,'excludes')"
              />
            </a-form-item>
          </a-col>
          <a-col :span="12">
            <a-form-item>
              <span slot="label">
                <span>自定义全局变量</span>
                <a-tooltip placement="topLeft" >
                  <template slot="title">
                    <span>自定义全局变量可在高级任务中使用，一行一个，格式：VAR=value。更多预置变量：http://walle-web.io/docs/configuration-project.html</span>
                  </template>
                  <a-icon type="exclamation-circle" theme="filled" :style="{cursor:'pointer',marginLeft:'6px'}"/>
                </a-tooltip>
              </span>
              <codemirror
                style="width:100%"
                class="codemirror"
                ref="cmEditor"
                v-decorator="['task_vars', {rules: [{required: false, message: '请输入服务器ip'}]}]"
                :options="cmOptions"
                @ready="onCmReady"
                @focus="onCmFocus"
                @input="(code)=>onCmCodeChange(code,'task_vars')"
              />
            </a-form-item>
          </a-col>

          <a-col :span="12">
            <a-form-item>
              <span slot="label">
                <span>高级任务-Deploy前置任务</span>
                <a-tooltip placement="topLeft" >
                  <template slot="title">
                    <span>在宿主机未检出代码前的前置任务，常为安装依赖、配置环境变量等</span>
                  </template>
                  <a-icon type="exclamation-circle" theme="filled" :style="{cursor:'pointer',marginLeft:'6px'}"/>
                </a-tooltip>
              </span>
              <codemirror
                style="width:100%"
                class="codemirror"
                ref="cmEditor"
                v-decorator="['prev_deploy', {rules: [{required: false, message: '请输入服务器ip'}]}]"
                :options="cmOptions"
                @ready="onCmReady"
                @focus="onCmFocus"
                @input="(code)=>onCmCodeChange(code,'prev_deploy')"
              />
            </a-form-item>
          </a-col>
          <a-col :span="12">
            <a-form-item>
              <span slot="label">
                <span>高级任务-Deploy后置任务</span>
                <a-tooltip placement="topLeft" >
                  <template slot="title">
                    <span>在宿主机检出代码后的后置任务，常为编译、清除文件</span>
                  </template>
                  <a-icon type="exclamation-circle" theme="filled" :style="{cursor:'pointer',marginLeft:'6px'}"/>
                </a-tooltip>
              </span>
              <codemirror
                style="width:100%"
                class="codemirror"
                ref="cmEditor"
                v-decorator="['post_deploy', {rules: [{required: false, message: '请输入服务器ip'}]}]"
                :options="cmOptions"
                @ready="onCmReady"
                @focus="onCmFocus"
                @input="(code)=>onCmCodeChange(code,'post_deploy')"
              />
            </a-form-item>
          </a-col>
          <a-col :span="12">
            <a-form-item>
              <span slot="label">
                <span>高级任务-Release前置任务</span>
                <a-tooltip placement="topLeft" >
                  <template slot="title">
                    <span>在目标服务器同步代码到版本库后，服务切换的前置任务，常为停服、摘机器等</span>
                  </template>
                  <a-icon type="exclamation-circle" theme="filled" :style="{cursor:'pointer',marginLeft:'6px'}"/>
                </a-tooltip>
              </span>
              <codemirror
                style="width:100%"
                class="codemirror"
                ref="cmEditor"
                v-decorator="['prev_release', {rules: [{required: false, message: '请输入服务器ip'}]}]"
                :options="cmOptions"
                @ready="onCmReady"
                @focus="onCmFocus"
                @input="(code)=>onCmCodeChange(code,'prev_release')"
              />
            </a-form-item>
          </a-col>
          <a-col :span="12">
            <a-form-item>
              <span slot="label">
                <span>高级任务-Release后置任务</span>
                <a-tooltip placement="topLeft" >
                  <template slot="title">
                    <span>在目标服务器新版本服务切换后的后置任务，常为启动服务、启动节点</span>
                  </template>
                  <a-icon type="exclamation-circle" theme="filled" :style="{cursor:'pointer',marginLeft:'6px'}"/>
                </a-tooltip>
              </span>
              <codemirror
                style="width:100%"
                class="codemirror"
                ref="cmEditor"
                v-decorator="['post_release', {rules: [{required: false, message: '请输入服务器ip'}]}]"
                :options="cmOptions"
                @ready="onCmReady"
                @focus="onCmFocus"
                @input="(code)=>onCmCodeChange(code,'post_release')"
              />
            </a-form-item>
          </a-col>

          <a-col :span="12">
            <a-form-item label="上线通知">
              <!-- <a-input
                v-decorator="['notice_type', {rules: [{required: true, message: '请输入服务器ip'}]}]"
                placeholder="请输入服务器ip"></a-input> -->
              <a-select
                v-decorator="['notice_type', {rules: [{required: false, message: '请选择环境'}]}]"
                placeholder="请选择环境"
              >
                <a-select-option value="">无</a-select-option>
                <a-select-option value="email">邮箱通知</a-select-option>
                <a-select-option value="dingding">钉钉通知</a-select-option>
              </a-select>
            </a-form-item>
          </a-col>
          <a-col :span="12">
            <a-form-item v-if="form.getFieldValue('notice_type')">
              <span slot="label">
                {{ form.getFieldValue('notice_type') === 'dingding'?'钉钉hook地址':'邮箱地址' }}
              </span>
              <a-input
                v-decorator="['notice_hook', {rules: [{required: true, message: '请输入内容'}]}]"
                placeholder="请输入邮箱地址，多条英文分号分隔"></a-input>
            </a-form-item>
          </a-col>
          <a-col :span="24">
            <a-form-item label="">

              <a-checkbox
                :checked="form.getFieldValue('task_audit')"
                v-decorator="['task_audit', {rules: [{required: false, message: '请输入内容'}]}]"
              >
                上线单是否开启审核
              </a-checkbox>
            </a-form-item>
          </a-col>

        </a-row>
      </a-card>
      <a-form-item >
        <a-button :loading="saveLoading" htmlType="submit" type="primary">保存</a-button>
        <a-button style="margin-left: 16px" @click="cancelInfo">返回</a-button>
      </a-form-item>
    </a-form>
  </a-card>
</template>

<script>

import { getServers } from '@/api/server'
import { getProject, putProject, postProject } from '@/api/project'
import { getEnvironments } from '@/api/environment'

export default {
  name: 'ProjectAddOrEdit',
  // mixins: [apps],
  data () {
    return {
      form: this.$form.createForm(this),
      editId: this.$route.params.id || 'add',
      isAdd: true,
      pidOptions: [],
      environments: [],
      servers: [],
      data: [],
      mockData: [],
      targetKeys: [],
      code: 'const a = 10',
      bash: {},
      projectInfo: {},
      cmOptions: {
        tabSize: 4,
        mode: 'base',
        theme: 'base16-dark',
        lineNumbers: true,
        line: true
      },
      saveLoading: false
    }
  },
  computed: {
    codemirror () {
      return this.$refs.cmEditor && this.$refs.cmEditor.codemirror
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
    // this.getMock()
    this.$nextTick(() => {
      console.log('the current CodeMirror instance object:', this.codemirror)
    })
    this.getEnvironments()
    this.getServers()
  },
  methods: {
    getEnvironments () {
      getEnvironments().then(res => {
        this.environments = res.data.list
      })
    },
    getServers () {
      getServers().then(res => {
        this.servers = res.data.list.map(item => {
          item.key = item.id.toString()
          item.title = item.name
          return item
        })
        // this.targetKeys = targetKeys

        // this.getMock(res.data.list)
      })
    },

    // renderItem (item) {
    //   const customLabel = (
    //     <span class="custom-item">
    //       {item.title} - {item.description}
    //     </span>
    //   )

    //   return {
    //     label: customLabel, // for displayed item
    //     value: item.title // for title and filter matching
    //   }
    // },
    handleChange (targetKeys, direction, moveKeys) {
      console.log(targetKeys, direction, moveKeys)
      this.targetKeys = targetKeys
    },

    /**
     * 编辑页需要获取当前编辑的资源详细信息，这里需要请
     * 求接口获取数据，并回填到表单里边
     */
    getItemDetail () {
      getProject({ id: this.editId }).then(res => {
        if (res.code === 0) {
          const data = res.data
          this.projectInfo = data
          const values = this.form.getFieldsValue()
          Object.keys(values).forEach(item => {
            values[item] = data[item]
            if (item === 'task_audit') {
              values[item] = data[item] === 1
            }
          })
          this.targetKeys = data.servers_info.map(item => item.id.toString())
          this.form.setFieldsValue(values)
          this.$nextTick(() => {
            this.form.setFieldsValue({
              notice_hook: data.notice_hook
            })
          })
          // this.form.setFieldsValue({
          //   name: data.name,
          //   repo_url: data.repo_url,
          //   post_deploy: data.post_deploy,
          //   post_release: data.post_release,
          //   prev_deploy: data.prev_deploy,
          //   prev_release: data.prev_release,
          //   task_vars: data.task_vars,
          //   repo_mode: data.repo_mode,
          //   target_releases: data.target_releases,
          //   target_root: data.target_root,
          //   notice_hook: data.notice_hook,
          //   notice_type: data.notice_type
          // })
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
          ...this.projectInfo,
          ...values,
          ...this.bash
        }
        data.server_ids = this.targetKeys.join(',')
        data.space_id = 1
        data.user_id = 11
        if (data.task_audit) {
          data.task_audit = 1
        } else {
          data.task_audit = 0
        }
        this.saveLoading = true
        if (this.isAdd) {
          try {
            postProject(data).then(res => {
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
          data.status = 1
          putProject({ ...data, id: ~~this.editId }).then(res => {
            this.saveLoading = false
            if (res.code === 0) {
              this.$message.success('更新成功')
              this.$router.back()
            } else {
              this.$message.error(res.msg)
            }
          }).catch(() => {
            this.saveLoading = false
          })
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
    },
    onCmReady (cm) {
      console.log('the editor is readied!', cm)
    },
    onCmFocus (cm) {
      console.log('the editor is focused!', cm)
    },
    onCmCodeChange (newCode, key) {
      this.bash[key] = newCode.toString()
      // this.form.setFieldsValue({ key: newCode })
    }
  }
}
</script>

<style lang="less">
// .codemirror-box{
// }
.project{
  .CodeMirror {
    border: 1px solid #eee;
    height: 100px;
  }
}

// .codemirror{
//   border: 1px solid #dcdfe6;
//   height: auto;
// }
</style>
