package config

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func GetCORSConfig() gin.HandlerFunc {
	corsConfig := cors.Config{
		AllowOrigins:     []string{"http://18.204.74.129"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}
	return cors.New(corsConfig)
}
