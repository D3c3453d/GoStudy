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

func (h *Handler) phoneByName(c *gin.Context) {

}

func (h *Handler) descByName(c *gin.Context) {

}

func (h *Handler) showByName(c *gin.Context) {
	name := c.Param("name")

	list, err := h.services.AccountsServiceI.GetByName(name)
	if err != nil {
		handler.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, list)
}

func (h *Handler) findByPhone(c *gin.Context) {
	phone := c.Param("phone")

	list, err := h.services.AccountsServiceI.GetByPhone(phone)
	if err != nil {
		handler.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, list)

}
