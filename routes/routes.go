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
		admin.POST("/house/:id/logo", handlers.UpdateHouseLogo)
		admin.POST("/announcements", handlers.CreateAnnouncement)
		admin.DELETE("/announcements/:id", handlers.DeleteAnnouncement)
	}

	captain := r.Group("/api/captain")
	captain.Use(middleware.AuthRequired())
	{
		captain.POST("/announcements", handlers.CreateAnnouncement)
		captain.DELETE("/announcements/:id", handlers.DeleteAnnouncement)
	}
}
