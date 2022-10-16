package controller

import (
	"net/http"

	"github.com/PattanasakGit/project-sa/entity"
	"github.com/gin-gonic/gin"
)

// POST /videos
func CreateBlood_type(c *gin.Context) {
	var Blood_type entity.Blood_type
	if err := c.ShouldBindJSON(&Blood_type); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&Blood_type).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": Blood_type})
}

// GET /video/:id
func GetBlood_type(c *gin.Context) {
	var Blood_type entity.Blood_type

	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM Blood_types WHERE id = ?", id).Find(&Blood_type).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": Blood_type})
}

// GET /videos
func ListBlood_type(c *gin.Context) {
	var Blood_type []entity.Blood_type
	if err := entity.DB().Raw("SELECT * FROM Blood_types").Find(&Blood_type).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": Blood_type})
}

// PATCH /videos
func UpdateBlood_type(c *gin.Context) {
	var Blood_type entity.Blood_type
	if err := c.ShouldBindJSON(&Blood_type); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", Blood_type.ID).First(&Blood_type); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Blood_type not found"})
		return
	}

	if err := entity.DB().Save(&Blood_type).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": Blood_type})
}
