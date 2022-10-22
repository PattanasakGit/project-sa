package controller

import (
	"net/http"

	"github.com/PattanasakGit/project-sa/entity"
	"github.com/gin-gonic/gin"
)

// POST / สำรับสร้างข้อมูล
func CreateRIGHTS(c *gin.Context) {
	var RIGHTS entity.RIGHTS
	if err := c.ShouldBindJSON(&RIGHTS); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&RIGHTS).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": RIGHTS})
}

// GET / แบบเฉพาะเจาะจง
func GetRIGHTS(c *gin.Context) {
	var RIGHTS entity.RIGHTS

	id := c.Param("id")
	if err := entity.DB().Preload("Owner").Raw("SELECT * FROM RIGHTS WHERE id = ?", id).Find(&RIGHTS).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": RIGHTS})
}

// GET / สำหรับเรียกดูข้อมูลทั้งหมด เอาไปใช้กับ combobox ใน fontend
func ListRIGHTS(c *gin.Context) {
	var RIGHTS []entity.RIGHTS
	if err := entity.DB().Raw("SELECT * FROM rights").Find(&RIGHTS).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": RIGHTS})
}

// PATCH /สำหรับการอัพเตค่า **แต่ไม่ได้ใช้ในระบบย่อยนี้**
func UpdateRIGHTS(c *gin.Context) {
	var RIGHTS entity.RIGHTS
	if err := c.ShouldBindJSON(&RIGHTS); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", RIGHTS.ID).First(&RIGHTS); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "RIGHTS not found"})
		return
	}

	if err := entity.DB().Save(&RIGHTS).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": RIGHTS})
}
