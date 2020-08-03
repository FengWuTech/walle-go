// eslint-disable-next-line
import { UserLayout, BasicLayout, RouteView, BlankLayout, PageView } from '@/layouts'
// import { dashboard } from '@/core/icons'
// import Role from './modules/Role'
// import Resource from './modules/Resource'
// import Action from './modules/Action'
// import User from './modules/User'
import { action } from '@/core/icons'

export const asyncRouterMap = [

  {
    path: '/',
    component: BasicLayout,
    meta: { title: '首页' },
    redirect: '/index',
    children: [
      {
        path: '/index',
        name: 'Index',
        meta: { title: '首页', keepAlive: false, icon: 'home', permission: ['dashboard'] },
        component: () => import('@/views/home/index')
      },
      // {
      //   path: '/space/index',
      //   name: 'SpaceList',
      //   meta: { title: '空间管理', keepAlive: false, icon: 'robot', permission: ['dashboard'] },
      //   component: () => import('@/views/space/list')
      // },
      // {
      //   path: '/space/add',
      //   name: 'SpaceCreate',
      //   hidden: true,
      //   component: () => import('@/views/space/addOrEdit'),
      //   meta: { title: '添加空间', keepAlive: true, permission: ['dashboard'] }
      // },
      // {
      //   path: '/space/:id',
      //   name: 'SpaceEdit',
      //   hidden: true,
      //   component: () => import('@/views/space/addOrEdit'),
      //   meta: { title: '编辑空间', keepAlive: true, permission: ['dashboard'] }
      // },
      {
        path: '/user/index',
        name: 'UserList',
        meta: { title: '用户管理', keepAlive: false, icon: 'user', permission: ['dashboard'] },
        component: () => import('@/views/user/list')
      },
      {
        path: '/user/add',
        name: 'UserCreate',
        hidden: true,
        component: () => import('@/views/user/addOrEdit'),
        meta: { title: '添加用户', keepAlive: true, permission: ['dashboard'] }
      },
      {
        path: '/user/:id',
        name: 'UserEdit',
        hidden: true,
        component: () => import('@/views/user/addOrEdit'),
        meta: { title: '编辑用户', keepAlive: true, permission: ['dashboard'] }
      },
      {
        path: '/project/index',
        name: 'ProjectList',
        component: PageView,
        meta: { title: '项目中心', keepAlive: false, icon: 'project', permission: ['dashboard'] },
        redirect: '/environment/index',
        children: [
          {
            path: '/environment/index',
            name: 'EnvironmentList',
            meta: { title: '环境管理', keepAlive: false, permission: ['dashboard'] },
            component: () => import('@/views/environment/list')
          },
          {
            path: '/environment/add',
            name: 'EnvironmentCreate',
            hidden: true,
            component: () => import('@/views/environment/addOrEdit'),
            meta: { title: '添加环境', keepAlive: true, permission: ['dashboard'] }
          },
          {
            path: '/environment/:id',
            name: 'EnvironmentEdit',
            hidden: true,
            component: () => import('@/views/environment/addOrEdit'),
            meta: { title: '编辑环境', keepAlive: true, permission: ['dashboard'] }
          },
          {
            path: '/server/index',
            name: 'ServerList',
            meta: { title: '服务器管理', keepAlive: false, permission: ['dashboard'] },
            component: () => import('@/views/server/list')
          },
          {
            path: '/server/add',
            name: 'ServerCreate',
            hidden: true,
            component: () => import('@/views/server/addOrEdit'),
            meta: { title: '添加服务器', keepAlive: true, permission: ['dashboard'] }
          },
          {
            path: '/server/:id',
            name: 'ServerEdit',
            hidden: true,
            component: () => import('@/views/server/addOrEdit'),
            meta: { title: '编辑服务器', keepAlive: true, permission: ['dashboard'] }
          },
          {
            path: '/project/index',
            name: 'ProjectList',
            meta: { title: '项目管理', keepAlive: false, permission: ['dashboard'] },
            component: () => import('@/views/project/list')
          },
          {
            path: '/project/add',
            name: 'ProjectCreate',
            hidden: true,
            component: () => import('@/views/project/addOrEdit'),
            meta: { title: '添加项目', keepAlive: true, permission: ['dashboard'] }
          },
          {
            path: '/project/:id',
            name: 'ProjectEdit',
            hidden: true,
            component: () => import('@/views/project/addOrEdit'),
            meta: { title: '编辑项目', keepAlive: true, permission: ['dashboard'] }
          },
          {
            path: '/project/read/:id',
            name: 'ProjectReadEdit',
            hidden: true,
            component: () => import('@/views/project/addOrEdit'),
            meta: { title: '查看项目', keepAlive: true, permission: ['dashboard'] }
          },
          {
            path: '/project/detection/:id',
            name: 'ProjectDetectionEdit',
            hidden: true,
            component: () => import('@/views/project/detection'),
            meta: { title: '项目检测', keepAlive: true, permission: ['dashboard'] }
          }
        ]
      },
      {
        path: '/deploy/index',
        name: 'DeployList',
        meta: { title: '部署管理', keepAlive: false, icon: 'hdd', permission: ['dashboard'] },
        component: () => import('@/views/task/list')
      },
      {
        path: '/task/create',
        name: 'TaskList',
        component: PageView,
        meta: { title: '部署管理', keepAlive: false, icon: action, permission: ['dashboard'] },
        // redirect: '/deploy/index',
        hidden: true,
        children: [
          {
            path: '/deploy/index',
            name: 'DeployList',
            hidden: true,
            meta: { title: '部署管理', keepAlive: false, icon: action, permission: ['dashboard'] },
            component: () => import('@/views/task/list')
          }, {
            path: '/task/create',
            name: 'taskCreate',
            hidden: true,
            component: () => import('@/views/task/create'),
            meta: { title: '添加上线单', keepAlive: true, permission: ['dashboard'] }
          },
          {
            path: '/task/create/:id',
            name: 'taskIdCreate',
            hidden: true,
            component: () => import('@/views/task/addOrEdit'),
            meta: { title: '添加上线单', keepAlive: true, permission: ['dashboard'] }
          },
          {
            path: '/task/edit/:id',
            name: 'taskEditCreate',
            hidden: true,
            component: () => import('@/views/task/addOrEdit'),
            meta: { title: '编辑上线单', keepAlive: true, permission: ['dashboard'] }
          },
          {
            path: '/task/deploy/:id',
            name: 'taskDeployCreate',
            hidden: true,
            component: () => import('@/views/task/deploy'),
            meta: { title: '上线单', keepAlive: true, permission: ['dashboard'] }
          }
        ]
      }

      // User,
      // Role,
      // Resource,
      // Action
    ]
  },
  {
    path: '*', redirect: '/404', hidden: true
  }
]

