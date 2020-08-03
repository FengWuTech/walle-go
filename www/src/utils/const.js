// 此文件存放一些常量

// 收银台查询维度
const chargeDimension = [
  {
    id: 1,
    label: '房屋',
    value: 'fangwu',
    placeholder: '输入房号搜索'
  },
  {
    id: 2,
    label: '车位',
    value: 'chewei',
    placeholder: '输入车位号搜索'
  },
  {
    id: 3,
    label: '住户',
    value: 'zhuhu',
    placeholder: '输入住户姓名搜索'
  }
]

const payTypeMap = {
  '1': '房屋',
  '2': '车位',
  '3': '住户',
  '4': '户表',
  '5': '公表',
  '6': '其他'
}

const feeSearchTypes = [
  {
    label: '房屋相关',
    value: '1'
  },
  {
    label: '车位相关',
    value: '2'
  },
  {
    label: '住户相关',
    value: '3'
  },
  {
    label: '租金',
    value: '4'
  },
  {
    label: '其他',
    value: '5'
  }
]

const hasZhuHu = [
  {
    label: '有住户',
    value: '1'
  },
  {
    label: '无住户',
    value: '0'
  }
]

const feeExportTypes = [
  {
    label: '房屋',
    value: '1'
  },
  {
    label: '车位',
    value: '2'
  }
]

const chargeImportConfig = {
  // 导入未缴账单
  'nopay': { title: '导入未缴账单' },
  'payed': { title: '导入已缴账单' },
  'prepay': { title: '导入预存款' }
}

const importTypeOptions = [
  {
    label: '通过手机号',
    value: 'shouji'
  },
  {
    label: '通过房屋号',
    value: 'fangwu'
  },
  {
    label: '通过车位号',
    value: 'chewei'
  }
]

const payedSearchTypes = [
  {
    label: '房屋',
    value: '1'
  },
  {
    label: '车位',
    value: '2'
  },
  {
    label: '住户',
    value: '3'
  }
]

const feeStatusTypes = [
  {
    label: '已收',
    value: '1'
  },
  {
    label: '已退',
    value: '2'
  },
  {
    label: '退费',
    value: '3'
  }
]

const wayOptions = [
  {
    label: '预存自动抵扣',
    value: '1'
  },
  {
    label: '线上支付-微信公众号支付',
    value: '2'
  },
  {
    label: '线上支付-支付宝物业缴费',
    value: '3'
  },
  {
    label: '线上支付-支付宝生活号支付',
    value: '4'
  },
  {
    label: '线上支付-微信小程序支付',
    value: '5'
  },
  {
    label: '线下支付-现金',
    value: '6'
  },
  {
    label: '线下支付-支票',
    value: '7'
  },
  {
    label: '线下支付-银行转账',
    value: '8'
  },
  {
    label: '线下支付-pos机刷卡',
    value: '9'
  },
  {
    label: '线下支付-支付宝直接转账',
    value: '10'
  },
  {
    label: '线下支付-微信直接转账',
    value: '11'
  },
  {
    label: '小区收款码-支付宝',
    value: '12'
  },
  {
    label: '小区收款码-微信',
    value: '13'
  },
  {
    label: '小区收款码-银联支付',
    value: '14'
  },
  {
    label: '其他收费方式',
    value: '15'
  }
]

const payedStatusTypes = [
  {
    label: '全部',
    value: '0'
  },
  {
    label: '已缴',
    value: '1'
  },
  {
    label: '已退',
    value: '2'
  }
]

const sendStatusTypes = [
  {
    label: '全部',
    value: '0'
  },
  {
    label: '已发送',
    value: '1'
  },
  {
    label: '未发送',
    value: '2'
  }
]

const discountOptions = [
  {
    label: '按优惠值',
    value: '1'
  },
  {
    label: '按优惠比率',
    value: '2'
  }
]

const chargePrepayTypes = [
  {
    label: '通用预存款充值',
    value: 1
  },
  {
    label: '专用预存款充值',
    value: 2
  }
]

