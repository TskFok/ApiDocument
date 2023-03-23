package router

import (
	"github.com/TskFok/ApiDocument/config"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter(Handle *gin.Engine) {
	Handle.GET("/swagger", func(context *gin.Context) {
		context.HTML(http.StatusOK, "swagger.html", gin.H{
			"list": config.GetSwaggerMap(),
		})
	})

}
