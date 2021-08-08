package handler

import (
	"github.com/HunkevychPhilip/todo/pkg/service"
	"github.com/HunkevychPhilip/todo/pkg/utils"
)

type Handler struct {
	services  *service.Service
	utilities *utils.Utils
}

func NewHandler(services *service.Service, utils *utils.Utils) *Handler {
	return &Handler{
		services:  services,
		utilities: utils,
	}
}
