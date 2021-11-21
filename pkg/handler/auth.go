package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/max99xam/todo-app"
)

func (h *Handler) signUp(context *gin.Context) {
	var input todo.User

	if err := context.BindJSON(&input); err != nil {
		newErrorResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	context.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) signIn(c *gin.Context) {

}
