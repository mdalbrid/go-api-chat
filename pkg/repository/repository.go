package repository

import (
	"github.com/jmoiron/sqlx"
	go_chat "go-chat"
)

type Chat interface {
	CreateChat(chat go_chat.Chats) (int, error)
	CreateMessage(chatId int, chat go_chat.Messages) (int, error)
	GetListIdMessages(chatId int, filter go_chat.FilterRequestBody) ([]int, error)
	GetMessageById(chatId int, msgId int) (go_chat.Messages, error)
}

type Repository struct {
	Chat
}

// NewRepository ...
func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Chat: NewChatPostgres(db),
	}
}
