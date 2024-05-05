package views

import (
	"net/http"
	"rm-go-blog/common"
	"rm-go-blog/config"
	"rm-go-blog/context"
)

func (*HTMLApi) Login(w http.ResponseWriter, r *http.Request) {

	loginTemplate := common.Template.Login
	loginTemplate.WriteDate(w, config.Cfg.Viewer)
}
func (*HTMLApi) LoginNew(ctx *context.RmContext) {

	loginTemplate := common.Template.Login
	loginTemplate.WriteDate(ctx.W, config.Cfg.Viewer)
}
