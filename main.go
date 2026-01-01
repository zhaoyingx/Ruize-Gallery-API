package main

import (
	"log"

	"api-template/config"
	"api-template/routes"
	"api-template/services"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("加载配置失败: %v", err)
	}

	if err := services.InitGalleryService(); err != nil {
		log.Fatalf("初始化 S3 失败: %v", err)
	}

	r := routes.SetupRouter()
	log.Printf("服务器启动于 :%s", cfg.Port)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatalf("服务器异常退出: %v", err)
	}
}
