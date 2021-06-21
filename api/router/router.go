package router

import (
	"github.com/blyndusk/image-resizer/api/controllers"
	"github.com/gin-gonic/gin"
)

func Setup(r *gin.Engine) {
	usersRoute(r)
	avatarsRoute(r)
	r.GET("/load_fixtures", controllers.LoadData)

}

func usersRoute(r *gin.Engine) {
	r.POST("/users", controllers.CreateUser)

	r.GET("/users", controllers.GetAllUsers)
	r.GET("/users/:id", controllers.GetUserById)

	r.PUT("/users/:id", controllers.UpdateUser)

	r.DELETE("/users/:id", controllers.DeleteUser)
}


func avatarsRoute(r *gin.Engine) {
	r.POST("/avatars", controllers.CreateUser)

	r.GET("/avatars", controllers.GetAllUsers)
	r.GET("/avatars/:id", controllers.GetUserById)

	r.PUT("/avatars/:id", controllers.UpdateUser)

	r.DELETE("/avatars/:id", controllers.DeleteUser)
}
