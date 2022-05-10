package model

import "github.com/gin-gonic/gin"

type Trafficviolationsystem struct {
	Trafficviolationid uint64 `gorm:"primary_key:auto_increment" json:"trafficviolationid,omitempty"`
	Violationdetailid  uint64 `gorm:"type:bigint(11)" json:"violationdetailid"`
	Vehicleregno       string `gorm:"type:varchar(30)" json:"vehicleregno"`
	City               string `gorm:"type:varchar(50)" json:"city"`
	Longitude          string `gorm:"type:varchar(10)" json:"longitude"`
	Latitude           string `gorm:"type:varchar(10)" json:"latitude"`
	State              string `gorm:"type:varchar(50)" json:"state"`
	Devicetype         string `gorm:"type:varchar(50)" json:"devicetype"`
	// Violationdetails   *[]Violationdetails `gorm:"-" json:"violationdetails"`
}

type Violationdetails struct {
	Violationdetailid uint64 `gorm:"primary_key:auto_increment" json:"violationdetailid,omitempty"`
	Name              string `gorm:"type:varchar(50)" json:"name"`
	Code              string `gorm:"type:varchar(30)" json:"code"`
	Charge            uint64 `gorm:"type:double" json:"charge"`
	Description       string `gorm:"type:varchar(255)" json:"description"`
	// Violationdetails  Trafficviolationsystem `gorm:"foreignkey:Violationdetailid;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"trafficviolationsystem"`
}
type TVS struct {
	Trafficviolation Trafficviolationsystem `json:"trafficviolation"`
	Violations       Violationdetails       `json:"violations"`
}

type ViolationService interface {
	GetViolationDetails(VehicleRegNo string) ([]TVS, error)
}

type ViolationRepository interface {
	GetViolationDetails(VehicleRegNo string) ([]TVS, error)
}

type ViolationController interface {
	Violation(c *gin.Context)
}
