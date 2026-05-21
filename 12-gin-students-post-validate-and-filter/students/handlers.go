package students

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllStudentsHandler(ctx *gin.Context) {
	filter := Filter{}

	// city
	city := ctx.Query("city")
	if city != "" {
		filter.City = &city
	}

	// min_age
	minAgeStr := ctx.Query("min_age")
	if minAgeStr != "" {

		minAge, err := strconv.Atoi(minAgeStr)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "некорректный параметр min_age",
			})
			return
		}

		filter.MinAge = &minAge
	}

	// max_age
	maxAgeStr := ctx.Query("max_age")
	if maxAgeStr != "" {

		maxAge, err := strconv.Atoi(maxAgeStr)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "некорректный параметр max_age",
			})
			return
		}

		filter.MaxAge = &maxAge
	}

	// sort
	sortValue := ctx.Query("sort")
	if sortValue != "" {

		if sortValue != "full_name" && sortValue != "age" {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "некорректный параметр sort",
			})
			return
		}

		filter.Sort = &sortValue
	}

	// limit
	limitStr := ctx.Query("limit")
	if limitStr != "" {

		limit, err := strconv.Atoi(limitStr)
		if err != nil || limit < 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "некорректный параметр limit",
			})
			return
		}

		filter.Limit = &limit
	}

	// offset
	offsetStr := ctx.Query("offset")
	if offsetStr != "" {

		offset, err := strconv.Atoi(offsetStr)
		if err != nil || offset < 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "некорректный параметр offset",
			})
			return
		}

		filter.Offset = &offset
	}

	ctx.JSON(http.StatusOK, GetAllStudents(filter))
}

func AddStudentHandler(ctx *gin.Context) {

	var student Student

	err := ctx.ShouldBindJSON(&student)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "некорректный JSON",
		})
		return
	}

	err = ValidateStudent(student)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	AddStudent(student)

	ctx.JSON(http.StatusCreated, student)
}