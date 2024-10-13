package Routes

import (
	"github.com/gin-gonic/gin"
	"github.com/valikak/go-practice/student/Controllers"
)

// SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	r := gin.Default()
	grp1 := r.Group("/student-app")
	{
		grp1.GET("student", Controllers.GetStudents)
		grp1.POST("student", Controllers.CreateStudent)
		grp1.GET("student/:id", Controllers.GetStudentByID)
		grp1.PUT("student/:id", Controllers.UpdateStudent)
		grp1.DELETE("student/:id", Controllers.DeleteStudent)
	}
	return r
}
