package config

import (
	"product-service/internal/initializers"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

// AppConfig holds all dependencies
type AppConfig struct {
	PG_DB *gorm.DB
	// RabbitMQ   *amqp.Connection
	// Redis      *redis.Client
	HTTPServer *gin.Engine
	GRPCServer *grpc.Server
}

// LoadConfig initializes and returns AppConfig
func LoadConfig() (*AppConfig, error) {

	// PostgreSQL connection
	pgDB, err := initializers.ConnectToPG_DB()
	if err != nil {
		return nil, err
	}

	// // RabbitMQ connection
	// rabbitMQ, err := initializers.ConnectToRabbitMQ()
	// if err != nil {
	// 	return nil, err
	// }

	// // Redis connection
	// redisClient, err := initializers.ConnectToRedis()
	// if err != nil {
	// 	return nil, err
	// }

	// Initialize gRPC server
	grpcServer := grpc.NewServer()

	// HTTP server initialization
	httpServer := gin.New()

	return &AppConfig{
		PG_DB: pgDB,
		// RabbitMQ:   rabbitMQ,
		// Redis:      redisClient,
		HTTPServer: httpServer,
		GRPCServer: grpcServer,
	}, nil
}
