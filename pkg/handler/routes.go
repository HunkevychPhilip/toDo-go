package handler

import "github.com/gin-gonic/gin"

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api", h.userIdentity)
	{
		lists := api.Group("/lists")
		{
			lists.POST("/", h.createList)
			lists.GET("/", h.getUserLists)
			lists.GET("/:id", h.getListByID)
			lists.PUT("/:id", h.updateList)
			lists.DELETE("/:id", h.deleteList)
		}
		items := lists.Group(":id/items")
		{
			items.POST("/", h.createItem)
			items.GET("/", h.getItems)
			items.GET("/:item_id", h.getItemByID)
			items.PUT("/:item_id", h.updateItem)
			items.DELETE("/:item_id", h.deleteItem)
		}
	}

	return router
}
