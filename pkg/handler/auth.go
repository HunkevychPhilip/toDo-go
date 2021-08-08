package handler

import (
	"github.com/HunkevychPhilip/todo/pkg/types"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) signIn(c *gin.Context) {

}

func (h *Handler) signUp(c *gin.Context) {
	var user types.User

	if err := c.BindJSON(&user); err != nil {
		h.utilities.ResponseHandler.ErrorResponseJSON(c, http.StatusBadRequest, err.Error())

		return
	}

	id, err := h.services.Authorization.CreateUser(&user)
	if err != nil {
		h.utilities.ResponseHandler.ErrorResponseJSON(c, http.StatusInternalServerError, err.Error())

		return
	}

	h.utilities.ResponseHandler.CommonResponseJSON(c, http.StatusOK, "id", id)
}
