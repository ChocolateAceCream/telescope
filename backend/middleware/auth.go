/*
* @fileName auth.go
* @author Di Sheng
* @date 2025/02/08 16:15:28
* @description auth middleware, to check if user exist in session, if not, return 401, otherwise, store user in context
 */
package middleware

import (
	db "github.com/ChocolateAceCream/telescope/backend/db/sqlc"
	"github.com/ChocolateAceCream/telescope/backend/model/response"
	"github.com/ChocolateAceCream/telescope/backend/utils"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return AuthHandler
}

func AuthHandler(c *gin.Context) {
	user, err := utils.GetValueFromSessionByKey[db.AUser](c, "user")
	if err != nil {
		response.FailWithUnauthorized(c, "error.session.expired")
		c.Abort()
	}
	c.Set("user", user)
}
