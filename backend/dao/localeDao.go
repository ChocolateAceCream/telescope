package dao

import (
	db "github.com/ChocolateAceCream/telescope/backend/db/sqlc"
	"github.com/ChocolateAceCream/telescope/backend/singleton"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type LocaleDao struct{}

func (l *LocaleDao) LoadTranslationMapper(c *gin.Context) (records []db.GetAllLocalesRow, err error) {
	records, err = singleton.DB.GetAllLocales(c)
	if err != nil {
		singleton.Logger.Error("GetAllLocales failed", zap.Error(err))
	}
	return
}
