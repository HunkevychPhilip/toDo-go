package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	authorizationHeaderKey = "Authorization"
	CtxUserID              = "userID"
)

func (h *Handler) userIdentity(c *gin.Context) {
	auth := c.GetHeader(authorizationHeaderKey)
	if len(auth) == 0 {
		h.utilities.ResponseHandler.ErrorResponseJSON(c, http.StatusUnauthorized, "Auth header is missing.")

		return
	}

	authParts := strings.Split(auth, " ")
	if len(authParts) != 2 {
		h.utilities.ResponseHandler.ErrorResponseJSON(c, http.StatusUnauthorized, "Wrong auth header.")

		return
	}

	userID, err := h.services.Authorization.ParseToken(authParts[1])
	if err != nil {
		h.utilities.ResponseHandler.ErrorResponseJSON(c, http.StatusUnauthorized, err.Error())

		return
	}

	c.Set(CtxUserID, userID)
}
