package dao

import "log"

func GetUserName(userId int) string {

	var userName string
	row := DB.QueryRow("select user_name from blog_user where uid=?", userId)
	if row.Err() != nil {
		log.Println(row.Err())
	}
	_ = row.Scan(&userName)
	return userName
}
