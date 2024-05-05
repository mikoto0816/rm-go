package common

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"rm-go-blog/config"
	models "rm-go-blog/modles"
	"sync"
)

var Template models.HtmlTemplate

func LoadTemplate() {

	wait := sync.WaitGroup{}
	wait.Add(1)
	//加入协程提高性能
	go func() {
		var err error
		Template, err = models.InitTemplate(config.Cfg.System.CurrentDir + "/template/")
		if err != nil {
			panic(err)
		}
		wait.Done()
	}()
	wait.Wait()
}

func Success(w http.ResponseWriter, data interface{}) {
	var result models.Result
	result.Code = 200
	result.Error = ""
	result.Data = data
	resultJson, _ := json.Marshal(result)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	_, err := w.Write(resultJson)
	if err != nil {
		log.Println(err)
	}
}

func Error(w http.ResponseWriter, err error) {
	var result models.Result
	result.Code = 400
	result.Error = err.Error()
	resultJson, _ := json.Marshal(result)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	_, err = w.Write(resultJson)
	if err != nil {
		log.Println(err)
	}
}

func GetRequestJsonParam(r *http.Request) map[string]interface{} {
	var params map[string]interface{}
	all, _ := io.ReadAll(r.Body)
	_ = json.Unmarshal(all, &params)
	return params
}
