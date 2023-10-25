package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// The server.
type Server struct {
	config *Config
	router *gin.Engine
	api    *API
}

func (s *Server) addRoutes() {
	s.router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})
	s.router.POST("/", func(c *gin.Context) {
		value := strings.TrimSpace(c.PostForm("email_login"))
		domain := s.api.GetClientDomain(value)
		if domain != "" {
			c.Redirect(
				http.StatusFound,
				fmt.Sprintf(s.config.BaseURL, domain),
			)
		} else {
			c.HTML(http.StatusOK, "index.html", gin.H{
				"error": "Неверный логин или email",
				"value": value,
			})
		}
	})
}

// Run starts the server.
func (s *Server) Run() {
	s.router.LoadHTMLFiles("templates/index.html")
	s.addRoutes()
	if err := s.router.Run(fmt.Sprintf(":%d", s.config.Port)); err != nil {
		panic(err)
	}

}

// NewServer creates a new server.
func NewServer(config *Config, api *API) *Server {
	if config.IsProd {
		gin.SetMode(gin.ReleaseMode)
	}

	return &Server{
		config: config,
		router: gin.Default(),
		api:    api,
	}
}
