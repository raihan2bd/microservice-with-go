package rest

import (
	"net/http"
	"product-service/internal/service"

	"github.com/gin-gonic/gin"
)

type RestHandler struct {
	Service service.Service
}

func NewRestHandler(service service.Service) *RestHandler {
	return &RestHandler{
		Service: service,
	}
}

// Example REST handler
func (h *RestHandler) GetProductByID(c *gin.Context) {
	// Your logic to handle the GET request for fetching product by ID
	c.JSON(http.StatusOK, "Get product by ID")
}
