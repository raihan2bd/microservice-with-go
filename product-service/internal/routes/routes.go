package routes

import (
	"product-service/internal/handlers/rest"

	"github.com/gin-gonic/gin"
)

type Routes struct {
	mux      *gin.Engine
	handlers *rest.RestHandler
}

func NewRoutes(mux *gin.Engine, h *rest.RestHandler) *Routes {
	return &Routes{mux: mux, handlers: h}
}

func (r *Routes) InitRoutes() {
	r.productRoutes()
}
