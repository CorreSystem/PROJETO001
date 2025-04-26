package routes

import (
	"micro-erp/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	api := r.Group("/api")
	{
		api.POST("/register", controllers.Register)
		api.POST("/login", controllers.Login)
	}
}
