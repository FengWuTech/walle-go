/**
 * 本JS是加载Lodop插件及CLodop服务
 */

let CLodopIsLocal
let CLodopJsState
let CreatedOKLodopObject

// 判断是否需要CLodop(那些不支持插件的浏览器)
export function needCLodop () {
  try {
    const ua = navigator.userAgent
    if (ua.match(/Windows\sPhone/i)) {
      return true
    }
    if (ua.match(/iPhone|iPod|iPad/i)) {
      return true
    }
    if (ua.match(/Android/i)) {
      return true
    }
    if (ua.match(/Edge\D?\d+/i)) {
      return true
    }

    const verTrident = ua.match(/Trident\D?\d+/i)
    const verIE = ua.match(/MSIE\D?\d+/i)
    let verOPR = ua.match(/OPR\D?\d+/i)
    let verFF = ua.match(/Firefox\D?\d+/i)
    const x64 = ua.match(/x64/i)
    if ((!verTrident) && (!verIE) && (x64)) {
      return true
    } else if (verFF) {
      verFF = verFF[0].match(/\d+/)
      if ((verFF[0] >= 41) || (x64)) {
        return true
      }
    } else if (verOPR) {
      verOPR = verOPR[0].match(/\d+/)
      if (verOPR[0] >= 32) {
        return true
      }
    } else if ((!verTrident) && (!verIE)) {
      let verChrome = ua.match(/Chrome\D?\d+/i)
      if (verChrome) {
        verChrome = verChrome[0].match(/\d+/)
        if (verChrome[0] >= 41) {
          return true
        }
      }
    }
    return false
  } catch (err) {
    return true
  }
}

