<template>
  <div>
    <a-form-item label="选择组团">
      <FwCascader :maxLevel="2" @change="onChange"></FwCascader>
    </a-form-item>
    <a-form-item label="选择楼宇" required>
      <a-select :mode="mode" :options="bindingOptions" v-model="myValue" :placeholder="placeholder" @change="onValueChange"></a-select>
    </a-form-item>
  </div>
</template>

<script>
import { doGetEstateList } from '@/utils'

export default {
  name: 'BuildingByGroup',
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
      default: '请选择楼宇'
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
    onChange (ids) {
      const id = ids[ids.length - 1]
      this.getEstateList(id)
    },

    /**
     * 当绑定组团被选中的时候，楼宇列表需要有所变化
     */
    getEstateList (id) {
      doGetEstateList(id, 2).then(res => {
        if (res.code === 200) {
          const data = res.data
          if (!data || !data.length) {
            return
          }
          const newOptions = data.map(item => {
            return {
              label: item.name,
              value: item.id
            }
          })
          console.log(newOptions)
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
