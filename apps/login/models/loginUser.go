package models

// 登录的用户，用于装填参数
type LoginUser struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}
