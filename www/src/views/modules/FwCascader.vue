<template>
  <a-cascader
    :disabled="disabled"
    :options="options"
    :fieldNames="fieldNames"
    :loadData="loadData"
    :changeOnSelect="changeOnSelect"
    v-model="myValue"
    :placeholder="newPlaceholder"
    @change="onChange"
  />
</template>

<script>
import { getEstateList } from '@/api/common'
const levelMap = {
  1: '项目',
  2: '组团',
  3: '楼宇',
  4: '单元',
  5: '房屋'
}

export default {
  name: 'FwCascader',
  props: {
    // 1项目、2组团、3楼宇、4单元、5房屋
    maxLevel: {
      type: Number,
      default: 2
    },
    inputValue: {
      type: Array,
      default () {
        return []
      }
    },
    disabled: {
      type: Boolean,
      default: false
    },
    fieldNames: {
      type: Object,
      default () {
        return {
          label: 'label',
          value: 'value',
          children: 'children'
        }
      }
    },
    placeholder: {
      type: String,
      default: ''
    },
    changeOnSelect: {
      type: Boolean,
      default: false
    }
  },
  data () {
    return {
      myValue: [],
      options: []
    }
  },
  watch: {
    inputValue () {
      this.myValue = this.inputValue
    }
  },
  computed: {
    newPlaceholder () {
      if (this.placeholder) {
        return this.placeholder
      } else {
        return '请选择所属' + levelMap[this.maxLevel]
      }
    }
  },
  created () {
    this.myValue = this.inputValue
    this.loadData()
  },
  methods: {
    onChange () {
      this.$emit('change', this.myValue)
    },

    /**
     * 获取 0项目、1组团、2楼宇、3单元、4房间列表的接口
     */
    doGetEstateList (pid, pidLevel) {
      return new Promise((resolve, reject) => {
        const data = {
          pid: pid,
          pid_level: pidLevel
        }
        getEstateList(data).then(res => {
          if (res.code === 200) {
            resolve(res)
          } else {
            reject(res)
          }
        }).catch(err => {
          reject(err)
        })
      })
    },

    /**
     * 加载数据,这里有两种情况需要注意，一种是初始化的时候获取第一级的数据，这时是在构造函数中触发的，selectedOptions为undefined
     * 这时候直接把结果赋值给options即可，其他情况都是当用户点击级联节点的时候触发的，一定存在selectedOptions字段
     * @param selectedOptions
     * @returns {Promise<void>}
     */
    async loadData (selectedOptions) {
      let [pid, level] = [0, 0]
      let targetOption = null
      if (selectedOptions) {
        targetOption = selectedOptions[selectedOptions.length - 1]
        pid = targetOption.value
        level = targetOption.level
        targetOption.loading = true
      }
      const res = await this.doGetEstateList(pid, level)
      const data = res.data
      let tempData = null
      if (data && data.length > 0) {
        tempData = data.map(item => {
          return {
            label: item.name,
            value: item.id,
            level: level + 1,
            isLeaf: this.maxLevel === level + 1
          }
        })
      }

      if (!targetOption) {
        this.options = tempData
      } else {
        targetOption.loading = false
        if (data && data.length > 0) {
          targetOption.children = tempData
        } else {
          targetOption.children = [{
            value: targetOption + Math.random(),
            disabled: true,
            label: '暂无数据',
            level: level + 1,
            isLeaf: true
          }]
        }
        this.options = [...this.options]
      }
    }
  }
}
</script>
