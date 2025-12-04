package config

import (
	"os"

	"github.com/joho/godotenv"
)

// Config 应用配置
type Config struct {
	Port         string
	DatabasePath string
}

// Load 从 .env 加载配置
func Load() (*Config, error) {
	_ = godotenv.Load() // 忽略 .env 不存在的情况

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	dbPath := os.Getenv("DATABASE_PATH")
	if dbPath == "" {
		dbPath = "database/GraduateEmploymentSurvey.db"
	}

	return &Config{
		Port:         port,
		DatabasePath: dbPath,
	}, nil
}
