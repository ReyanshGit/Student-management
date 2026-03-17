package controllers

import (
	"net/http"
	"strconv"
	"studentapi/config"
	"studentapi/models"

	"github.com/gin-gonic/gin"
)

func GetAllStudents(c *gin.Context) {

	page := c.DefaultQuery("page", "1")
	limit := c.DefaultQuery("limit", "10")
	search := c.DefaultQuery("search", "")
	course := c.DefaultQuery("course", "")
	sort := c.DefaultQuery("sort", "created_at")

	pageInt, _ := strconv.Atoi(page)
	limitInt, _ := strconv.Atoi(limit)
	offset := (pageInt - 1) * limitInt

	// Query banao
	query := config.DB.Model(&models.Student{})

	
	// Search — naam mein dhundho ← Yeh add karo
	if search != "" {
		query = query.Where("name LIKE ?", "%"+search+"%")
	}
	// Filter by course ← Line 2 add karo
	if course != "" {
		query = query.Where("course = ?", course)
	}
	// Total count
	var total int64
	query.Count(&total)

	// Students lo
	var students []models.Student
	query.Order(sort).Limit(limitInt).Offset(offset).Find(&students)

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
