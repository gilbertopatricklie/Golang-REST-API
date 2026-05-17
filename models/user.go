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
	Role     string
}
 
func (u User) Save() error {
	
	if u.Role == "" {
		u.Role = "user"
	}
 
	query := "INSERT INTO users(email, password, role) VALUES(?, ?, ?)"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
 
	defer stmt.Close()
 
	hashedPass, err := utils.HashPassword(u.Password)
	if err != nil{
		return err
	}
 
	result, err := stmt.Exec(u.Email, hashedPass, u.Role)
	if err != nil{
		return err
	}
 
	userId, err := result.LastInsertId()
	u.ID = userId
	return err
}
 
func (u *User) Validate() error {
	
	query := "SELECT id, password, role FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, u.Email)
 
	var retrievedPass string
	err := row.Scan(&u.ID, &retrievedPass, &u.Role)
	if err != nil{
		return err
	}
 
	isValPassword := utils.CheckPassword(u.Password, retrievedPass)
	if !isValPassword {
		return errors.New("invalid credentials")
	}
 
	return nil
}