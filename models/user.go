package models

import (
	"errors"

	"github.com/event-backend-api/db"
	"github.com/event-backend-api/utils"
)

type User struct {
	ID       int64  `json:"id"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (u *User) Save() error {
	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}
	u.Password = hashedPassword

	result, err := db.DB.Exec(
		`INSERT INTO users (email, password) VALUES (?, ?)`,
		u.Email,
		u.Password,
	)
	if err != nil {
		return err
	}

	u.ID, err = result.LastInsertId()
	return err
}

func (u *User) ValidateCredentials() error {
	var query = `SELECT password FROM users WHERE email = ?`
	row := db.DB.QueryRow(query, u.Email)
	var storedHashedPassword string
	err := row.Scan(&storedHashedPassword)

	if err != nil {
		return errors.New("invalid credentials")
	}
	isPasswordValid := utils.CheckPasswordHash(u.Password, storedHashedPassword)
	if !isPasswordValid {
		return errors.New("invalid credentials")
	}
	return nil

}
