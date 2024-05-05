package router

import (
	"net/http"
	"rm-go-blog/api"
	"rm-go-blog/views"
)

func Router() {
	//返回json或静态资源
	http.HandleFunc("/", views.HTML.Index)
	http.HandleFunc("/c/", views.HTML.Category)
	http.HandleFunc("/login", views.HTML.Login)
	http.HandleFunc("/api/v1/post", api.API.SaveAndUpdatePost)
	http.HandleFunc("/api/v1/login", api.API.Login)
	http.HandleFunc("/p/", views.HTML.Detail)

	//静态资源映射
	http.Handle("/resource/", http.StripPrefix("/resource/",
		http.FileServer(http.Dir("public/resource/"))))
}
