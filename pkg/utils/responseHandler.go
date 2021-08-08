package utils

import (
	"github.com/HunkevychPhilip/todo/pkg/types"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type responseHandler struct {
}

func NewErrHandler() *responseHandler {
	return &responseHandler{}
}

func (r *responseHandler) ErrorResponseJSON(c *gin.Context, statusCode int, msg string) {
	logrus.Error(msg)

	c.AbortWithStatusJSON(statusCode, types.Error{
		Status:  statusCode,
		Message: msg,
	})
}

func (r *responseHandler) CommonResponseJSON(c *gin.Context, statusCode int, key string, val interface{}) {
	c.JSON(statusCode, map[string]interface{}{
		key: val,
	})
}
