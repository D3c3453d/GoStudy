package handler

import (
	"GoStudy/internal/user/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	accounts := router.Group("/accounts")
	{
		accounts.POST("/", h.add)
		accounts.GET("/", h.all)
		accounts.GET("/phone/:phone", h.find)
		accounts.GET("/name/:name", h.show)
	}
	return router
}
