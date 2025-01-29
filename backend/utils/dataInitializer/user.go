package dataInitializer

import (
	"context"

	db "github.com/ChocolateAceCream/telescope/backend/db/sqlc"
	"github.com/ChocolateAceCream/telescope/backend/singleton"
	"github.com/ChocolateAceCream/telescope/backend/utils"
	"go.uber.org/zap"
)

type user struct{}

const InitUserOrder = internalOrder + 1

func init() {
	Register(InitUserOrder, &user{})
}

func (u *user) Name() string {
	return "user"
}

func (u *user) Init(ctx context.Context) (next context.Context, err error) {
	err = singleton.DB.InitUsers(ctx, singleton.Config.DB.Password)
	if err != nil {
		singleton.Logger.Error("init user failed", zap.Error(err))
		return
	}
	next = ctx
	return
}

func (u *user) VerifyData(ctx context.Context) (ok bool) {
	ok, _ = singleton.DB.VerifyUserCredentials(ctx, db.VerifyUserCredentialsParams{
		Username: "admin",
		Password: utils.Sha256(singleton.Config.DB.Password),
	})
	return
}
