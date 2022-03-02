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

func (r *UsersRepo) GetAllActive() ([]models.User, error) {
	activeUsers := []models.User{}

	rows, err := r.db.Query("SELECT uuid, username, password, active FROM users WHERE active = true")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		activeUser := models.User{}
		if err := rows.Scan(&activeUser.UUID, &activeUser.Username, &activeUser.Password, &activeUser.Active); err != nil {
			return nil, err
		}
		activeUsers = append(activeUsers, activeUser)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return activeUsers, err
}

func (r *UsersRepo) SwitchToActive(userID int) error {
	_, err := r.db.Exec("UPDATE users SET active = true WHERE id = $1", userID)
	return err
}

func (r *UsersRepo) SwitchToInactive(userID int) error {
	_, err := r.db.Exec("UPDATE users SET active = false WHERE id = $1", userID)
	return err
}
