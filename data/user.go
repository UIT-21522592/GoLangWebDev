package data

import (
	"time"
)

type User struct {
	Id int 
	Uuid string
	Name string
	Email string
	Password string	
	CreatedAt time.Time
}

type Session struct {
	Id int
	Uuid string
	Email string
	UserId int
	CreatedAt time.Time
}

func (user *User) CreateSession() (session Session, err error) {
	statement := "INSERT INTO session (uuid,email,user_id,created_at) values ($1,$2,$3,$4) RETURNING id,email,user_id,created_at"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}

	defer stmt.Close()
	// use QueryRow to return a row and scan the returned id into the session struct
	err = stmt.QueryRow()
}