package main

import (
	"encoding/json"
	"log"
	"net/http"
	"rm-go-blog/common"
	"rm-go-blog/router"
)

type JsonData struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func init() {
	//模板加载
	common.LoadTemplate()
}
func main() {
	//http监听
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		var data JsonData
		data.Title = "测试页面"
		data.Description = "我的第一个博客页面"
		jsonStr, _ := json.Marshal(data)
		w.Write(jsonStr)
	})

	//请求路由
	router.Router()
	//监听server
	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
