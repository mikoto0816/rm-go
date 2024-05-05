package main

import (
	"encoding/json"
	"net/http"
	"rm-go-blog/common"
	"rm-go-blog/config"
	"rm-go-blog/server"
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
	server.App.Start(config.Cfg.System.IP, config.Cfg.System.Port)

	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		var data JsonData
		data.Title = "测试页面"
		data.Description = "我的第一个博客页面"
		jsonStr, _ := json.Marshal(data)
		w.Write(jsonStr)
	})
}
