//学习了！中间件真是个好东西（赞赏
package auth

import (
	gin "github.com/gin-gonic/gin"
	"go-picbed/database"
	"go-picbed/model"
	"net/http"
)

func JwtMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		// get authorization header
		//oauth2.0 "bearer "[7]
		TString := c.GetHeader("Authorization")[7:]
		token, claims, err := TokenParse(TString)

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "unauthorized"})
		}

		UserId := claims.UserId
		DB := database.GetDB()
		var user model.User
		DB.First(&user, UserId)

		// not registered
		if user.ID == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "unauthorized"})
		}

		// write
		c.Set("user", user)
		c.Next()
	}
}
