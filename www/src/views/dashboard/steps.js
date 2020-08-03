// 首页引导
const steps = [
  {
    element: '#add_project',
    popover: {
      title: '添加项目',
      description: '添加一个项目，所有的操作都是在某一项目下进行的',
      position: 'bottom'
    }
  },
  {
    element: '#add_group',
    popover: {
      title: '添加组团',
      description: '在某个项目下添加组团，组团也是我们通常意义上的小区',
      position: 'bottom'
    }
  },
  {
    element: '#add_building',
    popover: {
      title: '添加楼宇',
      description: '在某个组团下添加楼宇，用于管理单元和房屋',
      position: 'bottom'
    }
  },
  {
    element: '#add_room',
    popover: {
      title: '添加房屋',
      description: '批量添加房屋可以快速得到标准统一的房屋设置',
      position: 'bottom'
    }
  },
  {
    element: '#contact_dev',
    popover: {
      title: '技术支持',
      description: '有任何问题，请联系客服解决'
    }
  }
]

export default steps
