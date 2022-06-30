package repository

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	go_chat "go-chat"
	"testing"
)

func TestChatPostgres_CreateChat(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sql-mock")
	r := NewRepository(sqlxDB)

	type args struct {
		chat go_chat.Chats
	}

	tests := []struct {
		name         string
		input        args
		mockBehavior func()
		want         int
		wantErr      bool
	}{
		{
			name: "Ok",
			input: args{
				chat: go_chat.Chats{
					ChatName: "chat_test",
					Author:   "author_test",
				},
			},
			mockBehavior: func() {
				rows := sqlmock.NewRows([]string{"id"}).AddRow(1)
				mock.ExpectQuery("INSERT INTO chats").
					WithArgs("chat_test", "author_test").WillReturnRows(rows)
			},
			want:    1,
			wantErr: false,
		},
		{
			name: "Wrong Input",
			input: args{
				chat: go_chat.Chats{
					ChatName: "",
					Author:   "author_test",
				},
			},
			mockBehavior: func() {
				rows := sqlmock.NewRows([]string{"id"})
				mock.ExpectQuery("INSERT INTO chats").
					WithArgs("", "author_test").WillReturnRows(rows)
			},
			want:    0,
			wantErr: true,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehavior()

			got, err := r.CreateChat(testCase.input.chat)
			if testCase.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, testCase.want, got)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestChatPostgres_CreateMessage(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sql-mock")
	r := NewRepository(sqlxDB)

	type args struct {
		chatId int
		chat   go_chat.Messages
	}

	tests := []struct {
		name         string
		input        args
		mockBehavior func()
		want         int
		wantErr      bool
	}{
		{
			name: "Ok",
			input: args{
				chatId: 1,
				chat: go_chat.Messages{
					Nickname: "nickname_test",
					Msg:      "msg_test",
				},
			},
			mockBehavior: func() {
				rows := sqlmock.NewRows([]string{"id"}).AddRow(1)
				mock.ExpectQuery("INSERT INTO messages").
					WithArgs(1, "nickname_test", "msg_test").WillReturnRows(rows)
			},
			want:    1,
			wantErr: false,
		},
		{
			name: "Wrong Input",
			input: args{
				chatId: 1,
				chat: go_chat.Messages{
					Nickname: "",
					Msg:      "msg_test",
				},
			},
			mockBehavior: func() {
				rows := sqlmock.NewRows([]string{"id"})
				mock.ExpectQuery("INSERT INTO messages").
					WithArgs(1, "", "msg_test").WillReturnRows(rows)
			},
			want:    0,
			wantErr: true,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehavior()

			got, err := r.CreateMessage(testCase.input.chatId, testCase.input.chat)
			if testCase.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, testCase.want, got)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestChatPostgres_GetListIdMessages(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sql-mock")
	r := NewRepository(sqlxDB)

	type args struct {
		chatId int
		filter go_chat.FilterRequestBody
	}

	tests := []struct {
		name         string
		input        args
		mockBehavior func()
		want         []int
		wantErr      bool
	}{
		{
			name: "Ok",
			input: args{
				chatId: 1,
				filter: go_chat.FilterRequestBody{
					Limit: 3,
				},
			},
			mockBehavior: func() {
				rows := sqlmock.NewRows([]string{"list"}).AddRow(1).AddRow(2).AddRow(3)
				mock.ExpectQuery(
					"SELECT m.id FROM messages m INNER JOIN chats c on m.chat_id = c.id WHERE (.+)").
					WithArgs(1).WillReturnRows(rows)
			},
			want:    []int{1, 2, 3},
			wantErr: false,
		},
		{
			name: "No Records",
			input: args{
				chatId: 404,
				filter: go_chat.FilterRequestBody{
					Limit: 3,
				},
			},
			mockBehavior: func() {
				rows := sqlmock.NewRows([]string{"list"})
				mock.ExpectQuery(
					"SELECT m.id FROM messages m INNER JOIN chats c on m.chat_id = c.id WHERE (.+)").
					WithArgs().WillReturnRows(rows)
			},
			wantErr: false,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehavior()

			got, err := r.GetListIdMessages(testCase.input.chatId, testCase.input.filter)
			if testCase.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, testCase.want, got)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestChatPostgres_GetMessageById(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sql-mock")
	r := NewRepository(sqlxDB)

	type args struct {
		chatId int
		msgId  int
	}

	tests := []struct {
		name         string
		input        args
		mockBehavior func()
		want         go_chat.Messages
		wantErr      bool
	}{
		{
			name: "Ok",
			input: args{
				chatId: 1,
				msgId:  1,
			},
			mockBehavior: func() {
				rows := sqlmock.NewRows([]string{"id", "chat_id", "nickname", "msg"}).
					AddRow(1, 1, "nick_t", "msg_t")
				mock.ExpectQuery(
					"SELECT nickname, msg FROM messages WHERE (.+)").
					WithArgs(1, 1).WillReturnRows(rows)
			},
			want: go_chat.Messages{
				Id:       1,
				ChatId:   1,
				Nickname: "nick_t",
				Msg:      "msg_t",
			},
			wantErr: false,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehavior()

			got, err := r.GetMessageById(testCase.input.chatId, testCase.input.msgId)
			if testCase.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, testCase.want, got)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}
