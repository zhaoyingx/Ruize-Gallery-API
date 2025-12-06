package services

import (
	"api-template/models"
	"context"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go/aws"
)

var GalleryService = &galleryService{}

type galleryService struct{
	client     *s3.Client
    bucketName string
}

func InitGalleryService() error {
    cli, bucket, err := NewS3Client()
    if err != nil {
        return err
    }

    GalleryService = &galleryService{
        client:     cli,
        bucketName: bucket,
    }

    return nil
}

func (s *galleryService) QueryByYearFromS3(year int) ([]models.Images, error) {
	prefix := fmt.Sprintf("%d/", year)

    out, err := s.client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
        Bucket: aws.String(s.bucketName),
        Prefix: aws.String(prefix),
    })
    if err != nil {
        return nil, fmt.Errorf("S3 列表查询失败: %w", err)
    }

    var results []models.Images

    for _, item := range out.Contents {
        key := *item.Key

        if strings.HasSuffix(key, "/") {
            continue
        }

        url := fmt.Sprintf("https://%s.s3.amazonaws.com/%s", s.bucketName, key)

        results = append(results, models.Images{
            Key: key,
            Url: url,
        })
    }

    return results, nil
}


