package main

import (
	_ "embed"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kelseyhightower/envconfig"
	"github.com/konstfish/hits/config"
	"github.com/konstfish/hits/handler"
	"github.com/konstfish/hits/storage"
)

//go:embed static/logo.svg
var iconSVG string

func main() {
	var cfg config.Config
	if err := envconfig.Process("", &cfg); err != nil {
		log.Fatalf("Failed to process config: %v", err)
	}

	if !cfg.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	redisStore, err := storage.NewRedisStore(cfg.RedisAddr, cfg.RedisPassword, cfg.RedisDB)
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
	log.Println("Connected to Redis successfully")

	badgeHandler := handler.NewBadgeHandler(redisStore)

	router := gin.Default()

	router.GET("/static/logo.svg", func(c *gin.Context) {
		c.Header("Content-Type", "image/svg+xml")
		c.String(http.StatusOK, iconSVG)
	})

	router.GET("/", handler.HandleIndex)

	router.GET("/api/count/incr/badge.svg", badgeHandler.HandleIncrBadge)
	router.GET("/api/count/show/badge.svg", badgeHandler.HandleShowBadge)

	log.Printf("Server starting on :%s", cfg.Port)
	if err := router.Run(":" + cfg.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
