Walle-Go
========

鉴于python版本的 [walle-web](https://github.com/meolu/walle-web) 不再更新，我们有使用其功能的强烈需求，就使用golang进行了重写，确保了项目基本功能可用性，并对前端UI进行了优化。

如何跑起来？
---------

1. 下载代码
1. 安装后端依赖 `go mod download` ，从配置模版 `conf/conf.yaml.template` 复制出 `conf/conf.yaml`，修改为真实可用的配置，启动后端服务 `go run main.go` ， API接口和socket.io服务默认都运行在 `5000` 端口。
1. 安装前端依赖

```
cd www/
npm i
npm run dev
```
1. 修改www前端项目目录下，vue.config.js 文件，修改成后端接口！
```
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
        onProxyRes (proxyRes, req, res) {}
      },
```
