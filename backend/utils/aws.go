package utils

import (
	"context"
	"crypto"
	"crypto/ecdsa"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"

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

func LoadPrivateKey(path string) (signer crypto.Signer, err error) {
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
