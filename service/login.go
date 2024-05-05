package service

import (
	"errors"
	"rm-go-blog/dao"
	models "rm-go-blog/modles"
	"rm-go-blog/utils"
)

func Login(username, passwd string) (*models.LoginResp, error) {

	//md5
	passwd = utils.Md5Crypt(passwd, "rimomi")
	user := dao.GetUser(username, passwd)
	if user == nil {
		//登陆失败
		return nil, errors.New("账号密码不正确")
	}
	//生成token
	uid := user.Uid
	token, err := utils.Award(&uid)
	if err != nil {
		return nil, errors.New("token未生成")
	}
	var userInfo models.UserInfo
	userInfo.Uid = user.Uid
	userInfo.Username = user.Username
	userInfo.Avatar = user.Avatar
	loginResp := &models.LoginResp{
		token,
		userInfo,
	}
	return loginResp, nil
}
