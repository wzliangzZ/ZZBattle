package handler

import (
	"github.com/sirupsen/logrus"

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
	if err := c.Bind(&user); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		logrus.Errorf("bind user error: %v", err)
		return
	}
	loginOrRegister(&user)
	logrus.Infof("user %s is connected!", user.Name)
}

func loginOrRegister(user *models.LoginUser)  {
	
}



