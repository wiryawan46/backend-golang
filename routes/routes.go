package routes

import (
	"backend-golang/controller"
	"backend-golang/middlewares"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	// initialize gin router
	router := gin.Default()

	//setup cors
	router.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:  []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders: []string{"Content-Length"},
	}))

	// define routes
	router.POST("/api/register", controller.Register)
	router.POST("/api/login", controller.Login)

	//users
	router.GET("/api/users", middlewares.AuthMiddleware(), controller.FindUser)
	router.POST("/api/users", middlewares.AuthMiddleware(), controller.CreateUser)
	router.GET("/api/users/:id", middlewares.AuthMiddleware(), controller.FindByUserId)
	router.PUT("/api/users/:id", middlewares.AuthMiddleware(), controller.UpdateUser)
	router.DELETE("/api/users/:id", middlewares.AuthMiddleware(), controller.DeleteUser)

	return router
}
