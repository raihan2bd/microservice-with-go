// package main

// import (
// 	"context"
// 	"log"
// 	"net"
// 	"net/http"
// 	"os"
// 	"os/signal"
// 	"product-service/internal/config"
// 	"product-service/internal/handlers"
// 	"product-service/internal/repository/postgres"
// 	"product-service/internal/service"
// 	"syscall"
// 	"time"

// 	"google.golang.org/grpc"
// )

// func main() {
// 	// Load application configuration
// 	cfg, err := config.LoadConfig()
// 	if err != nil {
// 		log.Fatalf("Failed to load config: %v", err)
// 	}

// 	// type AppConfig struct {
// 	// 	PG_DB      *gorm.DB
// 	// 	RabbitMQ   *amqp.Connection
// 	// 	Redis      *redis.Client
// 	// 	HTTPServer *gin.Engine
// 	// 	GRPCServer *grpc.Server
// 	// }

// 	// return &AppConfig{
// 	// 	PG_DB:      pgDB,
// 	// 	RabbitMQ:   rabbitMQ,
// 	// 	Redis:      redisClient,
// 	// 	HTTPServer: httpServer,
// 	// 	GRPCServer: grpcServer,
// 	// }, nil

// 	// Create a context for graceful shutdown
// 	ctx, cancel := context.WithCancel(context.Background())
// 	defer cancel()

// 	// Channel to listen for OS signals
// 	quit := make(chan os.Signal, 1)
// 	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

// 	// Set up HTTP server
// 	httpServer := &http.Server{
// 		Addr:    ":8080",
// 		Handler: cfg.HTTPServer, // Gin engine
// 	}

// 	// Start HTTP server
// 	httpErrChan := make(chan error, 1)
// 	go func() {
// 		log.Println("Starting HTTP server on :8080")
// 		httpErrChan <- httpServer.ListenAndServe()
// 	}()

// 	// Start gRPC server
// 	grpcErrChan := make(chan error, 1)
// 	go func() {
// 		log.Println("Starting gRPC server on :50051")
// 		grpcErrChan <- startGRPCServer()
// 	}()

// 	// Wait for errors or termination signal
// 	select {
// 	case <-quit:
// 		log.Println("Shutting down...")
// 	case err := <-httpErrChan:
// 		log.Printf("HTTP server error: %v", err)
// 	case err := <-grpcErrChan:
// 		log.Printf("gRPC server error: %v", err)
// 	}

// 	// Gracefully shut down both servers
// 	shutdownHTTPServer(ctx, httpServer)
// 	shutdownDependencies(cfg)
// 	log.Println("Application stopped gracefully")

// 	// Initialize repository based on environment settings (PostgreSQL or MySQL)
// 	repo := postgres.NewPGRepository(cfg.PG_DB)

// 	// Initialize service layer
// 	service := service.NewService(repo)

// 	// Initialize both REST and gRPC handlers
// 	handler := handlers.NewHandler(*service)
// 	// handler.GRPCHandler
// 	// handler.RestHandler
// }

// // startGRPCServer starts the gRPC server
// func startGRPCServer() error {
// 	listener, err := net.Listen("tcp", ":50051")
// 	if err != nil {
// 		return err
// 	}

// 	grpcServer := grpc.NewServer()
// 	grpc_service.RegisterServices(grpcServer) // Custom method to register your gRPC services

// 	return grpcServer.Serve(listener)
// }

// // shutdownHTTPServer gracefully shuts down the HTTP server
// func shutdownHTTPServer(ctx context.Context, server *http.Server) {
// 	shutdownCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
// 	defer cancel()

// 	if err := server.Shutdown(shutdownCtx); err != nil {
// 		log.Printf("HTTP server shutdown error: %v", err)
// 	} else {
// 		log.Println("HTTP server stopped gracefully")
// 	}
// }

// // shutdownDependencies closes database connections and other resources
// func shutdownDependencies(cfg *config.AppConfig) {
// 	// Close PostgreSQL connection
// 	sqlDB, err := cfg.PG_DB.DB()
// 	if err == nil {
// 		if err := sqlDB.Close(); err != nil {
// 			log.Printf("Error closing PostgreSQL connection: %v", err)
// 		} else {
// 			log.Println("PostgreSQL connection closed")
// 		}
// 	}

