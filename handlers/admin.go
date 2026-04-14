package handlers

import (
	"BE_Go/config"
	"BE_Go/models"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func GetAdminDashboard(c *gin.Context) {
	var houses []models.House
	config.DB.Order("house_points desc").Find(&houses)

	var transactions []models.PointTransaction
	config.DB.Preload("House").Preload("Admin").
		Order("timestamp desc").
		Limit(10).
		Find(&transactions)

	c.JSON(200, gin.H{
		"houses":              houses,
		"recent_transactions": transactions,
	})
}

func AddPoints(c *gin.Context) {
	var input struct {
		HouseID int    `json:"house_id"`
		Points  int    `json:"points"`
		Reason  string `json:"reason"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	if input.Points <= 0 {
		c.JSON(400, gin.H{"error": "Points must be positive"})
		return
	}

	var house models.House
	if err := config.DB.First(&house, input.HouseID).Error; err != nil {
		c.JSON(404, gin.H{"error": "House not found"})
		return
	}

	// get admin from cookie
	adminID, _ := c.Cookie("user_id")

	house.HousePoints += input.Points
	config.DB.Save(&house)

	var adminIDInt int
	fmt.Sscan(adminID, &adminIDInt)

	transaction := models.PointTransaction{
		HouseID:      input.HouseID,
		PointsChange: input.Points,
		Reason:       input.Reason,
		Timestamp:    time.Now(),
		AdminID:      &adminIDInt,
	}
	config.DB.Create(&transaction)

	c.JSON(200, gin.H{
		"success": true,
		"message": fmt.Sprintf("Added %d points to %s", input.Points, house.Name),
	})
}

func DeductPoints(c *gin.Context) {
	var input struct {
		HouseID int    `json:"house_id"`
		Points  int    `json:"points"`
		Reason  string `json:"reason"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	if input.Points <= 0 {
		c.JSON(400, gin.H{"error": "Points must be positive"})
		return
	}

	var house models.House
	if err := config.DB.First(&house, input.HouseID).Error; err != nil {
		c.JSON(404, gin.H{"error": "House not found"})
		return
	}

	// get admin from cookie
	adminID, _ := c.Cookie("user_id")

	house.HousePoints -= input.Points
	config.DB.Save(&house)

	var adminIDInt int
	fmt.Sscan(adminID, &adminIDInt)

	transaction := models.PointTransaction{
		HouseID:      input.HouseID,
		PointsChange: -input.Points,
		Reason:       input.Reason,
		Timestamp:    time.Now(),
		AdminID:      &adminIDInt,
	}
	config.DB.Create(&transaction)

	c.JSON(200, gin.H{
		"success": true,
		"message": fmt.Sprintf("Deducted %d points from %s", input.Points, house.Name),
	})
}

// UpdateHouseLogo handles updating a house's logo image
func UpdateHouseLogo(c *gin.Context) {
	houseID := c.Param("id")

	file, header, err := c.Request.FormFile("logo")
	if err != nil {
		c.JSON(400, gin.H{"error": "No file provided"})
		return
	}
	defer file.Close()

	// validate size (5MB)
	if header.Size > 5*1024*1024 {
		c.JSON(400, gin.H{"error": "File too large, max 5MB"})
		return
	}

	var house models.House
	if err := config.DB.First(&house, houseID).Error; err != nil {
		c.JSON(404, gin.H{"error": "House not found"})
		return
	}

	// upload to cloudinary
	url, err := config.UploadImage(file, "houses", house.Name)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to upload image"})
		return
	}

	house.LogoURL = url
	config.DB.Save(&house)

	c.JSON(200, gin.H{
		"success": true,
		"url":     url,
	})
}

// CreateAnnouncement allows admins to create announcements for any house
func AdminCreateAnnouncement(c *gin.Context) {
	// Admins can specify which house the announcement is for via form input
	var input struct {
		HouseID string `form:"house_id"`
		Title   string `form:"title"`
		Content string `form:"content"`
	}

	if err := c.ShouldBind(&input); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	if input.Title == "" || input.Content == "" || input.HouseID == "" {
		c.JSON(400, gin.H{"error": "House ID, title, and content are required"})
		return
	}

	// handle optional image
	var imageURL *string
	file, header, err := c.Request.FormFile("image")
	if err == nil {
		defer file.Close()

		if header.Size > 5*1024*1024 {
			c.JSON(400, gin.H{"error": "File too large, max 5MB"})
			return
		}

		url, err := config.UploadImage(file, "announcements", fmt.Sprintf("ann_admin_%s_%d", input.HouseID, time.Now().Unix()))
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to upload image"})
			return
		}
		imageURL = &url
	}

	var houseIDInt int
	fmt.Sscan(input.HouseID, &houseIDInt)

	// Verify house exists
	var house models.House
	if err := config.DB.First(&house, houseIDInt).Error; err != nil {
		c.JSON(404, gin.H{"error": "House not found"})
		return
	}

	announcement := models.Announcement{
		Title:     input.Title,
		Content:   input.Content,
		ImageURL:  imageURL,
		HouseID:   houseIDInt,
		CreatedAt: time.Now(),
		// AdminID can be added to the model if needed for tracking
	}

	config.DB.Create(&announcement)

	c.JSON(200, gin.H{
		"success":      true,
		"announcement": announcement,
	})
}

// DeleteAnnouncement allows admins to delete any announcement
func AdminDeleteAnnouncement(c *gin.Context) {
	announcementID := c.Param("id")

	var announcement models.Announcement
	if err := config.DB.First(&announcement, announcementID).Error; err != nil {
		c.JSON(404, gin.H{"error": "Announcement not found"})
		return
	}

	config.DB.Delete(&announcement)
	c.JSON(200, gin.H{"success": true})
}

// GetAllAnnouncements allows admins to view all announcements (optional feature)
func GetAllAnnouncements(c *gin.Context) {
	var announcements []models.Announcement
	config.DB.Preload("House").Order("created_at desc").Find(&announcements)

	c.JSON(200, gin.H{
		"announcements": announcements,
	})
}
