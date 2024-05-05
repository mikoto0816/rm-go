package api

import (
	"errors"
	"log"
	"net/http"
	"rm-go-blog/common"
	"rm-go-blog/dao"
	models "rm-go-blog/modles"
	"rm-go-blog/service"
	"rm-go-blog/utils"
	"strconv"
	"strings"
	"time"
)

func (*Api) SaveAndUpdatePost(w http.ResponseWriter, r *http.Request) {
	//获取userId
	token := r.Header.Get("Authorization")
	_, claims, err := utils.ParseToken(token)
	if err != nil {
		log.Println("token err:", err)
		common.Error(w, errors.New("登录已过期"))
		return
	}
	userId := claims.Uid
	method := r.Method
	switch method {
	case http.MethodPost:
		json := common.GetRequestJsonParam(r)
		categoryId := json["categoryId"].(string)
		cId, _ := strconv.Atoi(categoryId)
		content := json["content"].(string)
		markdown := json["markdown"].(string)
		slug := json["slug"].(string)
		title := json["title"].(string)
		postType := json["type"].(float64)
		pType := int(postType)
		post := &models.Post{
			-1,
			title,
			slug,
			content,
			markdown,
			cId,
			userId,
			0,
			pType,
			time.Now(),
			time.Now(),
		}
		service.SavePost(post)
		common.Success(w, post)
	case http.MethodPut:
		//更新
		json := common.GetRequestJsonParam(r)
		categoryId := json["categoryId"].(float64)
		cId := int(categoryId)
		content := json["content"].(string)
		markdown := json["markdown"].(string)
		slug := json["slug"].(string)
		title := json["title"].(string)
		postType := json["type"].(float64)
		pIdFloat := json["pid"].(float64)
		pType := int(postType)
		pId := int(pIdFloat)
		post := &models.Post{
			pId,
			title,
			slug,
			content,
			markdown,
			cId,
			userId,
			0,
			pType,
			time.Now(),
			time.Now(),
		}
		service.UpdatePost(post)
		common.Success(w, post)

	}

}

func (*Api) GetPost(w http.ResponseWriter, r *http.Request) {

	path := r.URL.Path
	prefix := strings.TrimPrefix(path, "/api/v1/post/")
	pId, err := strconv.Atoi(prefix)
	if err != nil {
		common.Error(w, errors.New("“不识别此路径"))
	}
	post, err := dao.GetPostDetail(pId)
	if err != nil {
		common.Error(w, err)
		return
	}
	common.Success(w, post)
}
func (*Api) SearchPost(w http.ResponseWriter, r *http.Request) {

	_ = r.ParseForm()
	condition := r.Form.Get("val")
	searchReap := service.SearchPost(condition)
	common.Success(w, searchReap)
}
