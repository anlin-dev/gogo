package server

import (
	"github.com/gin-gonic/gin"
	"gogo/backend/init/db"
	"gogo/backend/init/log"
	"gogo/backend/init/router"
	"gogo/backend/init/viper"
	"net"
	"net/http"
)

func Start() {
	viper.Init()
	log.Init()
	db.Init()

	gin.SetMode(gin.DebugMode)

	rootRouter := router.Routes()

	tcpItem := "tcp4"

	server := &http.Server{
		Addr:    ":8080",
		Handler: rootRouter,
	}
	ln, err := net.Listen(tcpItem, server.Addr)
	if err != nil {
		panic(err)
	}
	type tcpKeepAliveListener struct {
		*net.TCPListener
	}

	if err := server.Serve(tcpKeepAliveListener{ln.(*net.TCPListener)}); err != nil {
		panic(err)
	}
}
