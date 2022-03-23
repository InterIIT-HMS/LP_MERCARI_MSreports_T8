package main

import (
	"github.com/wryonik/appointment/controllers"
	"github.com/wryonik/appointment/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Connect to database
	models.ConnectDatabase()

	// Routes
	r.GET("/reports", controllers.FindReports)
	r.POST("/reports", controllers.CreateReport)
	r.PATCH("/reports", controllers.UpdateReport)
	r.DELETE("/reports", controllers.DeleteReport)

	// Run the server
	r.Run()
}
