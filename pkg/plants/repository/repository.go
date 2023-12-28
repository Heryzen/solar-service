package repository

import (
	"github.com/jmoiron/sqlx"
)

type plantsRepo struct{
	db *sqlx.DB
}

// PlantsRepo ...
type PlantsRepo interface {
	//
}

// NewPlantsRepo ...
func NewPlantsRepo(db *sqlx.DB) PlantsRepo {
	return &plantsRepo{
		db,
	}
}