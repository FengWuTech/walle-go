import 'driver.js/dist/driver.min.css' // import driver.js css
import Driver from 'driver.js'

const driver = new Driver({
  className: 'fw-driver-class',
  allowClose: false,
  keyboardControl: false,
  doneBtnText: '完成', // Text on the final button
  closeBtnText: '不再显示', // Text on the close button for this step
  nextBtnText: '下一步', // Next button text for this step
  prevBtnText: '上一步' // Previous button text for this step
})

export default driver