// 获取LODOP对象主过程,判断是否安装、需否升级
export function getLodop (oOBJECT, oEMBED) {
  const strHtmInstall = '<br><font color="#FF00FF">打印控件未安装!点击这里<a href="http://www.lodop.net/demolist/install_lodop32.exe" target="_self">执行安装</a>,安装后请刷新页面或重新进入。</font>'
  const strHtmUpdate = '<br><font color="#FF00FF">打印控件需要升级!点击这里<a href="http://www.lodop.net/demolist/install_lodop32.exe" target="_self">执行升级</a>,升级后请重新进入。</font>'
  const strHtm64Install = '<br><font color="#FF00FF">打印控件未安装!点击这里<a href="http://www.lodop.net/demolist/install_lodop64.exe" target="_self">执行安装</a>,安装后请刷新页面或重新进入。</font>'
  const strHtm64Update = '<br><font color="#FF00FF">打印控件需要升级!点击这里<a href="http://www.lodop.net/demolist/install_lodop64.exe" target="_self">执行升级</a>,升级后请重新进入。</font>'
  const strHtmFireFox = '<br><br><font color="#FF00FF">（注意：如曾安装过Lodop旧版附件npActiveXPLugin,请在【工具】->【附加组件】->【扩展】中先卸它）</font>'
  const strHtmChrome = '<br><br><font color="#FF00FF">(如果此前正常，仅因浏览器升级或重安装而出问题，需重新执行以上安装）</font>'
  const strCLodopInstall1 = '<br><font color="#FF00FF">Web打印服务CLodop未安装启动，点击这里<a href="http://www.lodop.net/demolist/CLodop_Setup_for_Win32NT.exe" target="_self">下载执行安装</a>'
  const strCLodopInstall2 = '<br>（若此前已安装过，可<a href="CLodop.protocol:setup" target="_self">点这里直接再次启动</a>）'
  const strCLodopInstall3 = '，成功后请刷新本页面。</font>'
  const strCLodopUpdate = '<br><font color="#FF00FF">Web打印服务CLodop需升级!点击这里<a href="http://www.lodop.net/demolist/CLodop_Setup_for_Win32NT.exe" target="_self">执行升级</a>,升级后请刷新页面。</font>'
  let LODOP, is64IE

  try {
    const ua = navigator.userAgent
    const isIE = !!(ua.match(/MSIE/i)) || !!(ua.match(/Trident/i))
    if (needCLodop()) {
      try {
        LODOP = window.getCLodop()
      } catch (err) {
      }
      if (!LODOP && (document.readyState !== 'complete' || (isIE && CLodopJsState === 'loading'))) {
        alert('网页还没下载完毕，请稍等一下再操作.')
      }
      if (!LODOP) {
        document.body.insertAdjacentHTML('afterBegin', strCLodopInstall1 + (CLodopIsLocal ? strCLodopInstall2 : '') + strCLodopInstall3)
      } else {
        if (window.CLODOP.CVERSION < '3.0.9.3') {
          document.body.insertAdjacentHTML('afterBegin', strCLodopUpdate)
        }
        if (oEMBED && oEMBED.parentNode) {
          oEMBED.parentNode.removeChild(oEMBED)
        }
        if (oOBJECT && oOBJECT.parentNode) {
          oOBJECT.parentNode.removeChild(oOBJECT)
        }
      }
    } else {
      is64IE = isIE && !!(ua.match(/x64/i))
      // 如果页面有Lodop就直接使用,否则新建
      if (oOBJECT || oEMBED) {
        if (isIE) {
          LODOP = oOBJECT
        } else {
          LODOP = oEMBED
        }
      } else if (!CreatedOKLodopObject) {
        LODOP = document.createElement('object')
        LODOP.setAttribute('width', 0)
        LODOP.setAttribute('height', 0)
        LODOP.setAttribute('style', 'position:absolute;left:0px;top:-100px;width:0px;height:0px;')
        if (isIE) {
          LODOP.setAttribute('classid', 'clsid:2105C259-1E0C-4534-8141-A753534CB4CA')
        } else {
          LODOP.setAttribute('type', 'application/x-print-lodop')
        }
        document.documentElement.appendChild(LODOP)
        CreatedOKLodopObject = LODOP
      } else {
        LODOP = CreatedOKLodopObject
      }
      // Lodop插件未安装时提示下载地址
      if ((!LODOP) || (!LODOP.VERSION)) {
        if (ua.indexOf('Chrome') >= 0) {
          document.body.insertAdjacentHTML('afterBegin', strHtmChrome)
        }
        if (ua.indexOf('Firefox') >= 0) {
          document.body.insertAdjacentHTML('afterBegin', strHtmFireFox)
        }
        document.body.insertAdjacentHTML('afterBegin', (is64IE ? strHtm64Install : strHtmInstall))
        return LODOP
      }
    }
    if (LODOP && LODOP.VERSION < '6.2.2.6') {
      if (!needCLodop()) {
        document.body.insertAdjacentHTML('afterBegin', (is64IE ? strHtm64Update : strHtmUpdate))
      }
    }
    // 如下空白位置适合调用统一功能(如注册语句、语言选择等)
    return LODOP
  } catch (err) {
    alert('getLodop出错:' + err)
  }
}

// 引用CLodop的主JS,用双端口8000和18000(以防其中一个被占)
if (needCLodop()) {
  const head = document.head || document.getElementsByTagName('head')[0] || document.documentElement

  const JS1 = document.createElement('script')
  JS1.src = 'http://localhost:8000/CLodopfuncs.js?priority=1'
  head.insertBefore(JS1, head.firstChild)

  const JS2 = document.createElement('script')
  JS2.src = 'http://localhost:18000/CLodopfuncs.js?priority=0'
  head.insertBefore(JS2, head.firstChild)

  CLodopIsLocal = !!((JS1.src + JS2.src).match(/\/\/localho|\/\/127.0.0./i))

  if (JS1.attachEvent) {
    CLodopJsState = 'loading'
    const onChange = function () {
      if (window.event.srcElement.readyState === 'loaded') {
        CLodopJsState = 'complete'
      }
    }
    JS1.attachEvent('onreadystatechange', onChange)
    JS2.attachEvent('onreadystatechange', onChange)
  }
}
