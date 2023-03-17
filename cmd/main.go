package main

import (
	"github.com/TskFok/apiSwagger/config"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	config.InitConfig()

	router := gin.Default()

	router.LoadHTMLFiles("swagger/html/swagger.html")
	router.Static("swagger", "swagger")

	router.GET("/swagger", func(context *gin.Context) {
		context.HTML(http.StatusOK, "swagger.html", gin.H{
			"list": config.GetSwaggerMap(),
		})
	})

	router.Run(":8011")
}
