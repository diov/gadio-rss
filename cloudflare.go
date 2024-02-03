package main

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

const (
	bucketName = "assets-wtf"
	filePath   = "misc/gcores.xml"
)

var (
	r2Mgr *r2Manager
)

type r2Manager struct {
	client *s3.Client
}

func setupR2Manager(accountID, accessKeyID, accessKeySecret string) error {
	r2Resolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		return aws.Endpoint{
			URL: fmt.Sprintf("https://%s.r2.cloudflarestorage.com", accountID),
		}, nil
	})
	cfg, err := config.LoadDefaultConfig(context.Background(),
		config.WithEndpointResolverWithOptions(r2Resolver),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKeyID, accessKeySecret, "")),
		config.WithRegion("auto"),
	)
	if err != nil {
		return err
	}
	client := s3.NewFromConfig(cfg)
	r2Mgr = &r2Manager{client: client}
	return nil
}

func (m *r2Manager) uploadFeedFile(path string) error {
	if nil == m {
		return nil
	}

	file, err := os.Open(path)
	if nil != err {
		return err
	}

	_, err = m.client.PutObject(context.Background(), &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(filePath),
		Body:   file,
	})
	return err
}
