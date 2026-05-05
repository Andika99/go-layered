package routes

import (
	"go-layered/controllers"
	"github.com/gin-gonic/gin"
)

func NewUserRouter(ctrl controllers.UserController) *gin.Engine {
	router := gin.Default()
	
	router.GET("/users", ctrl.GetUsers)
	router.GET("/user/:id", ctrl.GetUserById)
	router.POST("/user", ctrl.CreateUser)
	router.PUT("/user", ctrl.UpdateUser)
	router.DELETE("user/:id", ctrl.DeleteUserById)
	
	return router
}