const prepaySearchTypes = [
  {
    label: '充值',
    value: '1'
  },
  {
    label: '抵扣',
    value: '2'
  },
  {
    label: '退预存款',
    value: '3'
  }
]

const sourceOptions = [
  {
    label: '系统',
    value: '1'
  },
  {
    label: '超级管理员',
    value: '2'
  },
  {
    label: '子管理员',
    value: '3'
  },
  {
    label: '住户',
    value: '4'
  }
]

const prepaySetTypes = [
  {
    label: '自由充值',
    value: 0,
    tips: '住户可以在微信公众号中自由选择预存款充值金额'
  },
  {
    label: '自定义套餐',
    value: 1,
    icon: 'favourable',
    tips: '预先设置预存款充值优惠套餐，住户只能选择所设置套餐金额进行充值'
  }
  // {
  //   label: '关闭充值',
  //   value: 2,
  //   tips: '住户在微信公众号无法进行预存款充值'
  // }
]

const showTypesOptions = [
  {
    label: '按收费项目显示',
    value: 0
  },
  {
    label: '按月份显示',
    value: 1
  }
]

const degreeOptions = [
  {
    label: '元',
    value: 0
  },
  {
    label: '角',
    value: 1
  },
  {
    label: '分',
    value: 2
  }
]

const znjSetTypes = [
  {
    label: '账单开始时间',
    value: 1
  },
  {
    label: '账单结束时间',
    value: 2
  }
]

const lackCycleTypes = [
  {
    label: '按周期计算',
    value: 1
  },
  {
    label: '按实际天数计算',
    value: 2
  }
]

const cycleTypes = [
  {
    label: '当期收当期',
    value: 1
  },
  {
    label: '当期收上期',
    value: 2
  },
  {
    label: '当期收下期',
    value: 3
  }
]

const cycleMonthOptions = [
  { label: '首月', value: 0 },
  { label: '尾月', value: 1 }
]

/**
 * 打印设置相关
 */
// 纸张大小
const paperSizes = [
  {
    label: 'A3',
    value: 0
  },
  {
    label: 'A4',
    value: 1
  },
  {
    label: 'A5',
    value: 2
  },
  {
    label: 'A6',
    value: 3
  }
]
// 打印布局
const printLayouts = [
  {
    label: '横向',
    value: 0
  },
  {
    label: '纵向',
    value: 1
  }
]
// 打印字体
const printFontSize = [
  {
    value: '12px'
  },
  {
    value: '13px'
  },
  {
    value: '14px'
  },
  {
    value: '15px'
  },
  {
    value: '16px'
  },
  {
    value: '17px'
  },
  {
    value: '18px'
  },
  {
    value: '19px'
  },
  {
    value: '20px'
  },
  {
    value: '21px'
  },
  {
    value: '22px'
  },
  {
    value: '23px'
  },
  {
    value: '24px'
  }
]
// 元素宽度
const printWidth = [
  {
    label: '33%',
    value: '33.3333%'
  },
  {
    label: '49%',
    value: '50%'
  },
  {
    label: '100%',
    value: '100%'
  }
]
// 元素对齐方式
const printTextAlign = [
  {
    label: '左对齐',
    value: 'left'
  },
  {
    label: '居中',
    value: 'center'
  },
  {
    label: '右对齐',
    value: 'right'
  }
]
// 固定表格元素样式
const tableItemStyles = {
  border: '1px solid #e8e8e8',
  fontWeight: '400',
  padding: '8px 0',
  textAlign: 'center'
}
// 固定表头、表尾元素样式
const normalItemStyle = { margin: '6px 0 2px 0' }
/**
 * 页眉区内容选项：
 *  label: 展示文案，在添加列的时候可以自定义，由接口返回
 *  value: 占位符，用于打印的时候替换成对应字段
 *  custom: 当选择该项的时候，可以定制化的属性
 *  styles: 该项自带的属性，不包含custom里的
 *  suitType: 在哪种打印模板类型下可添加，如：收款日期 只适用于收据模板
 *  type: 用于对应接口返回数据的字段，需要跟接口商量好
 */
