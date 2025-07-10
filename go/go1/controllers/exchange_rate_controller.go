package controllers

import (
	"errors"
	"example/global"
	"example/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"time"
)

func CreateExchangeRate(c *gin.Context) {
	var exchangeRate models.ExchangeRate

	if err := c.ShouldBindJSON(&exchangeRate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
			"code":  http.StatusBadRequest,
		})
		return
	}
	exchangeRate.Data = time.Now()
	//返回本地时间
	//在数据库建表
	err := global.Db.AutoMigrate(&exchangeRate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
			"code":  http.StatusInternalServerError,
		})
		return
	}
	//创建相关的行
	if err := global.Db.Create(&exchangeRate).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
			"code":  http.StatusInternalServerError,
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"data": exchangeRate,
		"code": http.StatusCreated,
	})
}

func GetExchangeRate(c *gin.Context) {
	var exchangeRate models.ExchangeRate
	if err := global.Db.First(&exchangeRate).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
			"code":  http.StatusInternalServerError,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": exchangeRate,
		"code": http.StatusOK,
	})
}

func GetExchangeRates(c *gin.Context) {
	var exchangeRate []models.ExchangeRate

	if err := global.Db.Find(&exchangeRate).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": exchangeRate,
	})
}
