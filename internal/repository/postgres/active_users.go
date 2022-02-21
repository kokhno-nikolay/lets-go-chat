package postgres

import "database/sql"

type ActiveUsersRepo struct {
	db *sql.DB
}

func NewActiveUsersRepo(db *sql.DB) *ActiveUsersRepo {
	return &ActiveUsersRepo{
		db: db,
	}
}

func (r *ActiveUsersRepo) Count() (int, error) {
	return 0, nil
}

func (r *ActiveUsersRepo) Add() error {
	return nil
}

func (r *ActiveUsersRepo) Delete() error {
	return nil
}
