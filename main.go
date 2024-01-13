package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/wzliangzZ/ZZBatle/apps/login/handler"
)

func init() {
	// 加载.env文件
	godotenv.Load(".env")
}

func main() {
	r := gin.New()
	r.LoadHTMLGlob("./apps/common/templates/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.tmpl", gin.H{
			"title": "index",
		})
	})
	r.POST("/Hall", func(c *gin.Context) {
		handler.NewLoginHandler().HandlerFunc(c)
		c.HTML(200, "hall.tmpl", gin.H{
			"title": "hall",
		})
		logrus.Infoln("login success, enter game hall")
	})
	r.Run()
}
