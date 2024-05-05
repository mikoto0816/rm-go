package dao

import (
	"log"
	models "rm-go-blog/modles"
)

func GetUserName(userId int) string {

	var userName string
	row := DB.QueryRow("select user_name from blog_user where uid=?", userId)
	if row.Err() != nil {
		log.Println(row.Err())
	}
	_ = row.Scan(&userName)
	return userName
}

func GetUser(username, passwd string) *models.User {
	row := DB.QueryRow("select * from blog_user where user_name=? and passwd=?", username, passwd)
	if row.Err() != nil {
		log.Println(row.Err())
		return nil
	}
	var user = &models.User{}
	err := row.Scan(&user.Uid, &user.Username, &user.Passwd, &user.Avatar, &user.CreateAt, &user.UpdateAt)
	if err != nil {
		log.Println(err)
		return nil
	}
	return user
}
