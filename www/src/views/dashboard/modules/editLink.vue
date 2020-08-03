<template>
  <a-modal
    title="快速开始"
    :width="520"
    :visible="visible"
    @ok="handleSubmit"
    cancelText="关闭"
    @cancel="handleCancel"
  >
    <a-checkbox-group v-model="selected">
      <a-row>
        <a-col :span="8" v-for="(item, index) in startLinks" :key="index" style="margin-bottom: 8px;">
          <a-checkbox :value="item.title">{{ item.title }}</a-checkbox>
        </a-col>
      </a-row>
    </a-checkbox-group>

    <template slot="footer">
      <a-button key="back" @click="handleCancel">取消</a-button>
      <a-button key="submit" type="primary" @click="handleSubmit">确定</a-button>
    </template>
  </a-modal>
</template>
<script>
import { labelCol, wrapperCol } from '@/utils'
import { startLinks } from '@/utils/const'
import Vue from 'vue'
export default {
  data () {
    return {
      labelCol,
      wrapperCol,
      visible: false,
      startLinks,
      selected: []
    }
  },
  methods: {
    add () {
      this.getSelectedLinks()
      this.visible = true
    },
    handleCancel (e) {
      this.visible = false
    },
    getSelectedLinks () {
      const data = Vue.ls.get('startFccsLink')
      this.selected = (data || []).map(item => item.title)
    },
    handleSubmit () {
      const data = this.startLinks.filter(item => this.selected.indexOf(item.title) > -1)
      Vue.ls.set('startFccsLink', data)
      this.$emit('ok', data)
      this.handleCancel()
    }
  }
}
</script>
