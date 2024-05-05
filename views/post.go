package views

import (
	"errors"
	"net/http"
	"rm-go-blog/common"
	"rm-go-blog/context"
	"rm-go-blog/service"
	"strconv"
	"strings"
)

func (*HTMLApi) Detail(w http.ResponseWriter, r *http.Request) {

	detail := common.Template.Detail
	//路径参数
	path := r.URL.Path
	prefix := strings.TrimPrefix(path, "/p/")
	pIdStr := strings.TrimSuffix(prefix, ".html")
	pId, err := strconv.Atoi(pIdStr)
	if err != nil {
		detail.WriteError(w, errors.New("无法识别此路径"))
	}
	postDetail, err := service.GetPostDetail(pId)
	if err != nil {
		detail.WriteError(w, errors.New("查询出错"))
		return
	}
	detail.WriteDate(w, postDetail)
}
func (*HTMLApi) DetailNew(ctx *context.RmContext) {

	detail := common.Template.Detail
	//路径参数
	path := ctx.Request.URL.Path
	prefix := strings.TrimPrefix(path, "/p/")
	pIdStr := strings.TrimSuffix(prefix, ".html")
	pId, err := strconv.Atoi(pIdStr)
	if err != nil {
		detail.WriteError(ctx.W, errors.New("无法识别此路径"))
	}
	postDetail, err := service.GetPostDetail(pId)
	if err != nil {
		detail.WriteError(ctx.W, errors.New("查询出错"))
		return
	}
	detail.WriteDate(ctx.W, postDetail)
}
