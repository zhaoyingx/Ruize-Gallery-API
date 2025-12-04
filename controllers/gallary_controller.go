package controllers

import (
	"net/http"
	"strconv"

	"api-template/services"

	"github.com/gin-gonic/gin"
)

func QueryImagesByYear(ctx *gin.Context) {
	yearStr := ctx.Query("year")

	if yearStr == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "year 参数不能为空"})
		return
	}

	year, err := strconv.Atoi(yearStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "year 必须是数字"})
		return
	}

	results, err := services.GalleryService.QueryByYear(year)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": results,
	})
}