package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func middleware1(c *gin.Context) {
	log.Printf("here is middleware 1 in")
	c.Next()
	log.Printf("here is middleware 1 out")
}

func middleware2(c *gin.Context) {
	log.Printf("here is middleware 2 in")
	c.Next()
	log.Printf("here is middleware 2 out")
}
