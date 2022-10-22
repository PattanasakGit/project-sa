package controller

import (
	"net/http"

	"github.com/PattanasakGit/project-sa/entity"
	"github.com/gin-gonic/gin"
)

// POST/Blood_types สำหรับ สร้างข้อมูล
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

// GET/Blood_types/:id
func GetBlood_type(c *gin.Context) {
	var Blood_type entity.Blood_type

	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM Blood_types WHERE id = ?", id).Find(&Blood_type).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": Blood_type})
}

// GET /Blood_types สำหรับเรียกดูข้อมูลทั้งหมด เอาไปใช้กับ combobox ใน fontend
func ListBlood_type(c *gin.Context) {
	var Blood_type []entity.Blood_type
	if err := entity.DB().Raw("SELECT * FROM Blood_types").Find(&Blood_type).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": Blood_type})
}

// PATCH /Blood_types  สำหรับการอัพเตค่า **แต่ไม่ได้ใช้ในระบบย่อยนี้**
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
