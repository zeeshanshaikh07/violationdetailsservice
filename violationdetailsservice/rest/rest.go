package rest

import (
	"fmt"
	"net/http"
	"violationdetails/model"
	"violationdetails/utils"

	"github.com/gin-gonic/gin"
)

type violationController struct {
	violationService model.ViolationService
}

func NewViolationController(violationService model.ViolationService) *violationController {
	return &violationController{
		violationService: violationService,
	}
}

func (ctrl *violationController) Violation(c *gin.Context) {
	fmt.Println("inside violation")
	vno := c.Param("vehregno")
	if vno == "" {
		response := utils.BuildResponse("Failed to get id", 404, utils.EmptyObj{})
		c.JSON(http.StatusBadRequest, response)
	}
	obj, err := ctrl.violationService.GetViolationDetails(vno)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Server error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Violation": obj})
}
