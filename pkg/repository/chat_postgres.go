package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	go_chat "go-chat"
)

type ChatPostgres struct {
	db *sqlx.DB
}

func NewChatPostgres(db *sqlx.DB) *ChatPostgres {
	return &ChatPostgres{db: db}
}

// CreateChat ...
func (r *ChatPostgres) CreateChat(chat go_chat.Chats) (int, error) {
	//tx, err := r.db.Begin()
	//if err != nil {
	//	return 0, err
	//}

	var id int

	createChatQuery := fmt.Sprintf("INSERT INTO %s (chat_name, author) VALUES ($1, $2) RETURNING id", chatsTable)
	row := r.db.QueryRow(createChatQuery, chat.ChatName, chat.Author)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	//createUserQuery := fmt.Sprintf("INSERT INTO %s nickname VALUES $1", usersTable)
	//tx.Exec(createUserQuery, chat.Author)
	//if err != nil {
	//	tx.Rollback()
	//	return 0, err
	//}

	return id, nil //tx.Commit()
}

// CreateMessage ...
func (r *ChatPostgres) CreateMessage(chatId int, msg go_chat.Messages) (int, error) {
	var id int

	createMessageQuery := fmt.Sprintf(
		"INSERT INTO %s (chat_id, nickname, msg) VALUES ($1, $2, $3) RETURNING id", messagesTable)
	row := r.db.QueryRow(createMessageQuery, chatId, msg.Nickname, msg.Msg)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

// GetListIdMessages ...
func (r *ChatPostgres) GetListIdMessages(chatId int, filter go_chat.FilterRequestBody) ([]int, error) {
	var list []int

	query := fmt.Sprintf(
		"SELECT m.id FROM %s m INNER JOIN %s c on m.chat_id = c.id WHERE chat_id = $1", messagesTable, chatsTable)

	if filter.Limit > 0 {
		query = fmt.Sprintf("%s LIMIT %d", query, filter.Limit)
	}
	if err := r.db.Select(&list, query, chatId); err != nil {
		return list, err
	}

	return list, nil
}

// GetMessageById ...
func (r *ChatPostgres) GetMessageById(chatId int, msgId int) (go_chat.Messages, error) {
	var data go_chat.Messages

	query := fmt.Sprintf(
		"SELECT nickname, msg FROM %s WHERE chat_id = $1 and id = $2", messagesTable)

	if err := r.db.Get(&data, query, chatId, msgId); err != nil {
		return data, err
	}

	return data, nil
}
