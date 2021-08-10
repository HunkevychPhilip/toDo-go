package handler

import (
	"github.com/HunkevychPhilip/todo/pkg/types"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) getLists(c *gin.Context) {

}

func (h *Handler) getListByID(c *gin.Context) {

}

func (h *Handler) createList(c *gin.Context) {
	userID, err := h.getUserID(c)
	if err != nil {
		return
	}

	var list types.List
	err = c.ShouldBindJSON(&list)
	if err != nil {
		h.utilities.ResponseHandler.ErrorResponseJSON(c, http.StatusBadRequest, err.Error())

		return
	}

	listID, err := h.services.List.Create(userID, &list)
	if err != nil {
		h.utilities.ResponseHandler.ErrorResponseJSON(c, http.StatusInternalServerError, err.Error())

		return
	}

	h.utilities.ResponseHandler.CommonResponseJSON(c, http.StatusOK, "list_id", listID)
}

func (h *Handler) updateList(c *gin.Context) {

}

func (h *Handler) deleteList(c *gin.Context) {

}
