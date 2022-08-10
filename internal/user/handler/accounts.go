package handler

import (
	"GoStudy/internal/config"
	"GoStudy/internal/user/entity"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) help(c *gin.Context, cfg *config.Commands) {
	h.services.Help(cfg)
}

func (h *Handler) add(c *gin.Context) {
	var input entity.Account

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err := h.services.Create(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "added",
	})
}

type allResponse struct {
	Data []entity.Account `json:"accounts"`
}

func (h *Handler) all(c *gin.Context) {
	list, err := h.services.AccountsServiceI.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, allResponse{
		Data: list,
	})
}

func (h *Handler) phone(c *gin.Context) {

}

func (h *Handler) desc(c *gin.Context) {

}

func (h *Handler) show(c *gin.Context) {

}

func (h *Handler) find(c *gin.Context) {
	phone := c.Param("phone")

	list, err := h.services.AccountsServiceI.GetByPhone(phone)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, list)

}
