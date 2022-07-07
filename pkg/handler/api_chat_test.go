package handler

import (
	"bytes"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	go_chat "go-chat"
	"go-chat/pkg/service"
	mock_service "go-chat/pkg/service/mocks"
	"net/http/httptest"
	"testing"
)

func TestHandler_createChatHandler(t *testing.T) {
	type mockBehavior func(r *mock_service.MockChat, chat go_chat.Chats)

	tests := []struct {
		name                 string
		inputBody            string
		inputChat            go_chat.Chats
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "Ok",
			inputBody: `{"chat_name": "chat_test", "author": "author_test"}`,
			inputChat: go_chat.Chats{
				ChatName: "chat_test",
				Author:   "author_test",
			},
			mockBehavior: func(r *mock_service.MockChat, chat go_chat.Chats) {
				r.EXPECT().CreateChat(chat).Return(1, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"id":1}`,
		},
		{
			name:                 "Wrong Input",
			inputBody:            `{"author": "author_test"}`,
			inputChat:            go_chat.Chats{},
			mockBehavior:         func(r *mock_service.MockChat, chat go_chat.Chats) {},
			expectedStatusCode:   400,
			expectedResponseBody: `{"message":"invalid input body"}`,
		},
		{
			name:      "Service Error",
			inputBody: `{"chat_name": "chat_test", "author": "author_test"}`,
			inputChat: go_chat.Chats{
				ChatName: "chat_test",
				Author:   "author_test",
			},
			mockBehavior: func(r *mock_service.MockChat, chat go_chat.Chats) {
				r.EXPECT().CreateChat(chat).Return(0, errors.New("something went wrong"))
			},
			expectedStatusCode:   500,
			expectedResponseBody: `{"message":"something went wrong"}`,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			// Зависимости
			c := gomock.NewController(t)
			defer c.Finish()

			chat := mock_service.NewMockChat(c)
			testCase.mockBehavior(chat, testCase.inputChat)

			services := &service.Service{Chat: chat}
			handler := NewHandler(services)

			// Эндпоинт
			router := gin.New()
			router.POST("/api", handler.createChatHandler)

			//Создание запроса
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/api", bytes.NewBufferString(testCase.inputBody))

			//Запрос
			router.ServeHTTP(w, req)

			//Сравнить
			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedResponseBody, w.Body.String())

		})

	}
}

