package handlers

import (
	"BE_Go/config"
	"BE_Go/models"

	"github.com/gin-gonic/gin"
)

func GetHouses(c *gin.Context) {
	var houses []models.House
	config.DB.Find(&houses)
	c.JSON(200, houses)
}

func GetLivePoints(c *gin.Context) {
	var houses []models.House
	config.DB.Order("house_points desc").Find(&houses)

	type HouseRank struct {
		Rank        int    `json:"rank"`
		Name        string `json:"name"`
		Points      int    `json:"points"`
		Description string `json:"description"`
		LogoURL     string `json:"logo_url"`
	}

	var result []HouseRank
	for i, h := range houses {
		result = append(result, HouseRank{
			Rank:        i + 1,
			Name:        h.Name,
			Points:      h.HousePoints,
			Description: h.Description,
			LogoURL:     h.LogoURL,
		})
	}
	c.JSON(200, result)
}

func GetAnnouncements(c *gin.Context) {
	var announcements []models.Announcement
	config.DB.Preload("House").Preload("Captain").
		Order("created_at desc").
		Find(&announcements)
	c.JSON(200, announcements)
}
