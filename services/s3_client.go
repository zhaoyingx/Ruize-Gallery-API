package services

import (
	"fmt"
    "context"
    "os"

    "github.com/aws/aws-sdk-go-v2/config"
    "github.com/aws/aws-sdk-go-v2/service/s3"
)

func NewS3Client() (*s3.Client, string, error) {
    bucket := os.Getenv("AWS_S3_BUCKET")
    region := os.Getenv("AWS_REGION")
    
    if bucket == "" || region == "" {
        return nil, "", fmt.Errorf("missing AWS_S3_BUCKET or AWS_REGION env variables")
    }

    cfg, err := config.LoadDefaultConfig(context.TODO(),
        config.WithRegion(region),
    )
    if err != nil {
        return nil, "", err
    }

    client := s3.NewFromConfig(cfg)
    return client, bucket, nil
}