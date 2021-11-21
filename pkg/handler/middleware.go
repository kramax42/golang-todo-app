package handler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
	userContext         = "userId"
)

func (h *Handler) userIdentity(context *gin.Context) {
	header := context.GetHeader(authorizationHeader)

	if header == "" {
		newErrorResponse(context, http.StatusUnauthorized, "Empty auth header.")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		newErrorResponse(context, http.StatusUnauthorized, "Invalid auth header")
		return
	}

	userId, err := h.services.ParseToken(headerParts[1])

	if err != nil {
		newErrorResponse(context, http.StatusUnauthorized, err.Error())
		return
	}

	context.Set(userContext, userId)
}
