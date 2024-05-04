package views

import (
	"errors"
	"log"
	"net/http"
	"rm-go-blog/common"
	"rm-go-blog/service"
	"strconv"
)

func (*HTMLApi) Index(w http.ResponseWriter, r *http.Request) {
	index := common.Template.Index
	//处理分页
	if err := r.ParseForm(); err != nil {
		log.Println("表单获取失败")
		index.WriteError(w, errors.New("系统错误，请联系管理员"))
		return
	}
	pageStr := r.Form.Get("page")
	page := 1
	if pageStr != "" {
		page, _ = strconv.Atoi(pageStr)
	}
	pageSize := 10
	hr, err := service.GetAllHomeInfo(page, pageSize)
	if err != nil {
		log.Println("首页获取数据出错：", err)
		index.WriteError(w, errors.New("系统错误，请联系管理员"))
	}
	index.WriteDate(w, hr)

}