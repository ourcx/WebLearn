package router

import (
	"example/controllers"
	"example/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	auth := gin.Default()
	v1 := auth.Group("/api/v1")
	{
		v1.POST("/login", controllers.Login)
		v1.POST("/register", controllers.Register)
	}

	api := auth.Group("/api")
	api.GET("/exchangeRates", controllers.GetExchangeRate)
	//不用登录和注册的功能
	api.Use(middlewares.AuthMiddleware())
	{
		//要注册和登录的功能
		api.POST("/exchangeRates", controllers.CreateExchangeRate)
		api.POST("/articles", controllers.CreateAeticle)
		api.GET("/articles", controllers.GetArticles)
		api.GET("articles/:id", controllers.GetArticleByID)
		api.POST("/articles/:id/like", controllers.LikeArticle)
		api.GET("/articles/:id/like", controllers.GetArticleLikes)
	}

	return auth
}
