package main

import (
	"github.com/gin-gonic/gin"
	"github.com/itsvagapov/students"
)

func main() {
	r := gin.Default()

	r.GET("/students", students.GetAllStudentsHandler)
	r.POST("/students", students.AddStudentHandler)

	r.Run(":8080")
}