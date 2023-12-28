package database

import (
	"solar-service/models"

	"github.com/jmoiron/sqlx"
)

// DatabaseRepo interface
type DatabaseRepo interface {
	Connect(*models.Config) (*sqlx.DB, error)
}
