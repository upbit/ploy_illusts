package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/golang/glog"
	"github.com/golang/groupcache"
	"github.com/spf13/viper"
)

func getImage(ctx groupcache.Context, url string, dest groupcache.Sink) error {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Referer", "https://app-api.pixiv.net/")
	response, err := client.Do(req)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}
	dest.SetBytes(body)
	return nil
}

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
	// Server configs
	addr := viper.GetString("server.listen")
	peers_addrs := strings.Split(viper.GetString("server.peers"), ",")
	log.Infof("Start server at %s", addr)

	// groupcache
	peers := groupcache.NewHTTPPool("http://" + addr)
	peers.Set(peers_addrs...)
	cache := groupcache.NewGroup("image", 8<<30, groupcache.GetterFunc(getImage))

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

	r.GET("/stat", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	r.GET("/image", func(c *gin.Context) {
		url := c.Query("url")
		var data []byte
		cache.Get(nil, url, groupcache.AllocatingByteSliceSink(&data))
		fmt.Printf("Get %s return %d bytes\n", url, len(data))

		c.Writer.Header().Set("Content-Type", "image/jpeg")
		w := gin.ResponseWriter(c.Writer)
		w.Write(data)
	})

	// Error handling
	r.NoRoute(func(c *gin.Context) {
		c.HTML(404, "not_found.tmpl", gin.H{})
	})

	s := &http.Server{
		Addr:           addr,
		Handler:        r,
		ReadTimeout:    2 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
