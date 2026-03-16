package controllers

import (
	"studentapi/config"
	"studentapi/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Sab students lao
func GetAllStudents(c *gin.Context) {
	var students []models.Student
	config.DB.Find(&students)
	c.JSON(http.StatusOK, gin.H{
		"total":    len(students),
		"students": students,
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