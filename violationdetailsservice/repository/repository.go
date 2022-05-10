package repository

import (
	"errors"
	"fmt"
	"violationdetails/model"

	"gorm.io/gorm"
)

type violationConnection struct {
	connection *gorm.DB
}

func NewViolationRepo(db *gorm.DB) *violationConnection {
	return &violationConnection{
		connection: db,
	}
}

func (repo *violationConnection) GetViolationDetails(VehicleRegNo string) ([]model.TVS, error) {
	tvs := []model.TVS{}
	trafficviolationsystem := []model.Trafficviolationsystem{}
	violationdetails := model.Violationdetails{}
	result := repo.connection.Where("vehicleregno = ?", VehicleRegNo).Find(&trafficviolationsystem)

	fmt.Println("result", result.RowsAffected)
	fmt.Println("Trafficviolationsystem", trafficviolationsystem)
	if result.RowsAffected == 0 {
		return tvs, errors.New("no record found in violation details table")
	}

	if result.Error != nil {
		return tvs, result.Error
	}
	var i int64

	for i = 0; i <= result.RowsAffected; i++ {
		// fmt.Println(trafficviolationsystem[i])
		result = repo.connection.Where("violationdetailsid = ?", trafficviolationsystem[i].Violationdetailid).Find(&violationdetails)
		fmt.Println("VIOLATION", violationdetails)
		if result.RowsAffected == 0 {
			return tvs, errors.New("no record found in violation details table")
		}

		if result.Error != nil {
			return tvs, result.Error
		}
		tmp := model.TVS{
			Trafficviolation: trafficviolationsystem[i],
			Violations:       violationdetails,
		}
		tvs = append(tvs, tmp)

	}
	return tvs, result.Error
}
