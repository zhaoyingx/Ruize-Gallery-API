package controllers

import (
	"bytes"
	"context"
	"image/jpeg"
	"net/http"
	"strings"

	"api-template/services"

	"github.com/adrium/goheif"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/gin-gonic/gin"
)

// ConvertImage 转换 HEIC 图片为 JPEG
func ConvertImage(ctx *gin.Context) {
	key := ctx.Query("key")

	if key == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "key 参数不能为空"})
		return
	}

	// 从 S3 下载图片
	result, err := services.GalleryService.GetClient().GetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: aws.String(services.GalleryService.GetBucketName()),
		Key:    aws.String(key),
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "下载图片失败"})
		return
	}
	defer result.Body.Close()

	// 检查是否是 HEIC 格式
	lowerKey := strings.ToLower(key)
	if strings.HasSuffix(lowerKey, ".heic") || strings.HasSuffix(lowerKey, ".heif") {
		// 读取 HEIC 文件
		buf := new(bytes.Buffer)
		_, err := buf.ReadFrom(result.Body)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "读取图片失败"})
			return
		}

		// 解码 HEIC
		img, err := goheif.Decode(buf)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "解析 HEIC 文件失败: " + err.Error()})
			return
		}

		// 编码为 JPEG
		outputBuf := new(bytes.Buffer)
		err = jpeg.Encode(outputBuf, img, &jpeg.Options{Quality: 90})
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "编码 JPEG 失败"})
			return
		}

		// 返回 JPEG 图片
		ctx.Header("Content-Type", "image/jpeg")
		ctx.Header("Cache-Control", "public, max-age=86400") // 缓存 24 小时
		ctx.Data(http.StatusOK, "image/jpeg", outputBuf.Bytes())
		return
	}

	// 非 HEIC 格式，直接返回原图
	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(result.Body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "读取图片失败"})
		return
	}

	// 根据文件扩展名设置 Content-Type
	contentType := "image/jpeg"
	if strings.HasSuffix(lowerKey, ".png") {
		contentType = "image/png"
	} else if strings.HasSuffix(lowerKey, ".gif") {
		contentType = "image/gif"
	} else if strings.HasSuffix(lowerKey, ".webp") {
		contentType = "image/webp"
	}

	ctx.Header("Content-Type", contentType)
	ctx.Header("Cache-Control", "public, max-age=86400")
	ctx.Data(http.StatusOK, contentType, buf.Bytes())
}
