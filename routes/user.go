package routes

import (
	"go-layered/controllers"
	"github.com/gin-gonic/gin"
)

func NewUserRouter(ctrl controllers.UserController) *gin.Engine {
	router := gin.Default()
	
    // Apply CORS middleware globally with permissive or specific configurations
    router.Use(cors.Default()) 
	
	router.GET("/users", ctrl.GetUsers)
	router.GET("/user/:id", ctrl.GetUserById)
	router.POST("/user", ctrl.CreateUser)
	router.PUT("/user", ctrl.UpdateUser)
	router.PATCH("/user/:id", ctrl.PatchUserById)
	router.DELETE("user/:id", ctrl.DeleteUserById)
	
	router.HEAD("/ping", ctrl.Ping)
	
	return router
}
