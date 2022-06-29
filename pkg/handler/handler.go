package handler

import (
	"github.com/gin-gonic/gin"
	"go-chat/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

// InitRoutes ...
func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		chat := api.Group("/chat")
		{
			chat.POST("/", h.createChatHandler)

			msg := chat.Group(":chat_id/msg")
			{
				msg.POST("/", h.createMessageByIdChatHandler)
				msg.GET("/", h.getListIdMessagesHandler)
				msg.GET("/:msg_id", h.getMessageByIdHandler)
			}
		}
	}

	return router
}
