package models

type Message struct {
	ID     int    `json:"id"`
	Text   string `json:"text"`
	UserID int    `json:"user_id"`
}
