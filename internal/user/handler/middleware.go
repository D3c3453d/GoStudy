package handler

import (
	"GoStudy/pkg/handler"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	authHeader = "Authorization"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authHeader)
	if header == "" {
		handler.NewErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}
	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		handler.NewErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
		return
	}
	userId, err := h.services.AuthServiceI.ParseToken(headerParts[1])
	if err != nil {
		handler.NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	c.Set("userId", userId)
}
