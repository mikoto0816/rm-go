package api

import (
	"net/http"
)

var API = Api{}

type Api struct {
}

func (a Api) Post(writer http.ResponseWriter, request *http.Request) {

}
