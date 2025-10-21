package main

import (
	"example/config"
	"example/router"
	"github.com/gin-gonic/gin"
)

func main() {
	config.InitConfig()

	//路由组
	r := router.SetupRouter()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	//要注意避免忘记设置端口
	err := r.Run(config.AppConfig.App.Port)
	if err != nil {
		return
	} // listen and serve on 0.0.0.0:8080
	//默认8080端口
}
