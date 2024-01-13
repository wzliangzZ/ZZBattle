package handler

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/wzliangzZ/ZZBatle/apps/login/models"
)

// 用户从 / 页面登录注册时调用
type LoginHandler struct{}

func NewLoginHandler() *LoginHandler {
	return &LoginHandler{}
}

func (h *LoginHandler) HandlerFunc(c *gin.Context) {
	var user models.LoginUser
	c.Bind(&user)
	log.Printf("user %s was connected!", user.Name)

	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "ok",
	})
}
