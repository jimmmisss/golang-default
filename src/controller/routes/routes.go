package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jimmmisss/golang-default/src/controller"
)

func InitRoutes(r *gin.RouterGroup) {
	r.GET("getUserById/:userId", controller.FindUserById)
	r.GET("getUserByEmail/:userEmail", controller.FindUserByEmail)
	r.POST("createUser/", controller.CreateUser)
	r.PUT("updateUser/:idUser", controller.UpdateUser)
	r.DELETE("deleteUser/:idUser", controller.DeleteUser)
}