// 	// Close RabbitMQ connection
// 	if cfg.RabbitMQ != nil {
// 		if err := cfg.RabbitMQ.Close(); err != nil {
// 			log.Printf("Error closing RabbitMQ connection: %v", err)
// 		} else {
// 			log.Println("RabbitMQ connection closed")
// 		}
// 	}

// 	// Close Redis connection
// 	if cfg.Redis != nil {
// 		if err := cfg.Redis.Close(); err != nil {
// 			log.Printf("Error closing Redis connection: %v", err)
// 		} else {
// 			log.Println("Redis connection closed")
// 		}
// 	}
// }

package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"product-service/internal/config"
	"product-service/internal/handlers"
	grpcHandler "product-service/internal/handlers/grpc" // Assuming this is the correct import for the gRPC handler
	"product-service/internal/repository/postgres"
	"product-service/internal/routes"
	"product-service/internal/service"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Load application configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Create a context for graceful shutdown
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Channel to listen for OS signals
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	// Initialize repository based on environment settings (PostgreSQL or MySQL)
	repo := postgres.NewPGRepository(cfg.PG_DB)

	// Initialize service layer
	svc := service.NewService(repo)

	// Initialize both REST and gRPC handlers
	handlers := handlers.NewHandler(*svc)
	// restHandler := handlers.NewHandler(*svc)

	// Assign REST routes to the Gin HTTP server
	routes := routes.NewRoutes(cfg.HTTPServer, handlers.RestHandler)
	routes.InitRoutes()
	// cfg.HTTPServer = restHandler.SetupRoutes(cfg.HTTPServer)

	// Channel for errors
	errChan := make(chan error, 2)

	// Start HTTP server
	go func() {
		log.Println("Starting HTTP server on :8080")
		errChan <- cfg.HTTPServer.Run() // Use Gin's Run() directly
	}()

	// // Start gRPC server
	// go func() {
	// 	log.Println("Starting gRPC server on :50051")
	// 	errChan <- startGRPCServer(handlers.GRPCHandler)
	// }()

	// Wait for termination signal or server errors
	select {
	case <-quit:
		log.Println("Shutting down...")
	case err := <-errChan:
		log.Printf("Server error: %v", err)
	}

	// Gracefully shut down the HTTP server
	shutdownHTTPServer(ctx, cfg.HTTPServer)
	shutdownDependencies(cfg)
	log.Println("Application stopped gracefully")
}

// startGRPCServer starts the gRPC server
func startGRPCServer(handler *grpcHandler.GRPCHandler) error {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()
	// Register the gRPC handler
	// handler.Service.Register(grpcServer, handler) // Replace this line if your protogen-generated code requires specific registration

	return grpcServer.Serve(listener)
}

// shutdownHTTPServer gracefully shuts down the HTTP server
func shutdownHTTPServer(ctx context.Context, server *gin.Engine) {
	// shutdownCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	// defer cancel()

	// if err := server.Shutdown(shutdownCtx); err != nil {
	// 	log.Printf("HTTP server shutdown error: %v", err)
	// } else {
	// 	log.Println("HTTP server stopped gracefully")
	// }
}

// shutdownDependencies closes database connections and other resources
func shutdownDependencies(cfg *config.AppConfig) {
	// Close PostgreSQL connection
	sqlDB, err := cfg.PG_DB.DB()
	if err == nil {
		if err := sqlDB.Close(); err != nil {
			log.Printf("Error closing PostgreSQL connection: %v", err)
		} else {
			log.Println("PostgreSQL connection closed")
		}
	}

	// // Close RabbitMQ connection
	// if cfg.RabbitMQ != nil {
	// 	if err := cfg.RabbitMQ.Close(); err != nil {
	// 		log.Printf("Error closing RabbitMQ connection: %v", err)
	// 	} else {
	// 		log.Println("RabbitMQ connection closed")
	// 	}
	// }

	// // Close Redis connection
	// if cfg.Redis != nil {
	// 	if err := cfg.Redis.Close(); err != nil {
	// 		log.Printf("Error closing Redis connection: %v", err)
	// 	} else {
	// 		log.Println("Redis connection closed")
	// 	}
	// }
}
