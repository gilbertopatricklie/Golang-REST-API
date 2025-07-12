package models

import (
	"errors"

	"example.com/restapi/db"
	"example.com/restapi/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u User) Save() error {
	query := "INSERT INTO users(email, password) VALUES(?, ?)"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	hashedPass, err := utils.HashPassword(u.Password)
	if err != nil{

	}

	result, err := stmt.Exec(u.Email, hashedPass)
	if err != nil{
		return err
	}

	userId, err := result.LastInsertId()
	u.ID = userId
	return err
}

func (u *User) Validate() error {
	query := "SELECT id, password FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, u.Email)

	var retrievedPass string
	err := row.Scan(&u.ID, &retrievedPass)
	if err != nil{
		return err
	}

	isValPassword := utils.CheckPassword(u.Password, retrievedPass)
	if !isValPassword {
		return errors.New("invalid credentials")
	}

	return nil
}