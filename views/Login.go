package views

import (
	"net/http"
	"rm-go-blog/common"
	"rm-go-blog/config"
)

func (*HTMLApi) Login(w http.ResponseWriter, r *http.Request) {

	loginTemplate := common.Template.Login
	loginTemplate.WriteDate(w, config.Cfg.Viewer)
}
