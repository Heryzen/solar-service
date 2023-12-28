package controller

import "solar-service/pkg/plants"

type plantsController struct {
	plantsUc plants.PlantUsecase
}

// PlantsController interface
type PlantsController interface {
	//
}

// NewPlantsController func
func NewPlantsController(plantsUc plants.PlantUsecase) PlantsController {
	return &plantsController{
		plantsUc,
	}
}