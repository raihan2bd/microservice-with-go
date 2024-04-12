package application

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/raihan2bd/microservice-with-go/handler"
)

func loadRoutes() *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	router.Route("/orders", loadOrderRoutes)
	return router
}

func loadOrderRoutes(r chi.Router) {
	orderHandler := &handler.Order{}
	r.Post("/", orderHandler.Create)
	r.Get("/", orderHandler.List)
	r.Get("/{id}", orderHandler.GetByID)
	r.Put("/{id}", orderHandler.UpdateByID)
	r.Delete("/{id}", orderHandler.DeleteByID)
}
