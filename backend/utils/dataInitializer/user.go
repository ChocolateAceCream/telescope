package dataInitializer

import (
	"context"

	db "github.com/ChocolateAceCream/telescope/backend/db/sqlc"
	"github.com/ChocolateAceCream/telescope/backend/singleton"
	"github.com/ChocolateAceCream/telescope/backend/utils"
	"github.com/jackc/pgx/v5"
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

func (u *user) Init(c context.Context) (next context.Context, err error) {
	tx, err := singleton.DB.BeginTx(c, pgx.TxOptions{})
	if err != nil {
		singleton.Logger.Error("begin tx failed", zap.Error(err))
		return
	}
	defer tx.Rollback(c)

	txQuery := singleton.Query.WithTx(tx)
	err = txQuery.InitUsers(c)
	if err != nil {
		singleton.Logger.Error("init user failed", zap.Error(err))
		return
	}
	err = txQuery.InitPasswordLogins(c, singleton.Config.DB.Password)
	if err != nil {
		singleton.Logger.Error("init password login failed", zap.Error(err))
		return
	}
	err = tx.Commit(c)
	if err != nil {
		singleton.Logger.Error("init user and password login commit tx failed", zap.Error(err))
	}
	next = c
	return
}

func (u *user) VerifyData(ctx context.Context) (ok bool) {
	ok, _ = singleton.Query.VerifyUserCredentials(ctx, db.VerifyUserCredentialsParams{
		Email:    "admin@admin.com",
		Password: utils.Sha256(singleton.Config.DB.Password),
	})
	return
}
