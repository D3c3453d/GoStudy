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
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api")
	{
		api.GET("/", h.all)

		accounts := api.Group("/accounts")
		{
			//accounts.POST("/", h.add)
			accounts.GET("/", h.all)
			accounts.GET("/findByPhone/:phone", h.findByPhone)
			accounts.GET("/showByName/:name", h.showByName)
			accounts.GET("/descByName/:name", h.descByName)
			accounts.GET("/phoneByName/:name", h.phoneByName)
		}
	}

	return router
}
