package views

import (
	"net/http"
	"rm-go-blog/common"
	"rm-go-blog/context"
	"rm-go-blog/service"
)

func (*HTMLApi) Pigeonhole(w http.ResponseWriter, r *http.Request) {
	pigeonholeTemp := common.Template.Pigeonhole
	pResp := service.FindPostPig()
	pigeonholeTemp.WriteDate(w, pResp)
}
func (*HTMLApi) PigeonholeNew(ctx *context.RmContext) {
	pigeonholeTemp := common.Template.Pigeonhole
	pResp := service.FindPostPig()
	pigeonholeTemp.WriteDate(ctx.W, pResp)
}
