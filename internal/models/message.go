package models

type ChatMessage struct {
	Username string `json:"username"`
	Text     string `json:"text"`
}

type Message struct {
	ID     int    `json:"id"`
	Text   string `json:"text"`
	UserID int    `json:"user_id"`
}
