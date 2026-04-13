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
