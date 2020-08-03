<template>
  <div>
    <a-drawer
      title="关联小区"
      :width="420"
      :mask="true"
      :visible="visible"
      :wrap-style="{height: 'calc(100% - 108px)',overflow: 'auto',paddingBottom: '108px'}"
      @close="onClose"
    >

      <!-- <a-input-search placeholder="请输入关键字" style="width: 300px" @search="onSearch" /> -->
      <a-tree
        v-model="value"
        checkable
        :expanded-keys="expandedKeys"
        :auto-expand-parent="autoExpandParent"
        :selected-keys="selectedKeys"
        :tree-data="data"
        :replace-fields="{ children:'children', title:'name', key:'id' }"
        @expand="onExpand"
      />
      <div
        :style="{
          position: 'absolute',
          bottom: 0,
          width: '100%',
          borderTop: '1px solid #e8e8e8',
          padding: '10px 16px',
          textAlign: 'right',
          left: 0,
          background: '#fff',
          borderRadius: '0 0 4px 4px',
        }"
      >
        <a-button style="marginRight: 8px" @click="onClose">
          取消
        </a-button>
        <a-button type="primary" :loading="loading" @click="handleSubmit">
          确定
        </a-button>
      </div>
    </a-drawer>
  </div>
</template>
<script>

import { getFpmsEstateCommonGroupTree } from '@/api/common'
export default {
  props: {
    staffId: {
      type: Number,
      default: null
    },
    checkedKeys: {
      type: Array,
      default: () => []
    }

  },
  data () {
    return {
      visible: false,
      labelCol: {
        xs: { span: 8 },
        sm: { span: 6 }
      },
      wrapperCol: {
        xs: { span: 8 },
        sm: { span: 10 }
      },
      listId: [],
      expandedKeys: [25, 26],
      loading: false,
      autoExpandParent: true,
      value: [],
      selectedKeys: [],
      data: []
    }
  },
  watch: {
    checkedKeys (val) {
      this.value = this.checkedKeys.map(item => {
        return `children_${item}`
      })
    }
  },
  beforeCreate () {
    this.form = this.$form.createForm(this, { name: 'validate_other' })
  },
  created () {
    this.loadDatas()
  },
  methods: {
    loadDatas () {
      getFpmsEstateCommonGroupTree().then(res => {
        this.data = this.joinData(res.data)
      })
    },
    joinData (data) {
      return data.map(item => {
        if (item.children && item.children.length) {
          item.children.map(citem => {
            citem.id = `children_${citem.id}`
            return citem
          })
        }
        return item
      })
    },
    open (id) {
      this.visible = true
    },
    onClose () {
      this.visible = false
    },
    handleSubmit () {
      const groupIds = this.value.filter(item => {
        if (typeof (item) === 'string' && ~item.indexOf('children')) {
          return true
        }
        return false
      }).map(item => {
        return ~~item.split('_')[1]
      })
      this.$emit('handleSubmit', groupIds)
    },
    onExpand (expandedKeys) {
      this.expandedKeys = expandedKeys
      this.autoExpandParent = false
    },
    onCheck (value) {
      this.value = value
    }
  }
}
</script>
