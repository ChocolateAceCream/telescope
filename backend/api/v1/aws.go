package apiV1

import (
	"github.com/ChocolateAceCream/telescope/backend/model/dbmodel"
	"github.com/ChocolateAceCream/telescope/backend/model/request"
	"github.com/ChocolateAceCream/telescope/backend/model/response"
	"github.com/ChocolateAceCream/telescope/backend/singleton"
	"github.com/ChocolateAceCream/telescope/backend/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AwsApi struct{}

func (a *AwsApi) Classify(c *gin.Context) {
	var req request.ClassifyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		singleton.Logger.Error("Failed to bind JSON", zap.Error(err))
		response.FailWithMessage(c, "error.missing.params")
		return
	}

	resp, err := awsService.ClassifyImage(c, req)
	if err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}
	response.OkWithFullDetails(c, resp, "success")
}

func (a *AwsApi) GeneratePresignedUrl(c *gin.Context) {
	var req request.S3PresignedUrlRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		singleton.Logger.Error("Failed to bind JSON", zap.Error(err))
		response.FailWithMessage(c, "error.missing.params")
		return
	}
	user, err := utils.GetValueFromSessionByKey[dbmodel.UserInfo](c, "user")
	if err != nil {
		singleton.Logger.Error("fail to get user from session", zap.Error(err))
		response.FailWithMessage(c, err.Error())
		return
	}
	resp, err := awsService.GetS3UploadPresignedUrl(c, user, req.FileName)
	if err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}
	response.OkWithFullDetails(c, resp, "success")
}

func (a *AwsApi) Download(c *gin.Context) {
	var req request.S3PresignedUrlRequest
	if err := c.ShouldBind(&req); err != nil {
		singleton.Logger.Error("Failed to bind JSON", zap.Error(err))
		response.FailWithMessage(c, "error.missing.params")
		return
	}
	user, err := utils.GetValueFromSessionByKey[dbmodel.UserInfo](c, "user")
	if err != nil {
		singleton.Logger.Error("fail to get user from session", zap.Error(err))
		response.FailWithMessage(c, err.Error())
		return
	}

	// approach #1: using s3 presigned url
	// url, err := awsService.GetS3DownloadPresignedUrl(c, user, req.FileName)
	// if err != nil {
	// 	response.FailWithMessage(c, err.Error())
	// 	return
	// }

	// approach #2: using cloudfront signed url
	url, err := awsService.GetCloudfrontSignedUrl(c, user, req.FileName)
	response.OkWithFullDetails(c, gin.H{"url": url}, "success")
}
