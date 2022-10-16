package controller

import (
	"net/http"

	"github.com/PattanasakGit/project-sa/entity"
	"github.com/gin-gonic/gin"
)

// POST /videos
func CreateDrug_Allergy(c *gin.Context) {
	var Drug_Allergy entity.Drug_Allergy
	if err := c.ShouldBindJSON(&Drug_Allergy); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&Drug_Allergy).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": Drug_Allergy})
}

// GET /video/:id
func GetDrug_Allergy(c *gin.Context) {
	var Drug_Allergy entity.Drug_Allergy

	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM drug_allergies WHERE id = ?", id).Find(&Drug_Allergy).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": Drug_Allergy})
}

// GET /videos
func ListDrug_Allergy(c *gin.Context) {
	var Drug_Allergy []entity.Drug_Allergy
	if err := entity.DB().Raw("SELECT * FROM drug_allergies").Find(&Drug_Allergy).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": Drug_Allergy})
}

// PATCH /videos
func UpdateDrug_Allergy(c *gin.Context) {
	var Drug_Allergy entity.Drug_Allergy
	if err := c.ShouldBindJSON(&Drug_Allergy); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", Drug_Allergy.ID).First(&Drug_Allergy); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Drug_Allergy not found"})
		return
	}

	if err := entity.DB().Save(&Drug_Allergy).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": Drug_Allergy})
}
