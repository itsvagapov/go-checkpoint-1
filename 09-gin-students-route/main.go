package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/itsvagapov/students"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "server is running")
	})

	r.GET("/students", func(c *gin.Context) {
		c.JSON(http.StatusOK, students.Students)
	})

	r.Run()
}