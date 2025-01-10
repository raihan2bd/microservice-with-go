package postgres

import (
	"gorm.io/gorm"
)

// SQLProductRepository is the PostgreSQL implementation of the ProductRepository
type PGRepo struct {
	DB *gorm.DB
}

// NewSQLProductRepository creates a new instance of PostgreSQL ProductRepository
func NewPGRepository(db *gorm.DB) *PGRepo {
	return &PGRepo{
		DB: db,
	}
}
