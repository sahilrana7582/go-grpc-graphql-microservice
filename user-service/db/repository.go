package db

import (
	"database/sql"
	"errors"

	"github.com/google/uuid"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) CreateUser(user *User) error {
	query := `INSERT INTO users (id, username, password, role) VALUES ($1, $2, $3, $4)`
	_, err := r.DB.Exec(query, user.ID, user.Username, user.Password, user.Role)
	return err
}

func (r *UserRepository) GetUserByID(id string) (*User, error) {
	userID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	var user User
	query := `SELECT id, username, password, role FROM users WHERE id = $1`
	row := r.DB.QueryRow(query, userID)
	err = row.Scan(&user.ID, &user.Username, &user.Password, &user.Role)
	if err == sql.ErrNoRows {
		return nil, errors.New("user not found")
	} else if err != nil {
		return nil, err
	}
	return &user, nil
}
