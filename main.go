package main

import (
	"context"
	"embed"
	"fmt"
	"github.com/TskFok/ApiDocument/app/process"
	"github.com/TskFok/ApiDocument/config"
	"github.com/TskFok/ApiDocument/router"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"html/template"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

//go:embed public/html/*
var HtmlFs embed.FS

//go:embed public/*
var SwaggerFs embed.FS

func main() {
	process.InitProcess()

	config.InitConfig()

	gin.SetMode(gin.ReleaseMode)
	Handle := gin.New()
	Handle.Use(gin.Recovery())
	Handle.Use(gin.Logger())

	//html模版资源
	htmlTemplate := template.Must(template.New("").ParseFS(HtmlFs, "public/html/*"))
	Handle.SetHTMLTemplate(htmlTemplate)

	//静态资源
	Handle.Any("public/*filepath", func(context *gin.Context) {
		staticServer := http.FileServer(http.FS(SwaggerFs))
		staticServer.ServeHTTP(context.Writer, context.Request)
	})

	router.InitRouter(Handle)

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", 8011),
		Handler:        Handle,
		ReadTimeout:    time.Duration(20) * time.Second,
		WriteTimeout:   time.Duration(20) * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := s.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			log.Printf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}
