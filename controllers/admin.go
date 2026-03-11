package controllers

import (
	"backend/db"
	"backend/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PromoteUser(c *gin.Context) {

	// Authorization check
	role, exists := c.Get("role")

	if !exists || role != "ADMIN" {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "admin access required",
		})
		return
	}

	userID := c.Param("user_id")

	var user models.User

	// fetch user
	if err := db.DB.First(&user, userID).Error; err != nil {
		log.Printf("error occurred during fetching user: %v", err)

		c.JSON(http.StatusNotFound, gin.H{
			"error": "error occurred during fetching user",
		})
		return
	}

	// promote
	user.Role = "ADMIN"

	if err := db.DB.Save(&user).Error; err != nil {
		log.Printf("error occurred during promoting user: %v", err)

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "error occurred during promoting user",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "user promoted to admin",
	})
}

func GetUsers(c *gin.Context) {

	role, exists := c.Get("role")

	if !exists || role != "ADMIN" {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "admin access required",
		})
		return
	}

	var users []models.User

	if err := db.DB.Find(&users).Error; err != nil {
		log.Printf("error occurred during fetching users: %v", err)

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "error occurred during fetching users",
		})
		return
	}

	c.JSON(http.StatusOK, users)
}