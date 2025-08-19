package models

import (
	"errors"
	"rest-api/m/db"
	"rest-api/m/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding: "required"`
}

func (u User) Save() error {
	query := "INSERT INTO users(email, password) VALUES(?, ?)"

	hashedPassword, err := utils.HashPassword(u.Password)

	if err != nil {
		return err
	}

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()
	result, err := stmt.Exec(u.Email, hashedPassword)

	if err != nil {
		return err
	}

	// in case of success
	userId, err := result.LastInsertId()

	u.ID = userId
	return err
}

// login validation logic

func (u *User) ValidateCredentials() error {
	query := "SELECT id, password FROM users WHERE email= ?"
	row := db.DB.QueryRow(query, u.Email)

	var retrivedPassword string
	err := row.Scan(u.ID, &retrivedPassword)

	if err != nil {
		return errors.New("Credentials Invalid")
	}

	passwordIsValid := utils.CheckHashedPassword(u.Password, retrivedPassword)

	if !passwordIsValid {
		return errors.New("Credentials Invalid!")
	}

	return nil
}
