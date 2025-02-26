package service

import (
	"bytes"
	"crypto"
	"crypto/ecdsa"
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/ChocolateAceCream/telescope/backend/model/dbmodel"
	"github.com/ChocolateAceCream/telescope/backend/model/request"
	"github.com/ChocolateAceCream/telescope/backend/model/response"
	"github.com/ChocolateAceCream/telescope/backend/singleton"
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

func (a *AwsService) GetS3UploadPresignedUrl(c *gin.Context, user dbmodel.UserInfo, fileName string) (url string, err error) {
	path := user.Username + "/" + fileName
	req, err := singleton.AWS.PresignClient.PresignPutObject(c, &s3.PutObjectInput{
		Bucket: &singleton.Config.AWS.S3.Bucket,
		Key:    &path, // object name, s3 will store the file with this name
	}, func(opts *s3.PresignOptions) {
		opts.Expires = time.Duration(singleton.Config.AWS.S3.PresignedUrlExpiration) * time.Second
	})
	if err != nil {
		return
	}
	url = req.URL
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
	privateKey, err := loadPrivateKey("certs/private_key.pem")
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

func loadPrivateKey(path string) (signer crypto.Signer, err error) {
	raw, err := os.ReadFile(path)
	if err != nil {
		return
	}

	// Decode the PEM block
	block, _ := pem.Decode(raw)
	if block == nil {
		return nil, fmt.Errorf("failed to decode PEM block from file: %s", path)
	}

	// Parse the private key based on the key type (RSA, ECDSA, etc.)
	switch block.Type {
	case "RSA PRIVATE KEY":
		privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
		if err != nil {
			return nil, fmt.Errorf("failed to parse RSA private key: %v", err)
		}
		return privateKey, nil

	case "EC PRIVATE KEY":
		privateKey, err := x509.ParseECPrivateKey(block.Bytes)
		if err != nil {
			return nil, fmt.Errorf("failed to parse EC private key: %v", err)
		}
		return privateKey, nil

	case "PRIVATE KEY": // This may handle PKCS#8 encoded keys
		privateKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
		if err != nil {
			return nil, fmt.Errorf("failed to parse PKCS#8 private key: %v", err)
		}

		// Check if it's of the correct type (RSA, ECDSA)
		switch key := privateKey.(type) {
		case *rsa.PrivateKey:
			return key, nil
		case *ecdsa.PrivateKey:
			return key, nil
		default:
			return nil, fmt.Errorf("unsupported private key type: %T", privateKey)
		}

	default:
		return nil, fmt.Errorf("unsupported private key type: %s", block.Type)
	}
}