/**
 * 基础路由
 * @type { *[] }
 */
export const constantRouterMap = [
  // {
  //   path: '/',
  //   name: '/',
  //   redirect: '/apps/list',
  //   component: BlankLayout,
  //   children: [
  //     {
  //       path: '/apps/list',
  //       name: 'AppsList',
  //       meta: { title: '应用管理', keepAlive: false, icon: action,permission: ['dashboard'] },
  //       component: () => import(/* webpackChunkName: "fail" */ '@/views/apphome/app')
  //     },
  //     {
  //       path: '/apps/:id',
  //       name: 'AppsEdit',
  //       component: () => import(/* webpackChunkName: "fail" */ '@/views/apphome/addOrEdit'),
  //       meta: { title: '编辑应用', keepAlive: false, icon: action,permission: ['dashboard'] }
  //     },
  //     {
  //       path: '/apps/add',
  //       name: 'AppCreate',
  //       hidden: true,
  //       component: () => import('@/views/app/addOrEdit'),
  //       meta: { title: '添加应用', keepAlive: false, icon: action,permission: ['dashboard'] }
  //     }
  //   ]
  // },
  {
    path: '/login',
    name: 'login',
    component: () => import('@/views/login/index')
  },

  {
    path: '/404',
    component: () => import(/* webpackChunkName: "fail" */ '@/views/exception/404')
  },

  {
    path: '/403',
    component: () => import(/* webpackChunkName: "fail" */ '@/views/exception/403')
  }

]
