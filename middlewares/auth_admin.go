package middlewares

import (
	"YourProjectName/helper"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// AuthAdminCheck is a middleware function that checks if the user is authenticated with admin role.
func AuthAdminCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		log.Println(auth)
		userClaim, err := helper.AnalyseToken(auth)
		if err != nil {
			c.Abort()
			c.JSON(http.StatusOK, gin.H{
				"code": http.StatusUnauthorized,
				"msg":  "当前未登录",
			})
			return
		}
		if userClaim == nil || userClaim.IsAdmin != 1 {
			c.Abort()
			c.JSON(http.StatusOK, gin.H{
				"code": http.StatusUnauthorized,
				"msg":  "未经授权",
			})
			return
		}
		c.Next()
	}
}
