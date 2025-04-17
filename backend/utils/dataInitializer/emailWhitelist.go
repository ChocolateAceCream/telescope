package dataInitializer

import (
	"context"

	"github.com/ChocolateAceCream/telescope/backend/singleton"
	"go.uber.org/zap"
)

type emailWhitelist struct{}

const InitEmailWhitelistOrder = internalOrder + 1

func init() {
	Register(InitUserOrder, &emailWhitelist{})
}

func (e *emailWhitelist) Name() string {
	return "email_whitelist"
}

func (e *emailWhitelist) Init(c context.Context) (next context.Context, err error) {
	err = singleton.Query.InitEmailWhitelist(c)
	if err != nil {
		singleton.Logger.Error("failed to init email whitelist", zap.Error(err))
		return
	}
	next = c
	return
}

func (e *emailWhitelist) VerifyData(ctx context.Context) (ok bool) {
	ok, _ = singleton.Query.VerifyEmailWhitelistExist(ctx, "admin@admin.com")
	return
}
