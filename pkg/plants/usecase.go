package plants

import (
	"solar-service/pkg/plants/repository"
)

type plantUsecase struct {
	plantsRepo       repository.PlantsRepo
}

// ClantUsecase ...
type PlantUsecase interface {
	//
}

// NewCategoriesUsecase ...
func NewPlantUsecase(plantsRepo repository.PlantsRepo) PlantUsecase {
	return &plantUsecase{
		plantsRepo,
	}
}