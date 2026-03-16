package routes

import (
	"net/http"
	"studentapi/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	// Health check
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Student Management API Ready! 🚀",
		})
	})

	// Student routes
	api := r.Group("/api")
	{
		api.GET("/students",     controllers.GetAllStudents)
		api.GET("/students/:id", controllers.GetStudentByID)
		api.POST("/students",    controllers.CreateStudent)
		api.PUT("/students/:id", controllers.UpdateStudent)
		api.DELETE("/students/:id", controllers.DeleteStudent)
	}

	return r
}