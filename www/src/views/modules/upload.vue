<template>
  <div>
    <a-upload
      name="file"
      :disabled="disabled"
      :accept="accept"
      :before-upload="beforeUpload"
      :remove="removeFile"
      :file-list="fileItems"
      :listType="listType"
      @preview="handlePreview"
      class="upload-list-inline"
    >
      <div v-if="fileItems.length < 1 && listType === 'picture-card'">
        <a-icon type="plus" />
        <div class="ant-upload-text">上传图片</div>
      </div>

      <a-button v-if="listType !== 'picture-card'" :disabled="disabled" :icon="icon" :loading="loading">
        <slot name="upload-txt">上传图片</slot>
      </a-button>
    </a-upload>
    <a-modal :visible="previewVisible" :footer="null" @cancel="handleCancel" >
      <img alt="example" style="width: 100%" :src="previewImage" />
    </a-modal>
  </div>

</template>

<script>

import COS from 'cos-js-sdk-v5'
import { getTecentUploadToken } from '@/api/common'

export default {
  name: 'UploadComponent',
  props: {
    icon: { // icon
      type: String,
      default: () => 'upload'
    },
    accept: { // 允许文件类型
      type: String,
      default: () => '.jpg,.png,.gif,.xlsx'
    },
    listType: { // 允许文件类型
      type: String,
      default: () => 'picture'
    },
    single: { // 是否只单个文件上传
      type: Boolean,
      default: () => true
    },
    max: {
      type: Number,
      default: () => 10000
    },
    disabled: {
      type: Boolean,
      default: false
    },
    fileList: {
      type: Array,
      default () {
        return []
      }
    }
  },
  data () {
    return {
      file: null,
      fileItems: [],
      loading: false,
      previewImage: null,
      cos: null,
      previewVisible: false
    }
  },
  watch: {
    fileList () {
      this.fileItems = this.fileList
    }
  },
  created () {
    this.fileItems = this.fileList
  },
  methods: {
    handleCancel () {
      this.previewVisible = false
    },
    handlePreview (file) {
      this.previewImage = file.url || file.thumbUrl
      this.previewVisible = true
    },
    removeFile (file) {
      if (this.disabled) {
        return false
      }
      this.file = null
      let idx = -1
      for (let i = 0; i < this.fileItems.length; i++) {
        const item = this.fileItems[i]
        if (item === file) {
          idx = i
        }
      }
      if (idx !== -1) {
        this.fileItems.splice(idx, 1)
      }
      this.$emit('change', this.fileItems)
    },
    beforeUpload (file) {
      if (this.fileItems.length + 1 > this.max) {
        this.$message.info(`最多支持上传${this.max}个文件`)
        return
      }
      this.file = file
      this.handleUpload()
      return false
    },
    handleUpload () {
      this.loading = true
      const { file } = this
      const suffix = file.name.substr(file.name.lastIndexOf('.') + 1)

      getTecentUploadToken({ suffix }).then(res => {
        res = res.data || {}
        this.cos = new COS({
          getAuthorization: (opt, callback) => {
            const data = {
              TmpSecretId: res.token.TmpSecretId,
              TmpSecretKey: res.token.TmpSecretKey,
              XCosSecurityToken: res.token.Token,
              StartTime: res.start_time,
              ExpiredTime: res.expired_time
            }
            callback(data)
          }
        })
        this.startUpload(res)
      })
    },
    startUpload  (res) {
      const that = this
      this.cos.putObject({
        Bucket: res.bucket || 'fpms-1301854395',
        Region: res.region || 'ap-nanjing',
        Key: res.filekey,
        Body: this.file,
        onProgress: function (progress) {
          if (progress && progress.percent === 1) {
            const url = `${res.url}/${res.filekey}`
            that.uploadComplete({
              url,
              status: 'done',
              name: that.file.name
            })
          }
        }
      }, function (err) {
        if (err) {
          that.uploadError(err)
        }
      })
    },
    uploadError (err) {
      this.loading = false
      console.log('err', err)
      this.$message.error('文件上传失败，请重试')
    },
    uploadComplete (result) {
      this.loading = false
      if (this.single) { this.fileItems = [] }
      this.fileItems.push({ url: result.url, name: result.name || this.file.name, uid: this.file.uid, status: 'done' })
      this.$emit('change', this.fileItems)
    }
  }
}
</script>
<style>
  .upload-list-inline .ant-upload-list-item {
    float: left;
    width: 200px;
    margin-right: 8px;
  }
</style>
