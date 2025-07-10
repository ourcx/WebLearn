package controllers

import (
	"example/global"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

func LikeArticle(c *gin.Context) {
	articleId := c.PostForm("id")
	likeKey := "article:" + articleId + ":likes"
	if err := global.RedisDB.Incr(likeKey).Err(); err != nil {
		c.JSON(500, gin.H{
			"code": 500,
			"msg":  err.Error(),
		})
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "ok,article",
	})
}

func GetArticleLikes(c *gin.Context) {
	articleId := c.PostForm("id")
	likeKey := "article:" + articleId + ":likes"
	likes, err := global.RedisDB.Get(likeKey).Result()

	if err == redis.Nil {
		likes = "0"
	} else if err != nil {
		c.JSON(500, gin.H{
			"code": 500,
			"msg":  err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"code":  200,
		"msg":   "ok,likes",
		"likes": likes,
	})
}
