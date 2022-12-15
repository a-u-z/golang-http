package main

import (
	"github.com/gin-gonic/gin"
)

func NewServer() *Server {
	r := gin.New()

	r.Use(gin.Recovery())
	s := &Server{
		router: r,
	}
	s.routesV2()
	return s
}
