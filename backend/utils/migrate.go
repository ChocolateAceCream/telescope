package utils

import (
	"github.com/ChocolateAceCream/telescope/backend/singleton"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"go.uber.org/zap"
)

func Migrate() (err error) {
	m, err := migrate.New(
		singleton.Config.DB.MigrationFileDir,
		singleton.Config.DB.Source,
	)
	if err != nil {
		singleton.Logger.Error("cannot create migration", zap.Error(err))
		return
	}
	if err = m.Up(); err != nil && err != migrate.ErrNoChange {
		singleton.Logger.Error("cannot migrate", zap.Error(err))
		return
	}
	singleton.Logger.Info("migrate success")
	return nil
}
