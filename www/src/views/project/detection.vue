<template>
  <a-card :body-style="{padding: '24px 32px'}" :bordered="false">
    <a slot="title" >
      检测报告：{{ projectInfo.name }}
    </a>
    <a slot="extra" >
      <a-button @click="$router.go(-1)">返回</a-button>
    </a>
    <a-spin :spinning="spinning" tip="Loading...">
      <div style="height:280px">
        <a-result
          v-if="detectionInfo.message"
          :status="detectionInfo.code === 0 ? 'success' : warning"
          :title="detectionInfo.message"
        >
        </a-result>
      </div>
    </a-spin>
  </a-card>
</template>

<script>
import { getProject, detectionProject } from '@/api/project'

export default {
  name: 'BaseDoorAddOrEdit',
  components: {
  },
  data () {
    return {
      form: this.$form.createForm(this),
      editId: this.$route.params.id || 'add',
      isAdd: true,
      spinning: true,
      projectInfo: {},
      detectionInfo: {},
      fileList: [],
      saveLoading: false
    }
  },
  created () {
    this.getItemDetail()
    this.detectionProject()
  },
  methods: {
    /**
     * 编辑页需要获取当前编辑的应用详细信息，这里需要请
     * 求接口获取数据，并回填到表单里边
     */
    getItemDetail () {
      getProject({ id: this.editId }).then(res => {
        if (res.code === 0) {
          this.projectInfo = res.data
        } else {
          this.$message.error(res.msg)
        }
      })
    },
    detectionProject () {
      //  Promise.all([detectionProject({ id: this.editId }), getProject({ id: this.editId })]).then(res => {
      //   if (res[1].code === 0) {
      //     this.projectInfo = res[1].data
      //   } else {
      //     this.$message.error(res.msg)
      //   }
      //   this.spinning = false
      //   this.detectionInfo = res[0]
      // })
      this.spinning = true
      detectionProject({ id: this.editId }).then(res => {
        this.detectionInfo = res
        this.spinning = false
      })
    }
  }
}
</script>
