package views

import (
	"errors"
	"log"
	"net/http"
	"rm-go-blog/common"
	"rm-go-blog/context"
	"rm-go-blog/service"
	"strconv"
	"strings"
)

func (a *HTMLApi) Category(w http.ResponseWriter, r *http.Request) {

	category := common.Template.Category
	path := r.URL.Path
	prefix := strings.TrimPrefix(path, "/c/")
	cId, err := strconv.Atoi(prefix)
	if err != nil {
		category.WriteError(w, errors.New("无法识别此路径"))
	}
	//处理分页
	if err := r.ParseForm(); err != nil {
		log.Println("表单获取失败")
		category.WriteError(w, errors.New("系统错误，请联系管理员"))
		return
	}
	pageStr := r.Form.Get("page")
	if pageStr == "" {
		pageStr = "1"
	}
	page, _ := strconv.Atoi(pageStr)
	pageSize := 10
	//调用service返回数据
	categoryResponse, err := service.GetPostByCateGoryId(cId, page, pageSize)
	if err != nil {
		category.WriteError(w, err)
		return
	}
	//回写
	category.WriteDate(w, categoryResponse)
}

func (a *HTMLApi) CategoryNew(ctx *context.RmContext) {
	category := common.Template.Category
	cIdStr := ctx.GetPathVariable("id")
	w := ctx.W
	cId, err := strconv.Atoi(cIdStr)
	if err != nil {
		category.WriteError(w, errors.New("无法识别此路径"))
	}
	//处理分页
	pageStr, _ := ctx.GetForm("page")
	if pageStr == "" {
		pageStr = "1"
	}
	page, _ := strconv.Atoi(pageStr)
	pageSize := 10
	//调用service返回数据
	categoryResponse, err := service.GetPostByCateGoryId(cId, page, pageSize)
	if err != nil {
		category.WriteError(w, err)
		return
	}
	//回写
	category.WriteDate(w, categoryResponse)
}
