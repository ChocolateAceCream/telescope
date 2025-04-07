package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/ChocolateAceCream/telescope/backend/model/dbmodel"
	"github.com/ChocolateAceCream/telescope/backend/model/request"
	"github.com/ChocolateAceCream/telescope/backend/model/response"
	"github.com/ChocolateAceCream/telescope/backend/singleton"
	"github.com/ChocolateAceCream/telescope/backend/utils"
	"github.com/aws/aws-sdk-go-v2/feature/cloudfront/sign"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gin-gonic/gin"
)

type AwsService struct{}

func (a *AwsService) ClassifyImage(c *gin.Context, req request.ClassifyRequest) (resp response.ClassifyResponse, err error) {
	lambdaUrl := singleton.Config.AWS.Lambda.Url
	jsonData, err := json.Marshal(req)
	httpResp, err := http.Post(
		lambdaUrl,
		"application/json",
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		return
	}
	defer httpResp.Body.Close()

	err = json.NewDecoder(httpResp.Body).Decode(&resp)
	return
}

func (a *AwsService) GetS3UploadPresignedUrl(c *gin.Context, user dbmodel.UserInfo, fileName string) (resp response.GetS3UploadPresignedUrlResponse, err error) {
	path := user.Username + "/" + fileName
	r, err := singleton.AWS.PresignClient.PresignPutObject(c, &s3.PutObjectInput{
		Bucket: &singleton.Config.AWS.S3.Bucket,
		Key:    &path, // object name, s3 will store the file with this name
	}, func(opts *s3.PresignOptions) {
		opts.Expires = time.Duration(singleton.Config.AWS.S3.PresignedUrlExpiration) * time.Second
	})
	if err != nil {
		return
	}
	resp.PresignedUrl = r.URL
	resp.ImageUrl = strings.Split(r.URL, "?")[0]
	return
}

func (a *AwsService) GetS3DownloadPresignedUrl(c *gin.Context, user dbmodel.UserInfo, fileName string) (url string, err error) {
	path := user.Username + "/" + fileName
	req, err := singleton.AWS.PresignClient.PresignGetObject(c, &s3.GetObjectInput{
		Bucket: &singleton.Config.AWS.S3.Bucket,
		Key:    &path,
	}, func(opts *s3.PresignOptions) {
		opts.Expires = time.Duration(singleton.Config.AWS.S3.PresignedUrlExpiration) * time.Second
	})
	if err != nil {
		return
	}
	url = req.URL
	return
}

func (a *AwsService) GetCloudfrontSignedUrl(c *gin.Context, user dbmodel.UserInfo, fileName string) (signedURL string, err error) {
	privateKey, err := utils.LoadPrivateKey("certs/private_key.pem")
	if err != nil {
		fmt.Println("failed to load private key:", err.Error())
		return
	}

	signer := sign.NewURLSigner(singleton.Config.AWS.CloudFront.KeyID, privateKey)
	rawURL := singleton.Config.AWS.CloudFront.Prefix + user.Username + "/" + fileName
	fmt.Println("rawURL:", rawURL)
	expiration := time.Now().Add(time.Duration(singleton.Config.AWS.CloudFront.SignedUrlExpiration) * time.Second)
	signedURL, err = signer.Sign(rawURL, expiration)
	if err != nil {
		fmt.Println("failed to sign URL:", err.Error())
	}
	return
}
