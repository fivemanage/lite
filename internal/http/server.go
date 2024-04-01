package http

import "github.com/gin-gonic/gin"

type Server struct {
	Engine *gin.Engine
}

func NewServer() Server {
	server := gin.Default()

	apiGroup := server.Group("/api")
	apiGroup.GET("/auth/login", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	return Server{
		Engine: server,
	}
}
