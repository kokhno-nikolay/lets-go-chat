package models

type Message struct {
	ID       int    `json:"id"`
	Text     string `json:"text"`
	UserUUID string `json:"user_uuid"`
}
