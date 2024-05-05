package service

import (
	"rm-go-blog/config"
	"rm-go-blog/dao"
	models "rm-go-blog/modles"
)

func FindPostPig() models.PigeonholeResp {

	//查询所有文章分月份、查询所有分类
	category, _ := dao.GetAllCategory()
	posts, _ := dao.GetAllPost()
	pigMap := make(map[string][]models.Post)
	for _, post := range posts {
		at := post.CreateAt
		month := at.Format("2006-01")
		pigMap[month] = append(pigMap[month], post)
	}
	return models.PigeonholeResp{
		config.Cfg.Viewer,
		config.Cfg.System,
		category,
		pigMap,
	}
}
