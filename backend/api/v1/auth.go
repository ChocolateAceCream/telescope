package apiV1

import (
	"net/http"

	"github.com/ChocolateAceCream/telescope/backend/model/request"
	"github.com/ChocolateAceCream/telescope/backend/model/response"
	"github.com/ChocolateAceCream/telescope/backend/singleton"
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
	response.OkWithFullDetails(c, user, "success")
}

func (a *AuthApi) GoogleLogin(c *gin.Context) {
	code := c.Query("code")
	singleton.Logger.Info("code: ", zap.String("code", code))
	if code == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Code not found"})
		return
	}

	// Exchange the authorization code for an access token
	err := authService.ExchangeCodeForToken(c, code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to exchange code for token"})
		return
	}
	c.Redirect(http.StatusTemporaryRedirect, singleton.Config.OAuth.Google.FrontendBaseUrl)
}

func (a *AuthApi) RefreshToken(c *gin.Context) {
	err := authService.RefreshToken(c)
	if err != nil {
		response.FailWithUnauthorized(c, "error.invalid.credentials")
		return
	}
	response.OkWithMessage(c, "success")
}
