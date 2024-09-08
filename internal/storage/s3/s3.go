package s3

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type Storage struct {
	client *s3.Client
}

func New() *Storage {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		// TODO: handle error properly
		panic(err)
	}

	// s3 will automatically use the following environment variables:
	// AWS_ACCESS_KEY_ID
	// AWS_SECRET_ACCESS_KEY
	client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.Region = "eu-west-1"
		// mainly used for cloudflare r2
		o.BaseEndpoint = aws.String("https://s3.eu-west-1.amazonaws.com")
	})

	return &Storage{
		client: client,
	}
}

func (s *Storage) UploadFile() error {
	return nil
}

func (s *Storage) DeleteFile() error {
	return nil
}
