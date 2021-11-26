package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/max99xam/todo-app"
)

func (h *Handler) createList(context *gin.Context) {
	userId, err := getUserId(context)
	if err != nil {
		return
	}

	var input todo.TodoList
	if err := context.BindJSON(&input); err != nil {
		newErrorResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.TodoList.Create(userId, input)
	if err != nil {
		newErrorResponse(context, http.StatusInternalServerError, err.Error())
	}

	context.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type getAllListsResponse struct {
	Data []todo.TodoList `json:"data"`
}

func (h *Handler) getAllLists(context *gin.Context) {
	userId, err := getUserId(context)
	if err != nil {
		return
	}

	lists, err := h.services.TodoList.GetAll(userId)
	if err != nil {
		newErrorResponse(context, http.StatusInternalServerError, err.Error())
	}

	context.JSON(http.StatusOK, getAllListsResponse{
		Data: lists,
	})
}

func (h *Handler) updateList(context *gin.Context) {

}

func (h *Handler) getListById(context *gin.Context) {
	userId, err := getUserId(context)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		newErrorResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	list, err := h.services.TodoList.GetById(userId, id)
	if err != nil {
		newErrorResponse(context, http.StatusInternalServerError, err.Error())
	}

	context.JSON(http.StatusOK, list)
}

func (h *Handler) deleteList(c *gin.Context) {

}
