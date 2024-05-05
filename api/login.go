package api

import (
	"net/http"
	"rm-go-blog/common"
	"rm-go-blog/context"
	"rm-go-blog/service"
)

func (*Api) Login(w http.ResponseWriter, r *http.Request) {
	//拿到json里的用户名和密码 返回结果json
	json := common.GetRequestJsonParam(r)
	username := json["username"].(string)
	password := json["passwd"].(string)
	loginResp, err := service.Login(username, password)
	if err != nil {
		common.Error(w, err)
		return
	}
	common.Success(w, loginResp)
}
func (*Api) LoginNew(ctx *context.RmContext) {
	w := ctx.W
	r := ctx.Request
	//拿到json里的用户名和密码 返回结果json
	json := common.GetRequestJsonParam(r)
	username := json["username"].(string)
	password := json["passwd"].(string)
	loginResp, err := service.Login(username, password)
	if err != nil {
		common.Error(w, err)
		return
	}
	common.Success(w, loginResp)
}
