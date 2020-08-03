package main

import (
	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
	engineio "github.com/googollee/go-engine.io"
	"github.com/googollee/go-engine.io/transport"
	"github.com/googollee/go-engine.io/transport/polling"
	"github.com/googollee/go-engine.io/transport/websocket"
	socketio "github.com/googollee/go-socket.io"
	"net/http"
	"walle-go/conf"
	"walle-go/logger"
	"walle-go/models"
	"walle-go/router"
	wsio "walle-go/socketio"
)

/*
 	==js==

	<script src="https://cdn.socket.io/socket.io-1.2.0.js"></script>
    <script>
      var socket = io.connect("http://localhost/walle")
      socket.on('construct', function(msg){
        console.log(msg);
      });
	</script>

 	==console==
    手动触发 socket.emit('open', {"project_id":3})
*/

/*
	==nodejs===

	var socket = require('socket.io-client')('ws://192.168.251.226:8001',{ transports: ['websocket'] });
	socket.on('connect', function(data){
		console.log('data',data)
		socket.emit('notice', '222')
	});
	socket.on('reply', function(data){
		console.log('event',data)
	});
*/
func main() {
	// 初始化日志
	logger.Setup()
	models.Setup(true)
	var c conf.Conf
	c.GetConf()

	sentry.Init(sentry.ClientOptions{
		Dsn: c.SentryDsn,
		AttachStacktrace: true,
	})

	pt := polling.Default

	wt := websocket.Default
	wt.CheckOrigin = func(req *http.Request) bool {
		return true
	}

	server, err := socketio.NewServer(&engineio.Options{
		Transports: []transport.Transport{
			pt,
			wt,
		},
	})
	if err != nil {
		logger.Fatal(err)
	}

	server.OnConnect("/", func(s socketio.Conn) error {
		logger.Debug("connected:", s.ID())
		return nil
	})

	server.OnDisconnect("/", func(so socketio.Conn, reason string) {
		logger.Debug("closed", reason)
	})

	server.OnError("/", func(s socketio.Conn, e error) {
		//sentry.CaptureException(e)
		logger.Error("meet error:", e)
	})


	//把处理转到wsio中
	processor := wsio.WebSocketIO{Namespace: "/walle",}
	processor.Setup(server)

	go func() {
		err := server.Serve()
		sentry.CaptureException(err)
	}()
	defer server.Close()

	routers := router.InitRouter()

	routers.GET("/socket.io/*any", gin.WrapH(server))
	routers.POST("/socket.io/*any", gin.WrapH(server))
	routers.Run(":5000")
	//listener := &http.Server{
	//	Addr:           ":5000",
	//	Handler:        routers,
	//	ReadTimeout:    60,
	//	WriteTimeout:   60,
	//	MaxHeaderBytes: 1 << 20,
	//}
	//logger.Info("Serving at localhost:5000...")
	//logger.Fatal(listener.ListenAndServe())
}
