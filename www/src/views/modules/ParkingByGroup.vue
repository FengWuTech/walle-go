<template>
  <div>
    <a-form-item label="选择组团">
      <FwCascader :maxLevel="2" @change="getParkingList"></FwCascader>
    </a-form-item>
    <a-form-item label="选择车位" required>
      <a-select :mode="mode" :options="bindingOptions" v-model="myValue" :placeholder="placeholder" @change="onValueChange"></a-select>
    </a-form-item>
  </div>
</template>

<script>
import { doGetParkingListByGroupId } from '../../utils'

export default {
  name: 'ParkingByGroup',
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
      default: '请选择车位'
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
     * 当组团被选中的时候，车位列表需要有所变化
     */
    getParkingList (ids) {
      const id = ids[ids.length - 1]
      doGetParkingListByGroupId(id).then(options => {
        this.bindingOptions = options
      })
    },

    onValueChange (value) {
      this.$emit('change', value)
    }
  }
}
</script>
