package controllers

import (
	"crud-go/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetUsers(c *gin.Context, db *gorm.DB) {
	var user []models.User
	db.Find(&user)
	c.JSON(http.StatusOK, user)
}

func CreateUser(c *gin.Context, db *gorm.DB) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Create(&user)
	c.JSON(http.StatusCreated, user)
}
