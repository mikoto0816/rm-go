package models

import (
	"html/template"
	"io"
	"log"
	"time"
)

func IsODD(num int) bool {
	return num%2 == 0
}
func GetNextName(str []string, index int) string {
	return str[index+1]
}
func Date(layout string) string {
	return time.Now().Format(layout)
}
func DateDay(date time.Time) string {
	return date.Format("2006-01-02 15:04:05")
}

type TemplateBLog struct {
	*template.Template
}
type HtmlTemplate struct {
	Index      TemplateBLog
	Category   TemplateBLog
	Custom     TemplateBLog
	Detail     TemplateBLog
	Login      TemplateBLog
	Pigeonhole TemplateBLog
	Writing    TemplateBLog
}

func InitTemplate(templatePath string) (HtmlTemplate, error) {

	tp, err := readTemplate([]string{"index", "category", "custom", "detail", "login", "pigeonhole", "writing"},
		templatePath,
	)
	var htmlTemplate HtmlTemplate
	if err != nil {
		return htmlTemplate, err
	}
	htmlTemplate.Index = tp[0]
	htmlTemplate.Category = tp[1]
	htmlTemplate.Custom = tp[2]
	htmlTemplate.Detail = tp[3]
	htmlTemplate.Login = tp[4]
	htmlTemplate.Pigeonhole = tp[5]
	htmlTemplate.Writing = tp[6]
	return htmlTemplate, nil
}

func readTemplate(templates []string, path string) ([]TemplateBLog, error) {
	var tbs []TemplateBLog
	for _, view := range templates {
		viewName := view + ".html"
		t := template.New(viewName)
		//设置模板
		home := path + "home.html"
		header := path + "layout/header.html"
		footer := path + "layout/footer.html"
		personal := path + "layout/personal.html"
		post := path + "layout/post-list.html"
		pagination := path + "layout/pagination.html"
		t.Funcs(template.FuncMap{"isODD": IsODD, "getNextName": GetNextName, "date": Date, "dateDay": DateDay})
		t, err := t.ParseFiles(path+viewName, home, header, footer, personal, post, pagination)
		if err != nil {
			log.Printf("解析模板出错：", err)
			return nil, err
		}
		var tb TemplateBLog
		tb.Template = t
		tbs = append(tbs, tb)
	}
	return tbs, nil
}
func (t *TemplateBLog) WriteDate(w io.Writer, data interface{}) {
	err := t.Execute(w, data)
	if err != nil {
		_, err := w.Write([]byte(err.Error()))
		if err != nil {
			log.Println(err)
			return
		}
	}
}
func (t *TemplateBLog) WriteError(w io.Writer, err error) {
	if err != nil {
		_, err := w.Write([]byte(err.Error()))
		if err != nil {
			log.Println(err)
			return
		}
	}
}
