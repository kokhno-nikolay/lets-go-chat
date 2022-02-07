package postgres

import (
	"database/sql"

	"github.com/google/uuid"

	"github.com/kokhno-nikolay/lets-go-chat/internal/models"
)

type UsersRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UsersRepo {
	return &UsersRepo{
		db: db,
	}
}

func (r *UsersRepo) GetUUID(username string) (string, error) {
	user := models.User{}

	row := r.db.QueryRow("SELECT uuid FROM users WHERE username = $1", username)
	err := row.Scan(&user.UUID)
	return user.UUID, err
}

func (r *UsersRepo) Create(user models.User) (string, error) {
	uuid := uuid.New().String()
	_, err := r.db.Exec(`INSERT INTO users ("uuid", "username", "password") VALUES ($1, $2, $3)`, uuid, user.Username, user.Password)
	return uuid, err
}
