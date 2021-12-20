package handler

import (
	"github.com/AlexKryvets/go-api/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRouters() *gin.Engine {
	router := gin.New()

	test := router.Group("/test")
	{
		test.GET("/1", h.test)
	}

	return router
}
