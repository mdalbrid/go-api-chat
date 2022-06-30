package handler

import (
	"github.com/gin-gonic/gin"
	go_chat "go-chat"
	"net/http"
	"strconv"
)

func (h *Handler) createChatHandler(c *gin.Context) {
	var input go_chat.Chats

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	id, err := h.services.CreateChat(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) createMessageByIdChatHandler(c *gin.Context) {
	chatId, err := strconv.Atoi(c.Param("chat_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid chat id param")
		return
	}

	var input go_chat.Messages
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.CreateMessage(chatId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type Struct struct {
	Id []go_chat.Messages
}

func (h *Handler) getListIdMessagesHandler(c *gin.Context) {
	chatId, err := strconv.Atoi(c.Param("chat_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid chat id param")
		return
	}

	var input go_chat.FilterRequestBody
	input.Limit, err = strconv.Atoi(c.DefaultQuery("limit", "0"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid Limit param")
		return
	}

	listId, err := h.services.GetListIdMessages(chatId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"list_id": listId,
	})

}

func (h *Handler) getMessageByIdHandler(c *gin.Context) {
	chatId, err := strconv.Atoi(c.Param("chat_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid chat id param")
		return
	}

	msgId, err := strconv.Atoi(c.Param("msg_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid message id param")
		return
	}

	data, err := h.services.GetMessageById(chatId, msgId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"nickname": data.Nickname,
		"msg":      data.Msg,
	})
}
