<template>
  <a-select
    showSearch
    :allowClear="allowClear"
    :size="size"
    :placeholder="placeholder"
    :showArrow="false"
    :filterOption="false"
    @search="handleSearch"
    @change="onChange"
    notFoundContent="暂无相关数据">
    <a-select-option v-for="d in dataSource" :key="d.value">{{ d.label }}</a-select-option>
  </a-select>
</template>

<script>
import { getRoomList } from '@/api/base/room'

export default {
  name: 'UserSelect',
  props: {
    allowClear: {
      type: Boolean,
      default: true
    },
    size: {
      type: String,
      default: ''
    },
    placeholder: {
      type: String,
      default: '搜索说明：输入房间号即可快速搜索'
    }
  },
  data () {
    return {
      dataSource: []
    }
  },
  methods: {
    // 输入补全
    handleSearch (name) {
      this.dataSource = []
      if (!name) { return }
      const params = { name, page: 1, page_size: 1000 }
      getRoomList(params).then(res => {
        if (res.code === 200) {
          const data = res.data.list || []
          if (!data || !data.length) { return }
          this.dataSource = data.map(item => {
            return {
              label: `${item.name}-${item.belong_group_name}-${item.belong_building_name}-${item.belong_unit_name}`,
              value: item.id
            }
          })
        } else {
          this.$message.error(res.msg)
        }
      })
    },

    onChange (value) {
      this.$emit('change', value)
    }
  }
}
</script>
