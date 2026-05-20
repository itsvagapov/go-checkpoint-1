package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/itsvagapov/students"
)

func main() {
	r := gin.Default()

	r.GET("/students", func(c *gin.Context) {
		c.JSON(http.StatusOK, students.Students)
	})

	r.GET("/students/:index", func(c *gin.Context) {

		index := c.Param("index")

		i, err := strconv.Atoi(index)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "индекс должен быть числом",
			})
			return
		}

		if i < 0 || i >= len(students.Students) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "студент не найден",
			})
			return
		}

		c.JSON(http.StatusOK, students.Students[i])
	})

	r.POST("/students", func(c *gin.Context) {
		c.String(http.StatusOK, "Студент добавлен")
	})

	r.Run()
}