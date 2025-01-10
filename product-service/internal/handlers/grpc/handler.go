package grpc

import (
	"product-service/internal/service"
	"product-service/protogen"
	// Import the generated protobuf Go code
)

type GRPCHandler struct {
	Service service.Service
	protogen.UnimplementedProductServiceServer
}

func NewGRPCHandler(service service.Service) *GRPCHandler {
	return &GRPCHandler{
		Service: service,
	}
}

// Implement the GetProductByID gRPC method
// func (h *ProductHandler) GetProductByID(ctx context.Context, req *protogen.Empty) (*protogen.Product, error) {
// 	log.Println("gRPC: Get product by ID")

// 	// Replace with your service logic to fetch product by ID
// 	// For now, I'm just returning a dummy product for demonstration
// 	product := &protogen.Product{
// 		Id:          "123",
// 		Name:        "Sample Product",
// 		Description: "This is a sample product description",
// 	}

// 	return product, nil
// }