const printHeadOptions = {
  title: {
    value: '标题',
    custom: ['fontSize'],
    styles: {
      textAlign: 'center',
      width: '100%'
    }
  },
  logo: {
    value: 'LOGO',
    custom: ['upload'],
    styles: {
      bottom: '0',
      height: '34px',
      left: '0',
      lineHeight: '25px',
      position: 'absolute',
      textAlign: 'right',
      width: '20%'
    }
  },
  code: {
    value: '编号',
    custom: ['fontSize'],
    styles: {
      color: 'rgb(241, 53, 42)',
      fontWeight: '600'
    }
  },
  print_date: {
    value: '打印日期',
    custom: ['label', 'fontSize', 'width', 'textAlign'],
    styles: normalItemStyle
  },
  group: {
    value: '小区',
    custom: ['label', 'fontSize', 'width', 'textAlign'],
    styles: normalItemStyle
  },
  room: {
    value: '房号',
    custom: ['label', 'fontSize', 'width', 'textAlign'],
    styles: normalItemStyle
  },
  name: {
    value: '住户姓名',
    custom: ['label', 'fontSize', 'width', 'textAlign'],
    styles: normalItemStyle
  },
  mobile: {
    value: '住户手机号',
    custom: ['label', 'fontSize', 'width', 'textAlign'],
    styles: normalItemStyle
  },
  statistics: {
    value: '合计',
    custom: ['label', 'fontSize', 'width', 'textAlign'],
    styles: normalItemStyle
  },
  rdate: {
    value: '收款日期',
    custom: ['label', 'fontSize', 'width', 'textAlign'],
    styles: normalItemStyle,
    suitType: [2]
  },
  pay_user: {
    value: '付款人',
    custom: ['label', 'fontSize', 'width', 'textAlign'],
    styles: normalItemStyle,
    suitType: [2]
  },
  recv_user: {
    value: '收款人',
    custom: ['label', 'fontSize', 'width', 'textAlign'],
    styles: normalItemStyle,
    suitType: [2]
  }
}
// 表格区内容选项
const printBodyOptions = {
  code: {
    value: '房号/车位号/住户',
    custom: ['label', 'fontSize'],
    styles: tableItemStyles
  },
  building: {
    value: '楼宇',
    custom: ['label', 'fontSize'],
    styles: tableItemStyles
  },
  unit: {
    value: '单元',
    custom: ['label', 'fontSize'],
    styles: tableItemStyles
  },
  recv_project: {
    value: '收费项目',
    custom: ['label', 'fontSize'],
    styles: tableItemStyles
  },
  recv_standard: {
    value: '收费标准',
    custom: ['label', 'fontSize'],
    styles: tableItemStyles
  },
  stime: {
    value: '开始时间',
    custom: ['label', 'fontSize'],
    styles: tableItemStyles
  },
  etime: {
    value: '结束时间',
    custom: ['label', 'fontSize'],
    styles: tableItemStyles
  },
  smeter: {
    value: '起度',
    custom: ['label', 'fontSize'],
    styles: tableItemStyles
  },
  emeter: {
    value: '止度',
    custom: ['label', 'fontSize'],
    styles: tableItemStyles
  },
  amount: {
    value: '数量',
    custom: ['label', 'fontSize'],
    styles: tableItemStyles
  },
  price: {
    value: '单价',
    custom: ['label', 'fontSize'],
    styles: tableItemStyles
  },
  total: {
    value: '金额',
    custom: ['label', 'fontSize'],
    styles: tableItemStyles
  },
  discount: {
    value: '优惠',
    custom: ['label', 'fontSize'],
    styles: tableItemStyles
  },
  latefee: {
    value: '违约金',
    custom: ['label', 'fontSize'],
    styles: tableItemStyles
  },
  should_fee: {
    value: '应收合计',
    custom: ['label', 'fontSize'],
    styles: tableItemStyles
  },
  batch: {
    value: '账单期数',
    custom: ['label', 'fontSize'],
    styles: tableItemStyles
  },
  remark: {
    value: '备注',
    custom: ['label', 'fontSize'],
    styles: tableItemStyles
  }
}
// 页脚区内容选项
const printFooterOptions = {
  print_date: {
    value: '打印日期',
    custom: ['label', 'fontSize', 'width', 'textAlign'],
    styles: normalItemStyle
  },
  group: {
    value: '小区',
    custom: ['label', 'fontSize', 'width', 'textAlign'],
    styles: normalItemStyle
  },
  room: {
    value: '房号',
    custom: ['label', 'fontSize', 'width', 'textAlign'],
    styles: normalItemStyle
  },
  name: {
    value: '住户姓名',
    custom: ['label', 'fontSize', 'width', 'textAlign'],
    styles: normalItemStyle
  },
  mobile: {
    value: '住户手机号',
    custom: ['label', 'fontSize', 'width', 'textAlign'],
    styles: normalItemStyle
  },
  statistics: {
    value: '合计',
    custom: ['label', 'fontSize', 'width', 'textAlign'],
    styles: normalItemStyle
  },
  rdate: {
    value: '收款日期',
    custom: ['label', 'fontSize', 'width', 'textAlign'],
    styles: normalItemStyle,
    suitType: [2]
  },
  pay_user: {
    value: '付款人',
    custom: ['label', 'fontSize', 'width', 'textAlign'],
    styles: normalItemStyle,
    suitType: [2]
  },
  recv_user: {
    value: '收款人',
    custom: ['label', 'fontSize', 'width', 'textAlign'],
    styles: normalItemStyle,
    suitType: [2]
  },
  recv_remark: {
    value: '收款备注',
    custom: ['label', 'fontSize', 'width', 'textAlign'],
    styles: normalItemStyle,
    suitType: [2]
  },
  recv_method: {
    value: '收款方式',
    custom: ['label', 'fontSize', 'width', 'textAlign'],
    styles: normalItemStyle,
    suitType: [2]
  },
  remark: {
    value: '说明',
    custom: ['label', 'fontSize', 'width', 'textAlign'],
    styles: normalItemStyle
  }
}

