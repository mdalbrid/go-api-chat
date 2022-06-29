package service

import (
	go_chat "go-chat"
	"go-chat/pkg/repository"
)

type ChatService struct {
	repo repository.Chat
}

func NewChatService(repo repository.Chat) *ChatService {
	return &ChatService{repo: repo}
}

// CreateChat ...
func (s *ChatService) CreateChat(chat go_chat.Chats) (int, error) {
	return s.repo.CreateChat(chat)
}

// CreateMessage ...
func (s *ChatService) CreateMessage(chatId int, chat go_chat.Messages) (int, error) {
	return s.repo.CreateMessage(chatId, chat)
}

// GetListIdMessages ...
func (s *ChatService) GetListIdMessages(chatId int, filter go_chat.FilterRequestBody) ([]int, error) {
	return s.repo.GetListIdMessages(chatId, filter)
}

// GetMessageById ...
func (s *ChatService) GetMessageById(chatId int, msgId int) (go_chat.Messages, error) {
	return s.repo.GetMessageById(chatId, msgId)
}
