package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/indexofnull/stuco2020/controllers"
)

//CORS attaches a wildcard Access-Control-Allow-Origin because gin's official package doesn't work for some dumb reason.
//Yes, this is insecure
func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		//c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		//c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		//c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		/*if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}*/

		c.Next()
	}
}

//SetupRouter registers all routes to their corresponding controllers as well as some middleware
func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(CORS())
	/*grp1 := r.Group("/api")
	{
		grp1.GET("ping", controllers.PongHandler)
		grp1.GET("students", controllers.GetAllStudents)
		//grp1.GET("deletestudent/:id", controllers.DeleteStudent)
	}*/
	classGroup := r.Group("/class")
	{
		classGroup.GET(":id/members", controllers.GetClassMembers)
		classGroup.GET(":id", controllers.GetClass)
	}
	r.GET("/student/:id", controllers.GetStudent)
	r.GET("/code/:id", controllers.ResolveCode)
	//r.POST("/vote/:code", controllers.CastVote)
	return r
}
