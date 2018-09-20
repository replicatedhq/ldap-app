package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

var (
	logger = log.New(os.Stdout, "", log.Lshortfile)
)

func main() {
	h := &handlers{}
	go h.waitForDB()

	r := gin.Default()
	r.GET("/v1/ping", h.handlePing)
	r.POST("/v1/user/create", h.handleCreate)
	r.POST("/v1/user/modify", h.handleModify)
	r.DELETE("/v1/user/:uuid", h.handleDelete)
	r.Run("0.0.0.0:3000")
}
