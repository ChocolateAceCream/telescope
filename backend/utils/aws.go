package utils

import (
	"context"

	"github.com/ChocolateAceCream/telescope/backend/singleton"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"go.uber.org/zap"
)

type Options func(*singleton.AWSClient)

func NewAWS(opts ...Options) *singleton.AWSClient {
	aws := &singleton.AWSClient{}
	for _, opt := range opts {
		opt(aws)
	}
	return aws
}

func WithS3(aws *singleton.AWSClient) {

	// TODO: move region to config
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-1"))
	if err != nil {
		singleton.Logger.Error("failed to load AWS config", zap.Error(err))
	}

	// create s3 service client
	s3Client := s3.NewFromConfig(cfg)
	aws.S3 = s3Client
	presignClient := s3.NewPresignClient(s3Client)
	aws.PresignClient = presignClient
}
