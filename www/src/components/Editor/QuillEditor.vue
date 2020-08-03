<template>
  <div>
    <quill-editor
      v-model="content"
      ref="myQuillEditor"
      class="editor"
      :options="editorOption"
      @blur="onEditorBlur($event)"
      @focus="onEditorFocus($event)"
      @change="onEditorChange($event)">
    </quill-editor>
    <a-upload
      v-show="false"
      name="file"
      ref="upload"
      :disabled="disabled"
      accept=".jpeg"
      :before-upload="beforeUpload"
      :remove="removeFile"
      :file-list="fileItems"
      list-type="picture"
      class="upload-list-inline"
    >
      <a-button :disabled="disabled" ref="imgInput" icon="upload" :loading="loading">
        上传图片
      </a-button>
    </a-upload>
  </div>
</template>
<script>
// require styles 引入样式
import 'quill/dist/quill.core.css'
import 'quill/dist/quill.snow.css'
import 'quill/dist/quill.bubble.css'
import Quill from 'quill'
import { quillEditor } from 'vue-quill-editor'

import upload from '@/views/modules/upload'
import { getUploadToken } from '@/api/common'
import * as qiniu from 'qiniu-js'

// 工具栏配置
const toolbarOptions = [
  ['bold', 'italic', 'underline', 'strike'], // 加粗 斜体 下划线 删除线 -----['bold', 'italic', 'underline', 'strike']
  ['blockquote', 'code-block'], // 引用  代码块-----['blockquote', 'code-block']
  [{ header: 1 }, { header: 2 }], // 1、2 级标题-----[{ header: 1 }, { header: 2 }]
  [{ list: 'ordered' }, { list: 'bullet' }], // 有序、无序列表-----[{ list: 'ordered' }, { list: 'bullet' }]
  // [{ script: 'sub' }, { script: 'super' }], // 上标/下标-----[{ script: 'sub' }, { script: 'super' }]
  // [{ indent: '-1' }, { indent: '+1' }], // 缩进-----[{ indent: '-1' }, { indent: '+1' }]
  // [{ 'direction': 'rtl' }], // 文本方向-----[{'direction': 'rtl'}]
  [{ size: ['small', false, 'large', 'huge'] }], // 字体大小-----[{ size: ['small', false, 'large', 'huge'] }]
  [{ header: [1, 2, 3, 4, 5, 6, false] }], // 标题-----[{ header: [1, 2, 3, 4, 5, 6, false] }]
  [{ color: [] }, { background: [] }], // 字体颜色、字体背景颜色-----[{ color: [] }, { background: [] }]
  [{ font: [] }], // 字体种类-----[{ font: [] }]
  [{ align: [] }], // 对齐方式-----[{ align: [] }]
  // ['clean'], // 清除文本格式-----['clean']
  ['image'] // 链接、图片、视频-----['link', 'image', 'video']
]
export default {
  components: { quillEditor, upload },
  props: {
    initValue: {
      type: String,
      default: ''
    }
  },
  data () {
    return {
      content: '',
      disabled: false,
      file: null,
      fileItems: [],
      loading: false,

      editorOption: {
        placeholder: '请输入文本...',
        theme: 'snow',
        modules: {
          toolbar: {
            container: toolbarOptions
          }
        }
      }
    }
  },
  watch: {
    initValue () {
      this.content = this.initValue
    }
  },
  created () {
    this.content = this.initValue
  },
  mounted () {
    // 为图片ICON绑定事件  getModule 为编辑器的内部属性
    this.$refs.myQuillEditor.quill.getModule('toolbar').addHandler('image', this.imgHandler)
    this.$refs.myQuillEditor.quill.getModule('toolbar').addHandler('video', this.videoHandler) // 为视频ICON绑定事件
  },
  methods: {
    onEditorBlur () {
      // 失去焦点事件
    },
    onEditorFocus () {
      // 获得焦点事件
    },
    onEditorChange () {
      this.$emit('onEditorChange', this.content)
      this.$emit('change', this.content)
      // 内容改变事件
    },
    changeContent (val) {
      this.content = val
    },
    removeFile () {
      this.file = null
      this.fileItems = this.fileItems.length && this.fileItems.filter(item => item.status !== 'removed')
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
      const suffix = 'jpeg'
      getUploadToken({ suffix }).then(res => {
        if (res.data) {
          const uploadToken = res.data.token || ''
          const uploadKey = res.data.filekey || file.name
          const observable = qiniu.upload(file, uploadKey, uploadToken)

          const subObj = {
            next: () => {},
            error: this.uploadError,
            complete: this.uploadComplete
          }
          observable.subscribe(subObj)
        } else {
          this.loading = false
        }
      }).catch(_ => {
        this.loading = false
      })
    },

    uploadError (err) {
      this.loading = false
      console.log('err', err)
      this.$message.error('文件上传失败，请重试')
    },
    uploadComplete (result) {
      console.log('res', result)
      this.loading = false
      const url = `http://pms-static.gmtech.top/${result.key}`
      this.addRange = this.$refs.myQuillEditor.quill.getSelection()
      this.$refs.myQuillEditor.quill.insertEmbed(this.addRange !== null ? this.addRange.index : 0, this.uploadType, url, Quill.sources.USER) // 调用编辑器的 insertEmbed 方法，插入URL
    },

    // 点击图片ICON触发事件
    imgHandler (state) {
      this.addRange = this.$refs.myQuillEditor.quill.getSelection()
      console.log(this.$refs.imgInput)
      if (state) {
        this.$refs.imgInput.$el.click()
      }
      this.uploadType = 'image'
    },
    // 点击视频ICON触发事件
    videoHandler (state) {
      this.addRange = this.$refs.myQuillEditor.quill.getSelection()
      if (state) {
        this.$refs.imgInput.$el.click()
      }
      this.uploadType = 'video'
    }
  }
}
</script>
<style lang="less">
@import url('../index.less');

