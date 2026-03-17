package controllers

import (
	"net/http"
	"strconv"
	"studentapi/config"
	"studentapi/models"

	"github.com/gin-gonic/gin"
)

// Sab students lao
func GetAllStudents(c *gin.Context) {

	// Page aur limit URL se lo
	page := c.DefaultQuery("page", "1")
	limit := c.DefaultQuery("limit", "10")

	// String to number
	pageInt, _ := strconv.Atoi(page)
	limitInt, _ := strconv.Atoi(limit)

	// Offset calculate karo
	offset := (pageInt - 1) * limitInt

	// Total count lo
	var total int64
	config.DB.Model(&models.Student{}).Count(&total)

	// Students lo
	var students []models.Student
	config.DB.Limit(limitInt).Offset(offset).Find(&students)

	// Total pages
	totalPages := (int(total) + limitInt - 1) / limitInt

	c.JSON(http.StatusOK, gin.H{
		"students":    students,
		"total":       total,
		"page":        pageInt,
		"limit":       limitInt,
		"total_pages": totalPages,
	})
}

// Ek student lao ID se
func GetStudentByID(c *gin.Context) {
	id := c.Param("id")
	var student models.Student
	if err := config.DB.First(&student, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Student nahi mila",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"student": student})
}

// Naya student add karo
func CreateStudent(c *gin.Context) {
	var input models.StudentInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Sab fields bharo",
		})
		return
	}

	student := models.Student{
		Name:   input.Name,
		Email:  input.Email,
		Phone:  input.Phone,
		Course: input.Course,
		Age:    input.Age,
	}

	if err := config.DB.Create(&student).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Email already exist karta hai",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "✅ Student add ho gaya!",
		"student": student,
	})
}

// Student update karo
func UpdateStudent(c *gin.Context) {
	id := c.Param("id")

	var student models.Student
	if err := config.DB.First(&student, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Student nahi mila",
		})
		return
	}

	var input models.StudentInput
	c.ShouldBindJSON(&input)

	config.DB.Model(&student).Updates(map[string]interface{}{
		"name":   input.Name,
		"email":  input.Email,
		"phone":  input.Phone,
		"course": input.Course,
		"age":    input.Age,
	})

	c.JSON(http.StatusOK, gin.H{
		"message": "✅ Student update ho gaya!",
		"student": student,
	})
}

// Student delete karo
func DeleteStudent(c *gin.Context) {
	id := c.Param("id")

	var student models.Student
	if err := config.DB.First(&student, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Student nahi mila",
		})
		return
	}

	config.DB.Delete(&student)
	c.JSON(http.StatusOK, gin.H{
		"message": "✅ Student delete ho gaya!",
	})
}
