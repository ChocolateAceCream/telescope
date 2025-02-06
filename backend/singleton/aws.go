package singleton

import "github.com/aws/aws-sdk-go-v2/service/s3"

type AWSClient struct {
	S3            *s3.Client
	PresignClient *s3.PresignClient
}

type AWSConfig struct {
	S3         S3Config         `mapstructure:"s3" json:"s3" yaml:"s3"`
	CloudFront CloudFrontConfig `mapstructure:"cloud-front" json:"cloud_front" yaml:"cloud-front"`
	Lambda     LambdaConfig     `mapstructure:"lambda" json:"lambda" yaml:"lambda"`
}

type LambdaConfig struct {
	Url string `mapstructure:"url" json:"url" yaml:"url"`
}
type S3Config struct {
	Bucket                 string `mapstructure:"bucket" json:"bucket" yaml:"bucket"`
	PresignedUrlExpiration int64  `mapstructure:"presigned-url-expiration" json:"presigned_url_expiration" yaml:"presigned-url-expiration"`
}

type CloudFrontConfig struct {
	Prefix              string `mapstructure:"resource-url-prefix" json:"resource_url_prefix" yaml:"resource-url-prefix"`
	KeyID               string `mapstructure:"key-pair-id" json:"key_pair_id" yaml:"key-pair-id"`
	SignedUrlExpiration int64  `mapstructure:"signed-url-expiration" json:"signed_url_expiration" yaml:"signed-url-expiration"`
}
