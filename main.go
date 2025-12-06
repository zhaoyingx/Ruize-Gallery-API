package main

import (
	"log"

	"api-template/config"
	"api-template/managers"
	"api-template/routes"
	"api-template/services"
)

func main() {
	// 1. 加载配置
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("加载配置失败: %v", err)
	}

	// 2. 初始化数据库
	db, err := managers.Init(cfg.DatabasePath)
	if err != nil {
		log.Fatalf("初始化数据库失败: %v", err)
	}

	// 3. 初始化 S3 client
	if err := services.InitGalleryService(); err != nil {
		log.Fatalf("初始化 S3 失败: %v", err)
	}

	// 4. 启动路由
	r := routes.SetupRouter(db)
	log.Printf("服务器启动于 :%s", cfg.Port)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatalf("服务器异常退出: %v", err)
	}
}
