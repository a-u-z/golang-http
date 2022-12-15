package main

import (
	"time"

	corss "github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
}

func (s *Server) routesV2() {

	s.router.Use(corss.New(corss.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "authorization", "Referer"},
		AllowCredentials: false,
		AllowAllOrigins:  true,
		MaxAge:           12 * time.Hour, // pre-flight request cache
	}))

	// middleware 與 next 的路徑走法
	// middleware 1 in => middleware 2 in => ping => middleware 2 out => middleware 1 out
	s.router.Use(middleware1) // 先加的 middleware 先進去
	s.router.Use(middleware2)
	s.router.GET("/", s.Ping)
}
