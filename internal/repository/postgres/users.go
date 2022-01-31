package postgres

import (
	"database/sql"
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

func (r *UsersRepo) Create(user models.User) error {

	_, err := r.db.Exec(`INSERT INTO users ("username", "password") VALUES ($1, $2)`, user.Username, user.Password)
	return err
}
