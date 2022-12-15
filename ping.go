package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func (s *Server) Ping(c *gin.Context) {
	log.Printf("here is ping")
}
