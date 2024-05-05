package models

type LoginResp struct {
	Token    string   `json:"token"`
	UserInfo UserInfo `json:"userInfo"`
}
