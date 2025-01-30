package utils

import (
	"context"
	"time"

	db "github.com/ChocolateAceCream/telescope/backend/db/sqlc"
	"github.com/ChocolateAceCream/telescope/backend/singleton"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

var DB *pgxpool.Pool

func InitDB() (err error) {
	ctx := context.Background()

	config, err := pgxpool.ParseConfig(singleton.Config.DB.Source)
	if err != nil {
		singleton.Logger.Error("cannot parse DB source", zap.Error(err))
		return
	}

	config.MaxConns = 10
	config.MaxConnIdleTime = time.Hour

	DB, err = pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		singleton.Logger.Error("cannot create DB connection pool", zap.Error(err))
		return
	}
	err = DB.Ping(ctx)
	if err != nil {
		singleton.Logger.Error("cannot ping DB", zap.Error(err))
		return
	}

	singleton.DB = db.New(DB)
	return
}

func InitTranslation() (err error) {
	records, err := singleton.DB.GetAllLocales(context.Background())
	if err != nil {
		singleton.Logger.Error("GetAllLocales failed", zap.Error(err))
		return
	}
	mapper := make(map[string]map[string]string)
	for _, record := range records {
		if _, ok := mapper[record.Language]; !ok {
			mapper[record.Language] = make(map[string]string)
		}
		mapper[record.Language][record.RawMessage] = record.TranslatedMessage
	}
	singleton.Translation = mapper
	return
}
