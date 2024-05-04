package common

import (
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
