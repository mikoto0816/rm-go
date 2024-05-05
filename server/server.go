package server

import (
	"log"
	"net/http"
	"rm-go-blog/router"
)

type RmServer struct {
}

var App = &RmServer{}

func (*RmServer) Start(ip, port string) {
	//http监听
	server := http.Server{
		Addr: ip + ":" + port,
	}
	//请求路由
	router.Router()
	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
