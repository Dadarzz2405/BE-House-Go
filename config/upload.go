package config

import (
	"BE_Go/models"
	"context"
	"mime/multipart"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gin-gonic/gin"
)

// UploadImage uploads a file to Cloudinary and returns the URL
func UploadImage(file multipart.File, folder, publicID string) (string, error) {
	// Configure upload parameters
	var overwriteBool = true
	uploadParams := uploader.UploadParams{
		Folder:    folder,
		PublicID:  publicID,
		Overwrite: &overwriteBool,
	}

	// Upload file to Cloudinary
	uploadResult, err := CLD.Upload.Upload(context.Background(), file, uploadParams)
	if err != nil {
		return "", err
	}

	return uploadResult.SecureURL, nil
}

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
	if err := DB.First(&house, houseID).Error; err != nil {
		c.JSON(404, gin.H{"error": "House not found"})
		return
	}

	// upload to cloudinary
	url, err := UploadImage(file, "houses", house.Name)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to upload image"})
		return
	}

	house.LogoURL = url
	DB.Save(&house)

	c.JSON(200, gin.H{
		"success": true,
		"url":     url,
	})
}