// 工单状态
const orderStatusTypes = [
  {
    label: '待受理',
    value: 0
  },
  {
    label: '待回复',
    value: 1
  },
  {
    label: '处理中',
    value: 2
  },
  {
    label: '已回复',
    value: 3
  },
  {
    label: '已完结',
    value: 4
  }
]

// 首页快速开始连接
const startLinks = [
  { title: '商户列表', name: 'PayMerchantList' },
  { title: '应该管理', name: 'WechatAppList' },
  { title: '系统通知', name: 'NoticeList' },
  { title: '添加系统通知', name: 'NoticeCreate' },
  { title: '工单列表', name: 'FeedbackManager' },
  { title: '公司列表', name: 'Company' },
  { title: '添加公司', name: 'companyCreate' },
  { title: '问题列表', name: 'CallbackList' },
  { title: '添加回访问题', name: 'CallbackCreate' },
  { title: '大区管理', name: 'basisBigarea' },
  { title: '添加大区', name: 'basisBigareaCreate' },
  { title: '费项管理', name: 'basisCharge' },
  { title: '产品管理', name: 'basisProduct' },
  { title: '添加产品', name: 'basisProductCreate' }
]

export {
  chargeDimension,
  feeSearchTypes,
  hasZhuHu,
  feeExportTypes,
  chargeImportConfig,
  importTypeOptions,
  payedSearchTypes,
  feeStatusTypes,
  wayOptions,
  payedStatusTypes,
  sendStatusTypes,
  discountOptions,
  chargePrepayTypes,
  prepaySearchTypes,
  sourceOptions,
  prepaySetTypes,
  showTypesOptions,
  degreeOptions,
  znjSetTypes,
  lackCycleTypes,
  cycleTypes,
  paperSizes,
  printLayouts,
  printHeadOptions,
  printFontSize,
  printWidth,
  printTextAlign,
  printBodyOptions,
  printFooterOptions,
  startLinks,
  orderStatusTypes,
  cycleMonthOptions,
  payTypeMap
}
