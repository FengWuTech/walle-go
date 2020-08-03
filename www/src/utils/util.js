export function timeFix () {
  const time = new Date()
  const hour = time.getHours()
  return hour < 9 ? '早上好' : hour <= 11 ? '上午好' : hour <= 13 ? '中午好' : hour < 20 ? '下午好' : '晚上好'
}

export function welcome () {
  const arr = ['休息一会儿吧', '准备吃什么呢?', '要不要打一把 DOTA', '我猜你可能累了']
  const index = Math.floor(Math.random() * arr.length)
  return arr[index]
}

/**
 * 触发 window.resize
 */
export function triggerWindowResizeEvent () {
  const event = document.createEvent('HTMLEvents')
  event.initEvent('resize', true, true)
  event.eventType = 'message'
  window.dispatchEvent(event)
}

export function handleScrollHeader (callback) {
  let timer = 0

  let beforeScrollTop = window.pageYOffset
  callback = callback || function () {}
  window.addEventListener(
    'scroll',
    event => {
      clearTimeout(timer)
      timer = setTimeout(() => {
        let direction = 'up'
        const afterScrollTop = window.pageYOffset
        const delta = afterScrollTop - beforeScrollTop
        if (delta === 0) {
          return false
        }
        direction = delta > 0 ? 'down' : 'up'
        callback(direction)
        beforeScrollTop = afterScrollTop
      }, 50)
    },
    false
  )
}

export function isIE () {
  const bw = window.navigator.userAgent
  const compare = (s) => bw.indexOf(s) >= 0
  const ie11 = (() => 'ActiveXObject' in window)()
  return compare('MSIE') || ie11
}

/**
 * Remove loading animate
 * @param id parent element id or class
 * @param timeout
 */
export function removeLoadingAnimate (id = '', timeout = 1500) {
  if (id === '') {
    return
  }
  setTimeout(() => {
    document.body.removeChild(document.getElementById(id))
  }, timeout)
}

// 获取dpi
const getDPI = () => {
  const arrDPI = []
  if (window.screen.deviceXDPI !== undefined) { // ie 9
    arrDPI[0] = window.screen.deviceXDPI
    arrDPI[1] = window.screen.deviceYDPI
  } else { // chrome firefox
    const tmpNode = document.createElement('DIV')
    tmpNode.style.cssText = 'width: 1in; height: 1in; position: absolute; left: 0px; top: 0px; z - index: 99; visibility: hidden'
    document.body.appendChild(tmpNode)

    arrDPI[0] = parseInt(tmpNode.offsetWidth)
    arrDPI[1] = parseInt(tmpNode.offsetHeight)

    tmpNode.parentNode.removeChild(tmpNode)
  }
  return arrDPI
}
// cm 转 px
export function cm2px (cm) {
  const dpi = getDPI()
  const px = parseFloat(cm) / 2.54 * dpi[0]

  return parseInt(px)
}
// px 转 cm
export function px2cm (px) {
  const dpi = getDPI()
  const cm = px * 2.54 / dpi[0]

  return parseFloat(cm.toFixed(4))
}

// 新开tab打开链接
export function openNewTab (url) {
  const a = document.createElement('a')
  a.setAttribute('href', url)
  a.setAttribute('target', '_blank')
  a.setAttribute('id', 'camnpr')

  document.body.appendChild(a)
  a.click()

  setTimeout(() => {
    a.remove()
  }, 1000)
}
