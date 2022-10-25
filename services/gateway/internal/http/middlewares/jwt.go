package middlewares

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"pad/services/common/jwt"
)

func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		bearToken := c.GetHeader("Authorization")
		strArr := strings.Split(bearToken, " ")
		if len(strArr) != 2 {
			c.JSON(http.StatusUnauthorized, fmt.Errorf("authorization token not found"))
			c.Abort()
			return
		}

		userID, err := jwt.ExtractUserFromToken(strArr[1])
		if err != nil {
			c.JSON(http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}
		c.Set("userID", userID)
		c.Next()
	}
}
