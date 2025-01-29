package apiV1

import (
	"strconv"

	db "github.com/ChocolateAceCream/telescope/backend/db/sqlc"
	"github.com/ChocolateAceCream/telescope/backend/lib"
	"github.com/ChocolateAceCream/telescope/backend/model/response"
	"github.com/ChocolateAceCream/telescope/backend/singleton"
	"github.com/ChocolateAceCream/telescope/backend/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type JobApi struct{}

var producer *lib.MyKafka

func init() {
	lib.RegisterProducer(producer)
}

func (j *JobApi) Upload(c *gin.Context) {
	user, err := utils.GetValueFromSessionByKey[db.AUser](c, "user")
	if err != nil {
		singleton.Logger.Error("fail to get user from session", zap.Error(err))
		return
	}
	uid := user.ID
	err = producer.Push("admin1", strconv.Itoa(int(uid)), "resizer")
	if err != nil {
		singleton.Logger.Error("Failed to push message to kafka", zap.String("error", err.Error()))
		response.FailWithMessage(c, "error.kafka.push")
		return
	}
	response.OkWithMessage(c, "success")
}
