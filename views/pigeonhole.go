package views

import (
	"net/http"
	"rm-go-blog/common"
	"rm-go-blog/service"
)

func (*HTMLApi) Pigeonhole(w http.ResponseWriter, r *http.Request) {
	pigeonholeTemp := common.Template.Pigeonhole
	pResp := service.FindPostPig()
	pigeonholeTemp.WriteDate(w, pResp)
}
