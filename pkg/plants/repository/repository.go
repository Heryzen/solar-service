package repository

import (
	"context"
	"log"
	"solar-service/models"

	"github.com/jmoiron/sqlx"
)

type plantsRepo struct{
	db *sqlx.DB
}

// PlantsRepo ...
type PlantsRepo interface {
	Create(ctx context.Context, req *models.Plants) error
	Update(ctx context.Context, req *models.Plants) error
}

// NewPlantsRepo ...
func NewPlantsRepo(db *sqlx.DB) PlantsRepo {
	return &plantsRepo{
		db,
	}
}

// Create Transact Record ...
func (r *plantsRepo) Create(ctx context.Context, req *models.Plants) error {
	_, err := r.db.ExecContext(ctx, `
		INSERT INTO public.plants (
			id,
			created_at
		)
		VALUES (
			$1, 
			CURRENT_TIMESTAMP
		)
	`,
		req.ID,
	)

	if err != nil {
		log.Println("[err] [plantsRepo] [Create] => ", err)
		return err
	}

	return nil
}

func (r *plantsRepo) Update(ctx context.Context, req *models.Plants) error {
	_, err := r.db.ExecContext(ctx, `
		UPDATE public.plants SET
			status = $1
		WHERE id = $2
	`,
		req.Status,
		req.ID,
	)

	if err != nil {
		log.Println("[err] [plantsRepo] [Update] => ", err)
		return err
	}

	return nil
}