package service

import (
	"errors"

	db "github.com/ChocolateAceCream/telescope/backend/db/sqlc"
	"github.com/ChocolateAceCream/telescope/backend/model/request"
	"github.com/gin-gonic/gin"
)

type AuthService struct{}

func (a *AuthService) Login(c *gin.Context, payload request.LoginRequest) (user db.AUser, err error) {
	user, err = userDao.GetUserByUsername(c, payload.Username)
	if err != nil {
		return
	}
	if user.Password != payload.Password {
		err = errors.New("error.invalid.credentials")
	}
	return
}
