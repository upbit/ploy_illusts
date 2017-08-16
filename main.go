package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/golang/glog"
	"github.com/spf13/viper"
)

func requiredField(key string) {
	if value := viper.Get(key); value == nil {
		log.Fatalf("Required field %s is not configured!", key)
	}
}

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath("./config")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		log.Fatalf("Fatal error config file: %s", err)
	}

	requiredField("server.listen")
}

func main() {
	r := gin.Default()

	// Assets
	r.Static("/static", "./assets/static")
	r.StaticFile("/favicon.ico", "./assets/favicon.ico")
	r.StaticFile("/robots.txt", "./assets/robots.txt")
	// Templates
	r.LoadHTMLGlob("templates/*")

	// Routers
	r.GET("/", func(c *gin.Context) {
		c.Redirect(301, "/stat")
	})

	// r.GET("/index", func(c *gin.Context) {
	//     c.HTML(200, "index.tmpl", gin.H{
	//         "ECharts入门示例 - 柱状图": "bar.html",
	//     })
	// })

	r.GET("/stat", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	// Error handling
	r.NoRoute(func(c *gin.Context) {
		c.HTML(404, "not_found.tmpl", gin.H{})
	})

	addr := viper.GetString("server.listen")
	log.Infof("Start server at %s", addr)

	s := &http.Server{
		Addr:           addr,
		Handler:        r,
		ReadTimeout:    2 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
