package routes

import (
	"violationdetails/repository"
	"violationdetails/rest"
	"violationdetails/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func HandleViolationRequests(r *gin.Engine, db *gorm.DB) {
	regCtrl := rest.NewViolationController(service.NewViolationService(repository.NewViolationRepo(db)))

	r.GET("/api/v1/violation/:vehregno", regCtrl.Violation)
}
