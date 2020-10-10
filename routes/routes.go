package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/indexofnull/stuco2020/controllers"
)

//SetupRouter registers all routes to their corresponding controllers
func SetupRouter() *gin.Engine {
	r := gin.Default()
	grp1 := r.Group("/api")
	{
		grp1.GET("ping", controllers.PongHandler)
		grp1.GET("students", controllers.GetAllStudents)
		grp1.GET("deletestudent/:id", controllers.DeleteStudent)
	}
	classGroup := r.Group("/class")
	{
		classGroup.GET(":id/members", controllers.GetClassMembers)
		classGroup.GET(":id", controllers.GetClass)
	}
	studentGroup := r.Group("/student")
	{
		studentGroup.GET(":id", controllers.GetStudent)
	}
	codeGroup := r.Group("/code")
	{
		codeGroup.GET(":id", controllers.ResolveCode)
	}
	r.GET("/vote/:code", controllers.CastVote)
	return r
}
