package service

import (
	"html/template"
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
