package handler

import (
	"GoStudy/internal/user/entity"
	"GoStudy/pkg/handler"
	"github.com/gin-gonic/gin"
	"net/http"
)

type allResponse struct {
	Data []entity.Account `json:"accounts"`
}

func (h *Handler) all(c *gin.Context) {
	list, err := h.services.AccountsServiceI.GetAll()
	if err != nil {
		handler.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
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
	name := c.Param("name")

	list, err := h.services.AccountsServiceI.GetByName(name)
	if err != nil {
		handler.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, list)
}

func (h *Handler) find(c *gin.Context) {
	phone := c.Param("phone")

	list, err := h.services.AccountsServiceI.GetByPhone(phone)
	if err != nil {
		handler.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, list)

}
