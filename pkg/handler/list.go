package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) getLists(c *gin.Context) {

}

func (h *Handler) getListByID(c *gin.Context) {

}

func (h *Handler) createList(c *gin.Context) {
	id, ok := c.Get(CtxUserID)
	if ok {
		h.utilities.ResponseHandler.CommonResponseJSON(c, http.StatusOK, "id", id)

		return
	}
	h.utilities.ResponseHandler.ErrorResponseJSON(c, http.StatusUnauthorized, "Something went wrong")
}

func (h *Handler) updateList(c *gin.Context) {

}

func (h *Handler) deleteList(c *gin.Context) {

}
