// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	go_chat "go-chat"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockChat is a mock of Chat interface.
type MockChat struct {
	ctrl     *gomock.Controller
	recorder *MockChatMockRecorder
}

// MockChatMockRecorder is the mock recorder for MockChat.
type MockChatMockRecorder struct {
	mock *MockChat
}

// NewMockChat creates a new mock instance.
func NewMockChat(ctrl *gomock.Controller) *MockChat {
	mock := &MockChat{ctrl: ctrl}
	mock.recorder = &MockChatMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockChat) EXPECT() *MockChatMockRecorder {
	return m.recorder
}

// CreateChat mocks base method.
func (m *MockChat) CreateChat(chat go_chat.Chats) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateChat", chat)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateChat indicates an expected call of CreateChat.
func (mr *MockChatMockRecorder) CreateChat(chat interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateChat", reflect.TypeOf((*MockChat)(nil).CreateChat), chat)
}

// CreateMessage mocks base method.
func (m *MockChat) CreateMessage(chatId int, chat go_chat.Messages) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateMessage", chatId, chat)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateMessage indicates an expected call of CreateMessage.
func (mr *MockChatMockRecorder) CreateMessage(chatId, chat interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateMessage", reflect.TypeOf((*MockChat)(nil).CreateMessage), chatId, chat)
}

// GetListIdMessages mocks base method.
func (m *MockChat) GetListIdMessages(chatId int, filter go_chat.FilterRequestBody) ([]int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetListIdMessages", chatId, filter)
	ret0, _ := ret[0].([]int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetListIdMessages indicates an expected call of GetListIdMessages.
func (mr *MockChatMockRecorder) GetListIdMessages(chatId, filter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetListIdMessages", reflect.TypeOf((*MockChat)(nil).GetListIdMessages), chatId, filter)
}

// GetMessageById mocks base method.
func (m *MockChat) GetMessageById(chatId, msgId int) (go_chat.Messages, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMessageById", chatId, msgId)
	ret0, _ := ret[0].(go_chat.Messages)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMessageById indicates an expected call of GetMessageById.
func (mr *MockChatMockRecorder) GetMessageById(chatId, msgId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMessageById", reflect.TypeOf((*MockChat)(nil).GetMessageById), chatId, msgId)
}
