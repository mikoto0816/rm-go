package service

import (
	"html/template"
	"rm-go-blog/config"
	"rm-go-blog/dao"
	models "rm-go-blog/modles"
)

func GetAllHomeInfo(page, pageSize int) (*models.HomeResponse, error) {
	categorys, err := dao.GetAllCategory()
	//分页
	if err != nil {
		return nil, err
	}
	posts, err := dao.GetPostPage(page, pageSize)
	var postsMores []models.PostMore
	for _, post := range posts {
		//查询categoryName
		categoryName := dao.GetCategoryName(post.CategoryId)
		//查询userName
		userName := dao.GetUserName(post.UserId)
		content := []rune(post.Content)
		if len(content) > 100 {
			content = content[0:100]
		}
		postsMore := models.PostMore{
			post.Pid,
			post.Title,
			post.Slug,
			template.HTML(content),
			post.CategoryId,
			categoryName,
			post.UserId,
			userName,
			post.ViewCount,
			post.Type,
			models.DateDay(post.CreateAt),
			models.DateDay(post.UpdateAt),
		}
		postsMores = append(postsMores, postsMore)
	}
	total := dao.CountGetAllPost()
	pagesCount := (total-1)/10 + 1
	var pages []int
	for i := 0; i < pagesCount; i++ {
		pages = append(pages, i+1)
	}
	var hr = &models.HomeResponse{
		config.Cfg.Viewer,
		categorys,
		postsMores,
		total,
		page,
		pages,
		page != pagesCount,
	}
	return hr, nil
}
