package middleware

import (
	"fmt"
	"go-clean-architecture/module/user/response"
	"go-clean-architecture/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		t := strings.Split(authHeader, " ")

		if len(t) < 2 {
			c.JSON(http.StatusUnauthorized, response.ErrorResponse{Message: "Not authorized"})
			c.Abort()
		}

		authToken := t[1]
		authorized, err := utils.IsAuthorized(authToken, secret)

		if authorized {
			userID, _ := utils.ExtractIDFromToken(authToken, secret)

			if err != nil {
				c.JSON(http.StatusUnauthorized, response.ErrorResponse{Message: err.Error()})
				c.Abort()
				return
			}

			c.Set("x-user-id", fmt.Sprint(userID))
			c.Next()
			return
		}

		c.JSON(http.StatusUnauthorized, response.ErrorResponse{Message: err.Error()})
		c.Abort()
	}
}
