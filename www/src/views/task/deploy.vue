<template>
  <a-card :body-style="{padding: '24px 32px'}" :bordered="false" class="deploy">
    <div slot="title" >
      <span>{{ taskInfo.project_name }} / {{ taskInfo.name }}</span>
      <a-button type="primary" :style="{marginLeft:'18px'}" @click="deployTask" :loading="taskInfo.status === 3" :disabled="taskInfo.status !== 1">
        <span v-if="taskInfo.status === 3">部署中</span>
        <span v-else-if="taskInfo.status === 4">部署完成</span>
        <span v-else>开始部署</span>
      </a-button>
    </div>
    <a slot="extra" >
      <a-button @click="$router.go(-1)">返回</a-button>
    </a>
    <a-spin :spinning="spinning" tip="Loading...">
      <codemirror
        style="width:100%"
        class="codemirror"
        ref="cmEditor"
        :value="code"
        :options="cmOptions"
      />
    </a-spin>
  </a-card>
</template>

<script>
import { getTask } from '@/api/task'
// import AnsiUp from 'ansi_up'

export default {
  name: 'DeployTask',
  components: {
  },
  data () {
    return {
      form: this.$form.createForm(this),
      editId: this.$route.params.id || 'add',
      deployLoading: false,
      isAdd: true,
      spinning: false,
      taskInfo: {},
      detectionInfo: {},
      fileList: [],
      code: '',
      cmOptions: {
        tabSize: 4,
        styleActiveLine: true,
        lineNumbers: true,
        line: true,
        foldGutter: true,
        styleSelectedText: true,
        mode: 'shell',
        matchBrackets: true,
        showCursorWhenSelecting: true,
        theme: 'xq-dark',
        lineWrapping: true,
        hintOptions: {
          completeSingle: false
        }
      },
      ansiUp: null,
      saveLoading: false
    }
  },
  created () {
    this.getItemDetail()
    // this.ansiUp = new AnsiUp()
    this.$nextTick(() => {
      this.$socket.emit('open', { 'task': this.editId })

      this.$socket.emit('logs', { 'task': ~~this.editId })

      // this.term = new Terminal()
      // this.term.open(this.$refs['terminal'])
      // this.term._initialized = true
      this.spinning = false
      console.log('mounted is going on')
      // this.term.write('Hello from \x1B[1;3;31mxterm.js\x1B[0m $ ')
    })
  },
  sockets: {
    connect () {
      console.log('socket connected')
    },
    console (data) {
      if (data.data.cmd) {
        this.code += `\n\r[${data.data.user}@${data.data.host}]$ ${data.data.cmd}`
      }
      if (data.data.success) {
        this.code += `\n\r${data.data.success}`
      }
      if (data.data.error) {
        this.code += `\n\r${data.data.error}`
      }
      this.getItemDetail()
    },
    construct (data) {
      this.taskInfo = data.data
    }
  },
  methods: {
    /**
     * 编辑页需要获取当前编辑的应用详细信息，这里需要请
     * 求接口获取数据，并回填到表单里边
     */
    getItemDetail () {
      getTask({ id: this.editId }).then(res => {
        if (res.code === 0) {
          this.taskInfo = res.data
        } else {
          this.$message.error(res.msg)
        }
      })
    },
    deployTask () {
      this.$socket.emit('deploy', { 'task': this.editId })
      this.taskInfo.status = 3
      this.getItemDetail()
    }

  }
}
</script>

<style lang="less">
// .codemirror-box{
// }
.deploy .CodeMirror {
  border: 1px solid #eee;
  height: auto;
}
// .codemirror{
//   border: 1px solid #dcdfe6;
//   height: auto;
// }
</style>
