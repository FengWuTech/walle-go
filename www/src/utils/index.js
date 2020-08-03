// import { getEstateList } from '@/api/common'
// import { getAllTagList } from '@/api/base/tag'
// import { getUserList } from '@/api/base/user'
// import { getParkingListByGroupId } from '@/api/base/parking'

/**
 * 通用的分页配置信息
 * @type {{current: number, total: number, pageSizeOptions: [string, string, string, string], showTotal(*, *): string, pageSize: number, showSizeChanger: boolean}}
 */
export const pagination = {
  current: 1,
  pageSize: 20,
  total: 0,
  showSizeChanger: true,
  pageSizeOptions: ['10', '20', '50', '100'],
  showTotal (total, range) {
    return `共${total}条记录，当前第${range[0]}条~第${range[1]}条。`
  }
}

/**
 * 通用表单label布局配置
 * @type {{xs: {span: number}, sm: {span: number}}}
 */
export const labelCol = {
  xs: { span: 24 },
  sm: { span: 6 }
}

/**
 * 通用表单value布局配置
 * @type {{xs: {span: number}, sm: {span: number}}}
 */
export const wrapperCol = {
  xs: { span: 24 },
  sm: { span: 14 }
}

/**
 * 必填文本域配置，统一放在一个地方容易管理
 * @param text
 * @returns {{message: string, required: boolean}}
 */
export function getRequiredRules (text) {
  return { required: true, message: `请输入${text}` }
}

/**
 * 返回数字类型的文本，往接口传的时候需要转一下
 * @returns {{validator(*, *=, *): void, type: string}}
 */
export function getNumicRules () {
  return {
    type: 'string',
    validator (rule, value, callback) {
      const message = '请填写数字类型值'
      const isNumeric = function (n) {
        return !isNaN(parseFloat(n)) && isFinite(n)
      }
      if (isNumeric(value) || value === '' || value === null) {
        callback()
      } else {
        callback(message)
      }
    }
  }
}

/**
 * 获取 pidLevel 0项目、1组团、2楼栋、3单元、4房间列表的接口
 */
// export function doGetEstateList (pid, pidLevel) {
//   return new Promise((resolve, reject) => {
//     const data = {
//       pid: pid,
//       pid_level: pidLevel
//     }
//     getEstateList(data).then(res => {
//       if (res.code === 200) {
//         resolve(res)
//       } else {
//         reject(res)
//       }
//     }).catch(err => {
//       reject(err)
//     })
//   })
// }

// 数组转对象
export function arr2obj (arr = [], key = 'value', value = 'label') {
  const obj = {}
  arr.forEach(item => {
    obj[item[key]] = item[value]
  })

  return obj
}

/**
 * 查询所有标签，用在标签下拉列表的下拉框里边
 * @param params
 * category	int	Y	枚举， 1房屋、2车位、3住户、4楼宇. （如入参不是上述值，则报参数错误）
 * name	string	Y	tag名字，支持模糊搜索
 * @param type
 */
// export function getAllTagListByCategory (category) {
//   return new Promise((resolve, reject) => {
//     getAllTagList({ category }).then(res => {
//       if (res.code === 200) {
//         if (!res.data.list || res.data.list.length === 0) {
//           resolve([])
//         }
//         const tagList = res.data.list.map(item => {
//           return {
//             label: item.name,
//             value: item.id
//           }
//         })
//         resolve(tagList)
//       } else {
//         reject(res)
//       }
//     })
//   })
// }

/**
 * 通过姓名过滤用户列表
 * @param name 用户名
 */
// export function filterUserByName (name) {
//   return new Promise((resolve, reject) => {
//     let searchResults = []
//     if (name.length < 2) {
//       resolve(searchResults)
//     }
//     getUserList({ name }).then(res => {
//       if (res.code === 200) {
//         if (res.data.list && res.data.list.length > 0) {
//           searchResults = res.data.list.map(item => {
//             return {
//               label: `${item.name}(${item.basic_mobile})`,
//               value: item.id
//             }
//           })
//         }
//         resolve(searchResults)
//       } else {
//         reject(res)
//       }
//     }).catch(res => {
//       reject(res)
//     })
//   })
// }

/**
 * 将options1和options2中value不同的项过滤出来，也就是合并两个options并保证value唯一
 * @param options1
 * @param options2
 * @returns {*[]}
 */
export function getUniqOptions (options1, options2) {
  const addonKeys = options1.map(item => item.value)
  const filterResOptions = options2.filter(item => addonKeys.indexOf(item.value) === -1)
  return [...options1, ...filterResOptions]
}

/**
 * 通过组团ID来获取车位列表
 * @param groupId
 * @returns {Promise<unknown>}
 */
// export function doGetParkingListByGroupId (groupId) {
//   return new Promise((resolve, reject) => {
//     getParkingListByGroupId({ group_id: groupId }).then(res => {
//       if (res.code === 200) {
//         let options = []
//         if (res.data && res.data.length > 0) {
//           options = res.data.map(item => {
//             return {
//               label: item.name,
//               value: item.id
//             }
//           })
//         }
//         resolve(options)
//       } else {
//         reject(res)
//       }
//     }).catch(err => {
//       reject(err)
//     })
//   })
// }

/**
 * 车位状态对应的badge status
 * @type {{'1': string, '2': string, '3': string, '4': string, '5': string, '6': string, '7': string}}
 */
export const parkingStatusMap = {
  '1': 'success',
  '2': 'success',
  '3': 'error',
  '4': 'default',
  '5': 'processing',
  '6': 'processing',
  '7': 'warning'
}

/**
 * 这里的格式化有点特点，将children的第一项赋值给父节点，然后删除第一项子节点，这样数据才能展示正常
 * @param zhList
 * @param childrenKey
 * @returns {Array|*}
 */
export function formatZhData (zhList, childrenKey) {
  if (!zhList) {
    return []
  }
  for (var i = 0; i < zhList.length; i++) {
    const item = zhList[i]
    if (item[childrenKey]) {
      Object.assign(item, item[childrenKey][0])
      item[childrenKey].shift()
      if (item[childrenKey].length === 0) {
        delete item[childrenKey]
      }
    }
  }
  return zhList
}

export function launchIntoFullscreen (element) {
  if (element.requestFullscreen) {
    element.requestFullscreen()
  } else if (element.mozRequestFullScreen) {
    element.mozRequestFullScreen()
  } else if (element.webkitRequestFullscreen) {
    element.webkitRequestFullscreen()
  } else if (element.msRequestFullscreen) {
    element.msRequestFullscreen()
  }
}

export function exitFullscreen () {
  if (document.exitFullscreen) {
    document.exitFullscreen()
  } else if (document.mozCancelFullScreen) {
    document.mozCancelFullScreen()
  } else if (document.webkitExitFullscreen) {
    document.webkitExitFullscreen()
  }
}

export function getFullscreenElement () {
  var fullscreenEle = document.fullscreenElement || document.mozFullScreenElement || document.webkitFullscreenElement
  return fullscreenEle
}

/**
 * 导出excel等操作会用到
 * @param blob
 * @param name
 */
export function blobToFile (blob, name) {
  const url = window.URL.createObjectURL(blob)
  var downloadElement = document.createElement('a')
  downloadElement.href = url
  downloadElement.download = name // 下载后文件名
  document.body.appendChild(downloadElement)
  downloadElement.click() // 点击下载
  document.body.removeChild(downloadElement) // 下载完成移除元素
  window.URL.revokeObjectURL(url) // 释放掉blob对象
}
