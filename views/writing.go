package views

import (
	"net/http"
	"rm-go-blog/common"
	"rm-go-blog/service"
)

func (*HTMLApi) Writing(w http.ResponseWriter, r *http.Request) {
	writingTemplate := common.Template.Writing
	//
	wr := service.Writing()
	writingTemplate.WriteDate(w, wr)
}
