package service

import (
	"context"
	"log"
	"product-service/internal/repository"
)

// Service contains the methods for handling business logic
type Service struct {
	Repo repository.ProductRepository
}

// NewService creates a new instance of the service layer
func NewService(repo repository.ProductRepository) *Service {
	return &Service{
		Repo: repo,
	}
}

// Example method
func (s *Service) GetProductByID(ctx context.Context, id int64) (interface{}, error) {
	// Logic to fetch product from repository
	log.Println("Fetching product by ID")
	return nil, nil
}
