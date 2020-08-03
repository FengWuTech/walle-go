const path = require('path')
const webpack = require('webpack')
const createThemeColorReplacerPlugin = require('./config/plugin.config')
// 打包1kb以上的为.gz
const CompressionPlugin = require('compression-webpack-plugin')
// 依赖库分析工具
// const BundleAnalyzerPlugin = require('webpack-bundle-analyzer').BundleAnalyzerPlugin
function resolve (dir) {
  return path.join(__dirname, dir)
}

/**
 * check production or preview(pro.loacg.com only)
 * @returns {boolean}
 */
function isProd () {
  return process.env.NODE_ENV === 'production'
}

// const publicPath = process.env.NODE_ENV === 'production' && process.env.VUE_APP_PREVIEW !== 'true' ? 'https://pms-static.gmtech.top/fe/src/walle/dist/' : '/'

const assetsCDN = {
  css: [],
  // https://unpkg.com/browse/vue@2.6.10/
  js: [
    '//pms-static.gmtech.top/fe/src/npm/vue@2.6.10/dist/vue.min.js',
    '//pms-static.gmtech.top/fe/src/npm/vue-router@3.1.3/dist/vue-router.min.js',
    '//pms-static.gmtech.top/fe/src/npm/vuex@3.1.1/dist/vuex.min.js',
    '//pms-static.gmtech.top/fe/src/npm/axios@0.19.0/dist/axios.min.js'
  ]
}

// webpack build externals
const prodExternals = {
  vue: 'Vue',
  'vue-router': 'VueRouter',
  vuex: 'Vuex',
  axios: 'axios'
}

// vue.config.js
const vueConfig = {
  // publicPath,

  configureWebpack: {
    // webpack plugins
    plugins: [
      // Ignore all locale files of moment.js
      new webpack.IgnorePlugin(/^\.\/locale$/, /moment$/),
      new CompressionPlugin({
        test: /\.(js|css)(\?.*)?$/i, // 需要压缩的文件正则
        threshold: 10240, // 文件大小大于这个值时启用压缩
        deleteOriginalAssets: false// 压缩后保留原文件
      })
      // 依赖大小分析工具
      // new BundleAnalyzerPlugin()
    ],
    // if prod is on, add externals
    externals: isProd() ? prodExternals : {}
  },

  chainWebpack: (config) => {
    config.resolve.alias
      .set('@$', resolve('src'))

    const svgRule = config.module.rule('svg')
    svgRule.uses.clear()
    svgRule
      .oneOf('inline')
      .resourceQuery(/inline/)
      .use('vue-svg-icon-loader')
      .loader('vue-svg-icon-loader')
      .end()
      .end()
      .oneOf('external')
      .use('file-loader')
      .loader('file-loader')
      .options({
        name: 'assets/[name].[hash:8].[ext]'
      })

    // if prod is on
    // assets require on cdn
    if (isProd()) {
      config.plugin('html').tap(args => {
        args[0].cdn = assetsCDN
        return args
      })
    }
  },

  css: {
    loaderOptions: {
      less: {
        modifyVars: {
          // less vars，customize ant design theme
          'primary-color': '#5D6EA4'
          // 'link-color': '#F5222D',
          // 'border-radius-base': '4px'
        },
        // do not remove this line
        javascriptEnabled: true
      }
    }
  },

  devServer: {
    // development server port 8000
    port: 8000,
    // If you want to turn on the proxy, please remove the mockjs /src/main.jsL11
    proxy: {
      '/walle/': {
        target: 'http://127.0.0.1:5000/',
        ws: true,
        changeOrigin: true,
        onProxyReq (proxyReq, req, res) {
          const cookie = req.headers['cookie']
          if (cookie) {
            proxyReq.setHeader('cookie', cookie)
          }
        },
        onProxyRes (proxyRes, req, res) {

        }
      },
      '/socket.io/': {
        target: 'http://127.0.0.1:5000/',
        ws: true,
        changeOrigin: true,
        onProxyReq (proxyReq, req, res) {
          const cookie = req.headers['cookie']
          if (cookie) {
            proxyReq.setHeader('cookie', cookie)
          }
        },
        onProxyRes (proxyRes, req, res) {

        }
      },
      '/api/': {
        target: 'http://127.0.0.1:5000/',
        ws: false,
        changeOrigin: true,
        onProxyReq (proxyReq, req, res) {
          const cookie = req.headers['cookie']
          if (cookie) {
            proxyReq.setHeader('cookie', cookie)
          }
        },
        onProxyRes (proxyRes, req, res) {

        }
      }
    }
  },

  // disable source map in production
  productionSourceMap: false,
  lintOnSave: undefined,
  // babel-loader no-ignore node_modules/*
  transpileDependencies: []
}

// preview.pro.loacg.com only do not use in your production;
if (process.env.VUE_APP_PREVIEW === 'true') {
  console.log('VUE_APP_PREVIEW', true)
  // add `ThemeColorReplacer` plugin to webpack plugins
  vueConfig.configureWebpack.plugins.push(createThemeColorReplacerPlugin())
}

module.exports = vueConfig
