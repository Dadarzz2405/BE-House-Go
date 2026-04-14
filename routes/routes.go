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
		public.GET("/houses/live", handlers.GetLivePoints)
		public.GET("/live-points", handlers.GetLivePoints)
		public.GET("/announcements", handlers.GetAnnouncements)
	}

	r.POST("/api/login", handlers.Login)
	r.POST("/api/logout", handlers.Logout)
	r.GET("/api/me", handlers.Me)

	admin := r.Group("/api/admin")
	admin.Use(middleware.AuthRequired())
	{
		admin.GET("/dashboard", handlers.GetAdminDashboard)
		admin.POST("/points/add", handlers.AddPoints)
		admin.POST("/points/deduct", handlers.DeductPoints)
		admin.POST("/house/:id/logo", handlers.UpdateHouseLogo)
		admin.POST("/announcements", handlers.AdminCreateAnnouncement)
		admin.DELETE("/announcements/:id", handlers.AdminDeleteAnnouncement)
		admin.GET("/announcements", handlers.GetAllAnnouncements)
	}

	captain := r.Group("/api/captain")
	captain.Use(middleware.AuthRequired())
	{
		captain.GET("/dashboard", handlers.GetCaptainDashboard)
		captain.POST("/announcements", handlers.CaptainCreateAnnouncement)
		captain.POST("/announcements/create", handlers.CaptainCreateAnnouncement)
		captain.DELETE("/announcements/:id", handlers.CaptainDeleteAnnouncement)
	}
}
