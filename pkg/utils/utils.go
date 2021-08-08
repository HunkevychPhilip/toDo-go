package utils

import (
	"github.com/gin-gonic/gin"
)

type ResponseHandler interface {
	ErrorResponseJSON(c *gin.Context, statusCode int, msg string)
	CommonResponseJSON(c *gin.Context, statusCode int, key string, val interface{})
}

type DBErrorHandler interface {
}

type Utils struct {
	ResponseHandler
}

func NewUtils(rh ResponseHandler) *Utils {
	return &Utils{
		ResponseHandler: rh,
	}
}
