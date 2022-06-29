package service

//go:generate mockgen -source=service.go -destination=mocks/mock.go

import (
	go_chat "go-chat"
	"go-chat/pkg/repository"
)

type Chat interface {
	CreateChat(chat go_chat.Chats) (int, error)
	CreateMessage(chatId int, chat go_chat.Messages) (int, error)
	GetListIdMessages(chatId int, filter go_chat.FilterRequestBody) ([]int, error)
	GetMessageById(chatId int, msgId int) (go_chat.Messages, error)
}

type Service struct {
	Chat
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Chat: NewChatService(repos.Chat),
	}
}
