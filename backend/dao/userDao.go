package dao

import (
	"errors"

	db "github.com/ChocolateAceCream/telescope/backend/db/sqlc"
	"github.com/ChocolateAceCream/telescope/backend/singleton"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserDao struct{}

func (u *UserDao) VerifyUserCredentials(c *gin.Context, username, password string) (err error) {
	ok, err := singleton.DB.VerifyUserCredentials(c, db.VerifyUserCredentialsParams{
		Username: username,
		Password: password,
	})
	if err != nil {
		singleton.Logger.Error("VerifyUserCredentials failed", zap.Error(err))
		err = errors.New("operation.failed")
		return
	}
	if !ok {
		err = errors.New("error.invalid.credentials")
		return
	}
	return
}

func (u *UserDao) GetUserByUsername(c *gin.Context, username string) (user db.AUser, err error) {
	user, err = singleton.DB.GetUserByUsername(c, username)
	if err != nil {
		singleton.Logger.Error("GetUserByUsername failed", zap.Error(err))
		err = errors.New("failed to get user by username")
	}
	return
}
