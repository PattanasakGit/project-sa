package main

import (
	"github.com/PattanasakGit/project-sa/controller"
	"github.com/PattanasakGit/project-sa/entity"
	"github.com/gin-gonic/gin"
)

func main() {

	entity.SetupDatabase()
	r := gin.Default()
	r.Use(CORSMiddleware())

	r.GET("/ListBlood_type", controller.ListBlood_type)
	r.POST("/CreateBlood_type", controller.CreateBlood_type)

	r.GET("/ListDrug_Allergy", controller.ListDrug_Allergy)
	r.POST("/CreateDrug_Allergy", controller.CreateDrug_Allergy)

	r.GET("/ListGender", controller.ListGender)
	r.POST("/CreateGender", controller.CreateGender)

	r.GET("/ListRIGHTS", controller.ListRIGHTS)
	r.POST("/CreateRIGHTS", controller.CreateRIGHTS)

	//r.GET("/GetPatient", controller.list)
	r.GET("/patient/:id", controller.GetPatient)
	r.POST("/CreatePatient", controller.CreatePatient)

	r.Run()
}
func CORSMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}
