package controllers

import (
	"example/global"
	"example/models"
	"example/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  err.Error(),
		})
		return
	}

	hashedPwd, err := utils.HashPassword(user.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 500,
			//服务器错误，无法完成加密
			"msg": err.Error(),
		})
		return
	}
	user.Password = hashedPwd
	//把明文转化为密文

	//生成jwt
	token, err := utils.GenerateJWT(user.Username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 500,
			"msg":  err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": gin.H{"token": token},
	})

	if err := global.Db.AutoMigrate(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 500,
			"msg":  err.Error(),
		})
		return
	}

	if err := global.Db.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

//对输入的密码进行加密和添加令牌和gin的一个函数

func Login(c *gin.Context) {
	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
		//json是序列化和反序列化的行为，把Username大写转小写
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
		})
		return
	}

	var user models.User

	if err := global.Db.Where("username = ?", input.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code": 401,
			"msg":  err.Error(),
			"tit":  "未注册",
		})
		return
	}
	//密码不一致
	if !utils.CheckPassword(user.Password, input.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code": 401,
			"tit":  "密码错误",
		})
		return
	}

	token, err := utils.GenerateJWT(input.Username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 500,
			"msg":  err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": gin.H{"token": token},
	})

}
