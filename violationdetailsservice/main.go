package main

import (
	"violationdetails/config"
	"violationdetails/routes"
	"violationdetails/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func initializeDbAndRoutes(r *gin.Engine) {
	var (
		db *gorm.DB = config.SetupDBConnection()
	)

	// defer config.CloseDBConnection(db)
	routes.HandleViolationRequests(r, db)

	conf := utils.NewConfig()
	PORT := conf.Database.Port
	r.Run(PORT)
}

func main() {

	r := gin.Default()
	initializeDbAndRoutes(r)

}
