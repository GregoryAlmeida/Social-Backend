package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Initialize() {

	router := gin.Default()

	// Cors Config
	corsConfig := cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"POST", "DELETE", "PUT", "GET", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}
	router.Use(cors.New(corsConfig))

	//InitializeRoutes
	InitializeRoutes(router)

	//API run at 9090 port
	err := router.Run(":9090")
	if err != nil {
		panic(err)
	}
}
