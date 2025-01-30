package service

import (
	"errors"

	"github.com/ChocolateAceCream/telescope/backend/model/request"
	"github.com/ChocolateAceCream/telescope/backend/model/response"
	"github.com/gin-gonic/gin"
)

type AuthService struct{}

func (a *AuthService) Login(c *gin.Context, payload request.LoginRequest) (resp response.LoginResponse, err error) {
	user, err := userDao.GetUserByUsername(c, payload.Username)
	if err != nil {
		return
	}
	if user.Password != payload.Password {
		err = errors.New("error.invalid.credentials")
	}
	resp = response.LoginResponse{
		Username: user.Username,
		Email:    user.Email,
		Role:     user.Role,
	}
	return
}
