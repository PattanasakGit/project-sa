package controller

import (
	"net/http"
	//"github.com/PattanasakGit/project-sa/backend/entity"
	"github.com/PattanasakGit/project-sa/entity"
	"github.com/gin-gonic/gin"
)

// POST /patients
func CreatePatient(c *gin.Context) {
	var patient entity.Patient
	var Blood_type entity.Blood_type
	var Gender entity.Gender
	var Drug_Allergy entity.Drug_Allergy
	var RIGHTS entity.RIGHTS

	// ผลลัพธ์ที่ได้จากขั้นตอนที่ 10 จะถูก bind เข้าตัวแปร Patient
	if err := c.ShouldBindJSON(&patient); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 11: ค้นหา เพศผู้ป่วย ด้วย id
	if tx := entity.DB().Where("id = ?", patient.GenderID).First(&Gender); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Gender not found"})
		return
	}
	// 12: ค้นหา หมู่เลือดผู้ป่วย ด้วย id
	if tx := entity.DB().Where("id = ?", patient.Blood_typeID).First(&Blood_type); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Blood_type not found"})
		return
	}
	// 13: ค้นหา ยาที่ผู้ป่วยแพ้ ด้วย id
	if tx := entity.DB().Where("id = ?", patient.Drug_AllergyID).First(&Drug_Allergy); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Drug_AllergyID not found"})
		return
	}
	// 14: ค้นหา สิทธิ์การรักษา ด้วย id
	if tx := entity.DB().Where("id = ?", patient.RIGHTSID).First(&RIGHTS); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "RIGHTSID not found"})
		return
	}
	// 15: ค้นหา สิทธิ์การรักษา ด้วย id
	pt := entity.Patient{
		ID_Card:       patient.ID_Card,
		Patient_Name:  patient.Patient_Name,
		Date_of_Birth: patient.Date_of_Birth,
		User:          patient.User,
		Gender:        Gender,
		Blood_type:    Blood_type,
		Drug_Allergy:  Drug_Allergy,
		RIGHTS:        RIGHTS,
		Addess:        patient.Addess,
		Other:         patient.Other,
	}
	// 16: บันทึก
	if err := entity.DB().Create(&pt).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": pt})

}

// GET /patient/:id
func GetPatient(c *gin.Context) {
	var patient entity.Patient
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM patients WHERE id = ?", id).Scan(&patient).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": patient})
}

// GET patient
func Listpatient(c *gin.Context) {
	var patient []entity.Patient
	if err := entity.DB().Raw("SELECT * FROM patient").Scan(&patient).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": patient})
}
