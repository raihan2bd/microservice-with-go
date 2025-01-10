package handlers

import (
	"product-service/internal/handlers/grpc"
	"product-service/internal/handlers/rest"
	"product-service/internal/service"
)

// Handler struct contains all handlers for both REST and gRPC services
type Handler struct {
	RestHandler *rest.RestHandler
	GRPCHandler *grpc.GRPCHandler
}

// NewHandler initializes both REST and gRPC handlers with the service layer
func NewHandler(service service.Service) *Handler {
	return &Handler{
		RestHandler: rest.NewRestHandler(service),
		GRPCHandler: grpc.NewGRPCHandler(service),
	}
}
