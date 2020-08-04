package app

import (
	"github.com/gin-gonic/gin"
	"github.com/nguyenbt456/todolist-go/controller"
)

// InitRouter init route for app
func InitRouter() *gin.Engine {
	router := gin.Default()

	auth := router.Group("/auth")
	{
		auth.POST("/register", controller.Register)
		auth.POST("/login", controller.Login)
		auth.POST("/logout")
	}

	router.GET("/", controller.GetTasks)
	router.POST("/", controller.CreateTask)
	router.PUT("/", controller.UpdateTask)
	router.DELETE("/", controller.DeleteTask)

	return router
}
