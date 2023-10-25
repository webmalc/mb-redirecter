package main

import (
	"fmt"
	"net/http"
	"webmalc/mb-redirector/common/config"

	"github.com/gin-gonic/gin"
)

func main() {
	config.Setup()
	conf := NewConfig()
	router := gin.Default()
	router.LoadHTMLFiles("templates/index.html")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "MB redirector",
		})
	})

	if err := router.Run(fmt.Sprintf(":%d", conf.Port)); err != nil {
		panic(err)
	}
}
