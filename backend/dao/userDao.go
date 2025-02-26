package dao

import (
	"errors"

	db "github.com/ChocolateAceCream/telescope/backend/db/sqlc"
	"github.com/ChocolateAceCream/telescope/backend/singleton"
	"github.com/ChocolateAceCream/telescope/backend/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserDao struct{}

func (u *UserDao) VerifyUserCredentials(c *gin.Context, email, password string) (err error) {
	ok, err := singleton.Query.VerifyUserCredentials(c, db.VerifyUserCredentialsParams{
		Email:    email,
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

func (u *UserDao) GetUserByEmail(c *gin.Context, email string) (user db.User, err error) {
	user, err = singleton.Query.GetUserByEmail(c, email)
	if err != nil {
		singleton.Logger.Error("GetUserByEmail failed", zap.Error(err))
		err = errors.New("failed to get user by email")
	}
	return
}

func (u *UserDao) GetUserByUsername(c *gin.Context, username string) (user db.User, err error) {
	user, err = singleton.Query.GetUserByUsername(c, username)
	if err != nil {
		singleton.Logger.Error("GetUserByUsername failed", zap.Error(err))
		err = errors.New("failed to get user by username")
	}
	return
}

func (u *UserDao) CreateGoogleLogin(c *gin.Context, payload db.GoogleLoginParams) (err error) {
	q := singleton.Query
	if tx, ok := utils.GetTx(c); ok {
		q = q.WithTx(tx)
	}
	err = q.GoogleLogin(c, payload)
	if err != nil {
		singleton.Logger.Error("Google login failed", zap.Error(err))
		err = errors.New("failed to create google login")
	}
	return
}

func (u *UserDao) CreateUser(c *gin.Context, payload db.CreateNewUserParams) (err error) {
	q := singleton.Query
	if tx, ok := utils.GetTx(c); ok {
		q = q.WithTx(tx)
	}
	err = q.CreateNewUser(c, payload)
	if err != nil {
		singleton.Logger.Error("CreateUser failed", zap.Error(err))
	}
	return
}
