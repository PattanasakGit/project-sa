package controller

import (
	"net/http"

	"github.com/PattanasakGit/project-sa/entity"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {

	var user entity.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})

}

// // GET /user/:id
// func GetUser(c *gin.Context) {
// 	var user entity.User
// 	id := c.Param("id")
// 	if err := entity.DB().Raw("SELECT * FROM users WHERE id = ?", id).Scan(&user).Error; err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"data": user})
// }

// // GET /usrs
// func ListUsers(c *gin.Context {
// 	var users []entity.Uer
// 	if err := entity.DB().Raw("SELECT * FROM users").Scan(&users).Error; err != ni {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(})
// 		retrn
//	}

// 	c.JSON(http.StatusOK, gin.H{"data": user})
// }
