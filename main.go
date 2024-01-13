package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/wzliangzZ/ZZBatle/apps/login/handler"
)

func init() {
	// 加载.env文件
	godotenv.Load(".env")
}

func main() {
	r := gin.New()
	r.LoadHTMLGlob("./apps/common/templates/index.tmpl")
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.tmpl", gin.H{
			"title": "index",
		})
	})
	r.POST("/Hall", func(ctx *gin.Context) {
		handler.NewLoginHandler().HandlerFunc(ctx)
	})
	r.Run()
}
