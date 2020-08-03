<template>
  <div>
    <a-form-item label="选择单元">
      <FwCascader :maxLevel="4" @change="getUserListByUnitId"></FwCascader>
    </a-form-item>
    <a-form-item label="选择住户" required>
      <a-select :mode="mode" :options="bindingOptions" v-model="myValue" :placeholder="placeholder" @change="onValueChange"></a-select>
    </a-form-item>
  </div>
</template>

<script>
import { getUserList } from '@/api/base/user'

export default {
  name: 'UserByUnit',
  props: {
    inputValue: {
      type: [String, Number, Array],
      default: undefined
    },
    inputOptions: {
      type: Array,
      default () {
        return []
      }
    },
    placeholder: {
      type: String,
      default: '请选择住户'
    },
    mode: {
      type: String,
      default: 'default'
    }
  },
  data () {
    return {
      bindForm: this.$form.createForm(this),
      myValue: undefined,
      bindingOptions: []
    }
  },
  watch: {
    inputValue () {
      this.myValue = this.inputValue
    },
    inputOptions () {
      this.bindingOptions = this.inputOptions
    }
  },
  created () {
    this.myValue = this.inputValue
    if (this.myValue && this.myValue !== undefined) {
      this.$emit('change', this.myValue)
    }
    this.bindingOptions = this.inputOptions
  },
  methods: {
    /**
     * 当单元被选中的时候，住户列表需要有所变化
     */
    getUserListByUnitId (ids) {
      const id = ids[ids.length - 1]
      getUserList({ unit_id: id, page: 1, page_size: 1000 }).then(res => {
        if (res.code === 200) {
          const data = res.data.list
          if (!data || !data.length) {
            return
          }
          const newOptions = data.map(item => {
            return {
              label: item.name,
              value: item.id
            }
          })
          this.bindingOptions = newOptions
        } else {
          this.$message.error(res.msg)
        }
      })
    },

    onValueChange (value) {
      console.log(value)
      this.$emit('change', value)
    }
  }
}
</script>
