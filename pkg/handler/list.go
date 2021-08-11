package handler

import (
	"github.com/HunkevychPhilip/todo/pkg/types"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) getUserLists(c *gin.Context) {
	userID, err := h.getUserID(c)
	if err != nil {
		return
	}

	lists, err := h.services.List.GetUserLists(userID)
	if err != nil {
		h.utilities.ResponseHandler.ErrorResponseJSON(c, http.StatusInternalServerError, err.Error())

		return
	}

	h.utilities.ResponseHandler.CommonResponseJSON(c, http.StatusOK, "user_lists", lists)
}

func (h *Handler) getListByID(c *gin.Context) {
	userID, err := h.getUserID(c)
	if err != nil {
		return
	}

	val := c.Param("id")
	if val == "" {
		h.utilities.ResponseHandler.ErrorResponseJSON(c, http.StatusBadRequest, "id param is missing")

		return
	}
	listID, err := strconv.Atoi(val)
	if err != nil {
		h.utilities.ResponseHandler.ErrorResponseJSON(c, http.StatusBadRequest, "id param is of invalid type")

		return
	}

	list, err := h.services.List.GetList(userID, listID)
	if err != nil {
		h.utilities.ResponseHandler.ErrorResponseJSON(c, http.StatusInternalServerError, err.Error())

		return
	}

	h.utilities.ResponseHandler.CommonResponseJSON(c, http.StatusOK, "list", list)
}

func (h *Handler) createList(c *gin.Context) {
	userID, err := h.getUserID(c)
	if err != nil {
		return
	}

	list := new(types.List)
	err = c.ShouldBindJSON(&list)
	if err != nil {
		h.utilities.ResponseHandler.ErrorResponseJSON(c, http.StatusBadRequest, err.Error())

		return
	}

	listID, err := h.services.List.Create(userID, list)
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
