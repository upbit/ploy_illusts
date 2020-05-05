package main

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	log "github.com/upbit/ploy_illusts/logger"
	"github.com/upbit/ploy_illusts/model"
)

func requiredField(key string) {
	if value := viper.Get(key); value == nil {
		log.Errorf("Required field %s is not configured!", key)
		os.Exit(1)
	}
}

func initConfigs() {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./config")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		log.Errorf("Fatal error config file: %s", err)
		os.Exit(1)
	}

	requiredField("server.listen")
}

func initHTTPServer() *gin.Engine {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.Redirect(302, "/api/stat")
	})

	api := r.Group("/api")
	{
		api.GET("/stat", func(c *gin.Context) {
			c.JSON(200, gin.H{"status": "ok"})
		})

		// api.GET("/illusts", illusts.GetAllIllusts)
		// api.GET("/illusts/:id", illusts.GetIllust)
	}

	r.GET("/download", func(c *gin.Context) {
		url := c.Query("url")
		buffer, contentType, err := model.GetPixivImage(context.Background(), url)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
		} else {
			c.Data(200, contentType, buffer)
		}
	})

	// Assets
	r.Static("/data", "./data")
	r.StaticFile("/favicon.ico", "./assets/favicon.ico")
	r.StaticFile("/robots.txt", "./assets/robots.txt")

	// Error handling
	r.NoRoute(func(c *gin.Context) {
		c.AbortWithStatus(http.StatusNotFound)
	})

	return r
}

func main() {
	initConfigs()

	// Server configs
	addr := viper.GetString("server.listen")
	log.Infof("Start server at %s", addr)

	s := &http.Server{
		Addr:           addr,
		Handler:        initHTTPServer(),
		ReadTimeout:    2 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Errorf("%v", s.ListenAndServe())
}
