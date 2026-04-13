package handlers

import (
	"BE_Go/config"
	"BE_Go/models"
	"fmt"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {
	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	// check admin first
	var admin models.Admin
	if err := config.DB.Where("username = ?", input.Username).First(&admin).Error; err == nil {
		if bcrypt.CompareHashAndPassword([]byte(admin.PasswordHash), []byte(input.Password)) == nil {
			// store in session cookie
			c.SetCookie("user_id", fmt.Sprintf("%d", admin.ID), 3600, "/", "", false, true)
			c.SetCookie("user_role", "admin", 3600, "/", "", false, true)
			c.JSON(200, gin.H{
				"success": true,
				"role":    "admin",
				"user": gin.H{
					"id":       admin.ID,
					"username": admin.Username,
					"name":     admin.Name,
				},
			})
			return
		}
	}

	// check captain
	var captain models.Captain
	if err := config.DB.Where("username = ?", input.Username).First(&captain).Error; err == nil {
		if bcrypt.CompareHashAndPassword([]byte(captain.PasswordHash), []byte(input.Password)) == nil {
			c.SetCookie("user_id", fmt.Sprintf("%d", captain.ID), 3600, "/", "", false, true)
			c.SetCookie("user_role", "captain", 3600, "/", "", false, true)
			c.JSON(200, gin.H{
				"success": true,
				"role":    "captain",
				"user": gin.H{
					"id":       captain.ID,
					"username": captain.Username,
					"name":     captain.Name,
				},
			})
			return
		}
	}

	c.JSON(401, gin.H{"error": "Invalid username or password"})
}

func Logout(c *gin.Context) {
	c.SetCookie("user_id", "", -1, "/", "", false, true)
	c.SetCookie("user_role", "", -1, "/", "", false, true)
	c.JSON(200, gin.H{"success": true})
}

func Me(c *gin.Context) {
	userID, _ := c.Cookie("user_id")
	userRole, _ := c.Cookie("user_role")

	if userRole == "admin" {
		var admin models.Admin
		config.DB.First(&admin, userID)
		c.JSON(200, gin.H{
			"id":       admin.ID,
			"username": admin.Username,
			"name":     admin.Name,
			"role":     "admin",
		})
		return
	}

	var captain models.Captain
	config.DB.First(&captain, userID)
	c.JSON(200, gin.H{
		"id":       captain.ID,
		"username": captain.Username,
		"name":     captain.Name,
		"role":     "captain",
		"house_id": captain.HouseID,
	})
}
