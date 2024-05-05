package router

import (
	"net/http"
	"rm-go-blog/api"
	"rm-go-blog/views"
)

func Router() {
	//http.Handle("/", context.Context)
	//context.Context.Handler("/c/{id}", views.HTML.CategoryNew)
	//context.Context.Handler("/login", views.HTML.LoginNew)
	//context.Context.Handler("/p/", views.HTML.DetailNew)
	//context.Context.Handler("/writing/", views.HTML.WritingNew)
	//context.Context.Handler("/pigeonhole/", views.HTML.PigeonholeNew)
	//context.Context.Handler("/api/v1/post", api.API.SaveAndUpdatePostNew)
	//context.Context.Handler("/api/v1/post/", api.API.GetPostNew)
	//context.Context.Handler("/api/v1/post/search", api.API.SearchPostNew)
	//context.Context.Handler("/api/v1/login", api.API.LoginNew)
	//返回json或静态资源
	http.HandleFunc("/", views.HTML.Index)
	http.HandleFunc("/c/", views.HTML.Category)
	http.HandleFunc("/login", views.HTML.Login)
	http.HandleFunc("/p/", views.HTML.Detail)
	http.HandleFunc("/writing/", views.HTML.Writing)
	http.HandleFunc("/pigeonhole/", views.HTML.Pigeonhole)
	http.HandleFunc("/api/v1/post", api.API.SaveAndUpdatePost)
	http.HandleFunc("/api/v1/post/", api.API.GetPost)
	http.HandleFunc("/api/v1/post/search", api.API.SearchPost)
	http.HandleFunc("/api/v1/login", api.API.Login)

	//静态资源映射
	http.Handle("/resource/", http.StripPrefix("/resource/",
		http.FileServer(http.Dir("public/resource/"))))
}
