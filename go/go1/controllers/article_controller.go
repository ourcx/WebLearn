package controllers

import (
	"encoding/json"
	"errors"
	"example/global"
	"example/models"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"gorm.io/gorm"
	"net/http"
	"time"
)

var cacheKey = "article"

// 创建文章
func CreateAeticle(c *gin.Context) {
	var article models.Article
	if err := c.ShouldBind(&article); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := global.Db.AutoMigrate(article); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := global.Db.Create(&article).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"article": article})

}

func GetArticles(c *gin.Context) {
	cacheData, err := global.RedisDB.Get(cacheKey).Result()
	//做了一个旁路缓存机制，可以先从mysql拿到相关的数据，然后缓存在redis里面
	if errors.Is(err, redis.Nil) {
		var articles []models.Article

		if err := global.Db.Find(&articles).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		marshal, err := json.Marshal(articles)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		global.RedisDB.Set(cacheKey, string(marshal), 10*time.Minute)

	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	} else {
		var articles []models.Article
		err := json.Unmarshal([]byte(cacheData), &articles)
		if err != nil {
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": articles})
	}
}

func GetArticleByID(c *gin.Context) {
	id := c.Param("id")
	//param可以拿到url的参数
	var articles []models.Article

	if err := global.Db.Where("id=?", id).Find(&articles, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"data": nil})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"data": nil})
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": articles})
}
