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

	r.GET("/avatars", controllers.GetAllAvatars)
	r.GET("/users/:id/avatar", controllers.GetAvatarById)

	r.PUT("/users/:id/avatar", controllers.UpdateAvatar)

	r.DELETE("/users/:id/avatar", controllers.DeleteAvatar)
}
