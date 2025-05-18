package routers

import (
	"crud-go/controllers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	r.GET("/staff-members/:id", func(c *gin.Context) {
		controllers.GetStaffMemberByID(c, db)
	})
	r.GET("/staff-members", func(c *gin.Context) {
		controllers.GetStaffMembers(c, db)
	})
	r.POST("/staff-members", func(c *gin.Context) {
		controllers.CreateStaffMember(c, db)
	})

	r.GET("/students/:id", func(c *gin.Context) {
		controllers.GetStudentByID(c, db)
	})
	r.GET("/students", func(c *gin.Context) {
		controllers.GetStudents(c, db)
	})
	r.POST("/students", func(c *gin.Context) {
		controllers.CreateStudent(c, db)
	})

	r.GET("/users", func(c *gin.Context) {
		controllers.GetUsers(c, db)
	})
	r.POST("/users", func(c *gin.Context) {
		controllers.CreateUser(c, db)
	})

	return r
}
