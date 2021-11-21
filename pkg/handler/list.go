package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createList(context *gin.Context) {
	id, _ := context.Get(userContext)
	context.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getAllLists(c *gin.Context) {

}

func (h *Handler) updateList(c *gin.Context) {

}

func (h *Handler) getListById(c *gin.Context) {

}

func (h *Handler) deleteList(c *gin.Context) {

}
