package routers

import (
	"pdf-generator-service/internal/handler"

	"github.com/gin-gonic/gin"
)

func SetUpRouters() *gin.Engine {
	router := gin.Default()

	var Handlers handler.Contollers

	router.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"status": "OK"})
	})

	router.POST("/generate-pdf", Handlers.GeneratePDFHandler)

	return router
}
