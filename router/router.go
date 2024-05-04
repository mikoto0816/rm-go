package router

import (
	"net/http"
	"rm-go-blog/api"
	"rm-go-blog/views"
)

func Router() {
	//返回json或静态资源
	http.HandleFunc("/", views.HTML.Index)
	http.HandleFunc("/api/post", api.API.Post)

	//静态资源映射
	http.Handle("/resource/", http.StripPrefix("/resource/",
		http.FileServer(http.Dir("public/resource/"))))
}
