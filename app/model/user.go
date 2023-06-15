package model

import "app/server"

type User struct {
	Id      int
	Name    string
	Surname string
}

func (u *User) Add() (err error) {
	query := `INSERT INTO users (name, surname) VALUES ($1, $2)`
	_, err = server.Db.Exec(query, u.Name, u.Surname)
	return
}

func (u *User) Delete() (err error) {
	query := `DELETE FROM users WHERE id = $1`
	_, err = server.Db.Exec(query, u.Id)
	return
}

func (u *User) Update() (err error) {
	query := `UPDATE users SET name = $1, surname = $2 WHERE id = $3`
	_, err = server.Db.Exec(query, u.Name, u.Surname, u.Id)
	return
}

func NewUser(name, surname string) *User {
	return &User{Name: name, Surname: surname}
}

func GetUserById(id int) (u User, err error) {
	query := `SELECT * FROM users WHERE id = $1`
	rows, err := server.Db.Query(query, id)
	if err != nil {
		return
	}
	rows.Next()
	err = rows.Scan(&u.Id, &u.Name, &u.Surname)
	return
}

func GetAllUsers() ([]User, error) {
	var users []User
	query := `SELECT * FROM users`
	rows, err := server.Db.Query(query)
	if err != nil {
		return users, err
	}
	defer rows.Close()

	for rows.Next() {
		res := User{}
		err := rows.Scan(&res.Id, &res.Name, &res.Surname)
		if err != nil {
			return users, err
		}
		users = append(users, res)
	}
	return users, err
}
