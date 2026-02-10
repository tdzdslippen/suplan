package main

import (
	config "backend/internal/app"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"net/http"
)

func main() {
	router := gin.Default()
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("config error: %v", err)
	}
	api := router.Group("/api")
	api.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	api.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	addr := ":" + cfg.Port
	log.Printf("starting server on %s", addr)
	if err := router.Run(addr); err != nil {
		log.Fatalf("server error: %v", err)
	}

}