/* 覆盖 quill 默认边框圆角为 ant 默认圆角，用于统一 ant 组件风格 */
.editor {
  /deep/ .ql-toolbar.ql-snow {
    border-radius: @border-radius-base @border-radius-base 0 0;
  }
  /deep/ .ql-container.ql-snow {
    border-radius: 0 0 @border-radius-base @border-radius-base;
  }
}

.editor {
  line-height: normal !important;
  height: 360px;
}
.ql-snow .ql-tooltip[data-mode=link]::before {
  content: "请输入链接地址:";
}
.ql-snow .ql-tooltip.ql-editing a.ql-action::after {
    border-right: 0px;
    content: '保存';
    padding-right: 0px;
}

.ql-snow .ql-tooltip[data-mode=video]::before {
    content: "请输入视频地址:";
}

.ql-snow .ql-picker.ql-size .ql-picker-label::before,
.ql-snow .ql-picker.ql-size .ql-picker-item::before {
  content: '14px';
}
.ql-snow .ql-picker.ql-size .ql-picker-label[data-value=small]::before,
.ql-snow .ql-picker.ql-size .ql-picker-item[data-value=small]::before {
  content: '10px';
}
.ql-snow .ql-picker.ql-size .ql-picker-label[data-value=large]::before,
.ql-snow .ql-picker.ql-size .ql-picker-item[data-value=large]::before {
  content: '18px';
}
.ql-snow .ql-picker.ql-size .ql-picker-label[data-value=huge]::before,
.ql-snow .ql-picker.ql-size .ql-picker-item[data-value=huge]::before {
  content: '32px';
}

.ql-snow .ql-picker.ql-header .ql-picker-label::before,
.ql-snow .ql-picker.ql-header .ql-picker-item::before {
  content: '文本';
}
.ql-snow .ql-picker.ql-header .ql-picker-label[data-value="1"]::before,
.ql-snow .ql-picker.ql-header .ql-picker-item[data-value="1"]::before {
  content: '标题1';
}
.ql-snow .ql-picker.ql-header .ql-picker-label[data-value="2"]::before,
.ql-snow .ql-picker.ql-header .ql-picker-item[data-value="2"]::before {
  content: '标题2';
}
.ql-snow .ql-picker.ql-header .ql-picker-label[data-value="3"]::before,
.ql-snow .ql-picker.ql-header .ql-picker-item[data-value="3"]::before {
  content: '标题3';
}
.ql-snow .ql-picker.ql-header .ql-picker-label[data-value="4"]::before,
.ql-snow .ql-picker.ql-header .ql-picker-item[data-value="4"]::before {
  content: '标题4';
}
.ql-snow .ql-picker.ql-header .ql-picker-label[data-value="5"]::before,
.ql-snow .ql-picker.ql-header .ql-picker-item[data-value="5"]::before {
  content: '标题5';
}
.ql-snow .ql-picker.ql-header .ql-picker-label[data-value="6"]::before,
.ql-snow .ql-picker.ql-header .ql-picker-item[data-value="6"]::before {
  content: '标题6';
}

.ql-snow .ql-picker.ql-font .ql-picker-label::before,
.ql-snow .ql-picker.ql-font .ql-picker-item::before {
  content: '标准字体';
}
.ql-snow .ql-picker.ql-font .ql-picker-label[data-value=serif]::before,
.ql-snow .ql-picker.ql-font .ql-picker-item[data-value=serif]::before {
  content: '衬线字体';
}
.ql-snow .ql-picker.ql-font .ql-picker-label[data-value=monospace]::before,
.ql-snow .ql-picker.ql-font .ql-picker-item[data-value=monospace]::before {
  content: '等宽字体';
}
</style>
