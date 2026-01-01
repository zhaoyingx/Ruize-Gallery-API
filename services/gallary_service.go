package services

import (
	"api-template/models"
	"context"
	"fmt"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go/aws"
)

var GalleryService = &galleryService{}

type galleryService struct {
	client     *s3.Client
	bucketName string
}

func (s *galleryService) GetClient() *s3.Client {
	return s.client
}

func (s *galleryService) GetBucketName() string {
	return s.bucketName
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

	results := make([]models.Images, 0)

	presignClient := s3.NewPresignClient(s.client)

	for _, item := range out.Contents {
		key := *item.Key

		if strings.HasSuffix(key, "/") {
			continue
		}

		var imageUrl string
		lowerKey := strings.ToLower(key)

		if strings.HasSuffix(lowerKey, ".heic") || strings.HasSuffix(lowerKey, ".heif") {
			baseURL := os.Getenv("BASE_URL")
			if baseURL == "" {
				baseURL = "http://localhost:8082"
			}
			imageUrl = fmt.Sprintf("%s/api/v1/image/convert?key=%s", baseURL, url.QueryEscape(key))
		} else {
			presignResult, err := presignClient.PresignGetObject(context.TODO(), &s3.GetObjectInput{
				Bucket: aws.String(s.bucketName),
				Key:    aws.String(key),
			}, s3.WithPresignExpires(time.Hour*1))

			if err != nil {
				return nil, fmt.Errorf("生成预签名 URL 失败: %w", err)
			}
			imageUrl = presignResult.URL
		}

		results = append(results, models.Images{
			Key: key,
			Url: imageUrl,
		})
	}

	return results, nil
}
