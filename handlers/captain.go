package handlers

import (
	"BE_Go/config"
	"BE_Go/models"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func GetCaptainDashboard(c *gin.Context) {
	captainID, _ := c.Cookie("user_id")

	var captain models.Captain
	config.DB.First(&captain, captainID)

	var house models.House
	config.DB.First(&house, captain.HouseID)

	var members []models.Member
	config.DB.Where("house_id = ?", captain.HouseID).Find(&members)

	var announcements []models.Announcement
	config.DB.Where("captain_id = ?", captainID).
		Order("created_at desc").
		Find(&announcements)

	c.JSON(200, gin.H{
		"house":            house,
		"members":          members,
		"my_announcements": announcements,
	})
}

func CaptainCreateAnnouncement(c *gin.Context) {
	captainID, _ := c.Cookie("user_id")

	// switch to form instead of JSON since we have a file
	title := c.PostForm("title")
	content := c.PostForm("content")

	if title == "" || content == "" {
		c.JSON(400, gin.H{"error": "Title and content are required"})
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

		url, err := config.UploadImage(file, "announcements", fmt.Sprintf("ann_%s_%d", captainID, time.Now().Unix()))
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to upload image"})
			return
		}
		imageURL = &url
	}

	var captain models.Captain
	config.DB.First(&captain, captainID)

	var captainIDInt int
	fmt.Sscan(captainID, &captainIDInt)

	announcement := models.Announcement{
		Title:     title,
		Content:   content,
		ImageURL:  imageURL,
		HouseID:   captain.HouseID,
		CaptainID: &captainIDInt,
		CreatedAt: time.Now(),
	}
	config.DB.Create(&announcement)

	c.JSON(200, gin.H{
		"success":      true,
		"announcement": announcement,
	})
}

func CaptainDeleteAnnouncement(c *gin.Context) {
	captainID, _ := c.Cookie("user_id")
	announcementID := c.Param("id")

	var announcement models.Announcement
	if err := config.DB.First(&announcement, announcementID).Error; err != nil {
		c.JSON(404, gin.H{"error": "Announcement not found"})
		return
	}

	var captainIDInt int
	fmt.Sscan(captainID, &captainIDInt)

	if *announcement.CaptainID != captainIDInt {
		c.JSON(403, gin.H{"error": "You can only delete your own announcements"})
		return
	}

	config.DB.Delete(&announcement)
	c.JSON(200, gin.H{"success": true})
}
