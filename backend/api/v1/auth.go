package apiV1

import (
	"github.com/ChocolateAceCream/telescope/backend/model/request"
	"github.com/ChocolateAceCream/telescope/backend/model/response"
	"github.com/ChocolateAceCream/telescope/backend/singleton"
	"github.com/ChocolateAceCream/telescope/backend/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AuthApi struct{}

func (a *AuthApi) Login(c *gin.Context) {
	var req request.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		singleton.Logger.Error("Failed to bind JSON", zap.Error(err))
		response.FailWithMessage(c, "error.missing.params")
		return
	}
	user, err := authService.Login(c, req)
	if err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}

	session := utils.GetSession(c)
	session.Set("user", user)
	response.OkWithFullDetails(c, user, "success")

}
