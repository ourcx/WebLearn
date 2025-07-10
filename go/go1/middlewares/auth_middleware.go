package middlewares

import (
	"example/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": http.StatusUnauthorized,
			})
			c.Abort()
			//有错误就不继续进行后续操作
			return
		}

		username := utils.ParseJWT(token)

		c.Set("username", username)
		c.Next()
		//set和get结合起来，可以在中间件中跨中间件取值

	}
}
