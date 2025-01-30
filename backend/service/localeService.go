package service

import "github.com/gin-gonic/gin"

type LocaleService struct{}

func (l *LocaleService) LoadTranslationMapper(c *gin.Context) (map[string]map[string]string, error) {
	records, err := LocaleDao.LoadTranslationMapper(c)
	if err != nil {
		return nil, err
	}
	mapper := make(map[string]map[string]string)
	for _, record := range records {
		if _, ok := mapper[record.Language]; !ok {
			mapper[record.Language] = make(map[string]string)
		}
		mapper[record.Language][record.RawMessage] = record.TranslatedMessage
	}
	return mapper, nil
}
