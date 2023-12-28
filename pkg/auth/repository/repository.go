package repository

import (
	"github.com/jmoiron/sqlx"
)

type authRepo struct{
	db *sqlx.DB
}

// PlantsRepo ...
type AuthRepo interface {
	//
}

// NewPlantsRepo ...
func NewAuthRepo(db *sqlx.DB) AuthRepo {
	return &authRepo{
		db,
	}
}