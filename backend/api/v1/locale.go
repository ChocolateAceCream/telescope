package apiV1

import (
	"github.com/ChocolateAceCream/telescope/backend/model/response"
	"github.com/ChocolateAceCream/telescope/backend/singleton"
	"github.com/gin-gonic/gin"
)

type LocaleApi struct{}

func (l *LocaleApi) LoadTranslation(c *gin.Context) {
	mapper, err := localeService.LoadTranslationMapper(c)
	if err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}
	singleton.Translation = mapper
	response.OkWithMessage(c, "success")
}
