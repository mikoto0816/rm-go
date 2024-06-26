package service

import (
	"html/template"
	"log"
	"rm-go-blog/config"
	"rm-go-blog/dao"
	models "rm-go-blog/modles"
)

func GetPostDetail(pId int) (*models.PostDetail, error) {
	post, err := dao.GetPostDetail(pId)
	if err != nil {
		return nil, err
	}
	categoryName := dao.GetCategoryName(post.CategoryId)
	userName := dao.GetUserName(post.UserId)
	postMore := models.PostMore{
		post.Pid,
		post.Title,
		post.Slug,
		template.HTML(post.Content),
		post.CategoryId,
		categoryName,
		post.UserId,
		userName,
		post.ViewCount,
		post.Type,
		models.DateDay(post.CreateAt),
		models.DateDay(post.UpdateAt),
	}
	var postDetail = &models.PostDetail{
		config.Cfg.Viewer,
		config.Cfg.System,
		postMore,
	}
	return postDetail, nil
}

func Writing() (wr models.WritingResp) {
	wr.Title = config.Cfg.Viewer.Title
	wr.CdnURL = config.Cfg.System.CdnURL
	category, err := dao.GetAllCategory()
	if err != nil {
		log.Println(err)
		return
	}
	wr.Categorys = category
	return
}

func SavePost(post *models.Post) {
	dao.SavePost(post)
}

func UpdatePost(post *models.Post) {
	dao.UpdatePost(post)
}
func SearchPost(condtion string) []models.SearchResp {
	posts, _ := dao.Search(condtion)
	var searchResp []models.SearchResp
	for _, post := range posts {
		searchResp = append(searchResp, models.SearchResp{
			post.Pid,
			post.Title,
		})
	}
	return searchResp
}
