package routes

import (
	"BE_Go/handlers"
	"BE_Go/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	public := r.Group("/api")
	{
		public.GET("/houses", handlers.GetHouses)
	}

	admin := r.Group("/api/admin")
	admin.Use(middleware.AuthRequired())
	{
		admin.POST("/points/add", handlers.AddPoints)
	}
}
