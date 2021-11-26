package handler

import (
	"errors"
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

func getUserId(context *gin.Context) (int, error) {
	id, ok := context.Get(userContext)
	if !ok {
		newErrorResponse(context, http.StatusInternalServerError, "User id is not found.")
		return 0, errors.New("User id not found.")
	}

	idInt, ok := id.(int)
	if !ok {
		newErrorResponse(context, http.StatusInternalServerError, "Invalid type of user id.")
		return 0, errors.New("User id not found.")
	}

	return idInt, nil
}
