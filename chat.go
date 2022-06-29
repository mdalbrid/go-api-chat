package go_chat

type Chats struct {
	Id       int    `json:"id" db:"id"`
	ChatName string `json:"chat_name" binding:"required"`
	Author   string `json:"author" binding:"required"`
}

type Messages struct {
	Id       int    `json:"id" db:"id"`
	ChatId   int    `json:"chat_id" db:"chat_id"`
	Nickname string `json:"nickname" db:"nickname" binding:"required"`
	Msg      string `json:"msg" db:"msg" binding:"required"`
}

type Users struct {
	Id       int    `json:"-"`
	Nickname string `json:"nickname" binding:"required"`
}
