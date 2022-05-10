package service

import (
	"violationdetails/model"
)

type violationService struct {
	violationRepository model.ViolationRepository
}

func NewViolationService(violationRep model.ViolationRepository) *violationService {
	return &violationService{
		violationRepository: violationRep,
	}
}

func (s *violationService) GetViolationDetails(VehicleRegNo string) ([]model.TVS, error) {
	return s.violationRepository.GetViolationDetails(VehicleRegNo)
}
