package views

import (
	"net/http"
	"rm-go-blog/common"
	"rm-go-blog/context"
	"rm-go-blog/service"
)

func (*HTMLApi) Writing(w http.ResponseWriter, r *http.Request) {
	writingTemplate := common.Template.Writing
	//
	wr := service.Writing()
	writingTemplate.WriteDate(w, wr)
}
func (*HTMLApi) WritingNew(ctx *context.RmContext) {
	writingTemplate := common.Template.Writing
	//
	wr := service.Writing()
	writingTemplate.WriteDate(ctx.W, wr)
}
