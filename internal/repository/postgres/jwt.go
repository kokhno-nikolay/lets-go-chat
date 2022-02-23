package postgres

import "database/sql"

type JWTRepo struct {
	db *sql.DB
}

func NewJWTRepo(db *sql.DB) *JWTRepo {
	return &JWTRepo{
		db: db,
	}
}
