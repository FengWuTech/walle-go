<template>
  <a-input-group class="ant-input-group-wrapper">
    <span class="ant-input-wrapper ant-input-group">
      <span v-if="!!beforeText" class="ant-input-group-addon">{{ beforeText }}</span>
      <a-input-number
        v-model="myInputNumber"
        :style="!!beforeText?innerStyle:innerRightStyle"
        :placeholder="placeholder"
        :min="min"
        :max="max"
        :disabled="disabled"
        :precision="precision"
        @change="inputChange" />
      <span class="ant-input-group-addon">{{ afterText }}</span>
    </span>
  </a-input-group>
</template>

<script>
export default {
  name: 'InputNumberWithAddon',
  props: {
    inputValue: {
      type: [Number, String],
      default: 0
    },
    placeholder: {
      type: String,
      default: ''
    },
    min: {
      type: Number,
      default: 0
    },
    max: {
      type: Number,
      default: Infinity
    },
    precision: {
      type: Number,
      default: 2
    },
    beforeText: {
      type: String,
      default: ''
    },
    afterText: {
      type: String,
      default: '元'
    },
    disabled: {
      type: Boolean,
      default: false
    }
  },
  data () {
    return {
      innerStyle: {
        width: '100%',
        borderRadius: 0
      },
      innerRightStyle: {
        width: '100%',
        borderTopRightRadius: 0,
        borderBottomRightRadius: 0
      },
      myInputNumber: undefined
    }
  },
  watch: {
    inputValue () {
      this.myInputNumber = this.inputValue
    }
  },
  created () {
    this.myInputNumber = this.inputValue
    if (this.myInputNumber && this.myInputNumber !== undefined) {
      this.$emit('change', this.myInputNumber)
    }
  },
  methods: {
    inputChange () {
      this.$emit('change', this.myInputNumber)
    }
  }
}
</script>
