package controller

import (
	"net/http"

	"github.com/PattanasakGit/project-sa/entity"
	"github.com/gin-gonic/gin"
)

// POST /สำหรับสร้างข้อมูล
func CreateGender(c *gin.Context) {
	var Gender entity.Gender
	if err := c.ShouldBindJSON(&Gender); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&Gender).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": Gender})
}

// GET / แบบเฉพาพเจาะจง
func GetGender(c *gin.Context) {
	var Gender entity.Gender

	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM Genders WHERE id = ?", id).Find(&Gender).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": Gender})
}

// GET /
func ListGender(c *gin.Context) {
	var Gender []entity.Gender
	if err := entity.DB().Raw("SELECT * FROM Genders").Find(&Gender).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": Gender})
}

// PATCH / สำหรับการอัพเตค่า **แต่ไม่ได้ใช้ในระบบย่อยนี้**
func UpdateGender(c *gin.Context) {
	var Gender entity.Gender
	if err := c.ShouldBindJSON(&Gender); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", Gender.ID).First(&Gender); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Gender not found"})
		return
	}

	if err := entity.DB().Save(&Gender).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": Gender})
}
