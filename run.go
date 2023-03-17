package apiSwagger

import (
	"github.com/TskFok/apiSwagger/config"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Run() {
	config.InitConfig()

	router := gin.Default()

	router.StaticFS("swagger", http.Dir("swagger"))
	router.LoadHTMLGlob("swagger/html/*")

	router.GET("/swagger", func(context *gin.Context) {
		context.HTML(http.StatusOK, "swagger.html", gin.H{
			"list": config.GetSwaggerMap(),
		})
	})

	router.Run(":8011")
}
