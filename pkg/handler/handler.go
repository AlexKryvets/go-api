package handler

import "github.com/gin-gonic/gin"

type Handler struct {
}

func (h *Handler) InitRouters() *gin.Engine {
	router := gin.New()

	test := router.Group("/test")
	{
		test.GET("/1", h.test)
	}

	return router
}
