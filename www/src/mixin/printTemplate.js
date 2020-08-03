// 打印模板预览和打印页面的公共方法
export const print = {
  props: {
    printInfo: {
      type: Object,
      default: () => {}
    }
  },
  data () {
    return {
      rowStyle: {
        width: '100%',
        display: 'flex'
      },
      tdStyle: {
        height: '30px',
        textAlign: 'center',
        border: '1px solid rgb(232, 232, 232)'
      },
      numStyle: {
        width: '20%',
        textAlign: 'right',
        lineHeight: '25px',
        position: 'absolute',
        bottom: 0,
        right: 0
      }
    }
  },
  computed: {
    headTitle () {
      return this.getSpecialField('header', 'title')
    },
    headLogo () {
      return this.getSpecialField('header', 'logo')
    },
    headNum () {
      return this.getSpecialField('header', 'code')
    }
  },
  methods: {
    // 获取某些特定的字段 - 标题、LOGO、编号、小区收款码
    getSpecialField (area, name) {
      if (!this['printInfo'][area]) { return null }
      const res = this['printInfo'][area].filter(item => item.type === name)
      return res && res[0] || null
    },
    // 判断是否是特定行
    isSpecialField (type) {
      return ['title', 'logo', 'code'].indexOf(type) > -1
    },
    // string类型的style转为对象
    parseStyle (item) {
      return typeof item.style === 'string' ? JSON.parse(item.style) : item.style
    }
  }
}
