package postgres

import (
	"database/sql"

	"github.com/kokhno-nikolay/lets-go-chat/internal/models"
)

type MessageRepo struct {
	db *sql.DB
}

func NewMessageRepo(db *sql.DB) *MessageRepo {
	return &MessageRepo{
		db: db,
	}
}

func (r *MessageRepo) GetAll() ([]models.Message, error) {
	messages := []models.Message{}
	rows, err := r.db.Query("SELECT id, text, user_uuid FROM messages")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		message := models.Message{}
		if err := rows.Scan(&message.ID, &message.Text, &message.UserID); err != nil {
			return nil, err
		}
		messages = append(messages, message)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return messages, nil
}

func (r *MessageRepo) Create(message models.Message) error {
	_, err := r.db.Exec(`INSERT INTO messages ("id", "text", "user_id") VALUES ($1, $2, $3)`, message.ID, message.Text, message.UserID)
	return err
}
