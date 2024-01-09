package db

import (
	"log"
	"time"
)

type User struct {
	Id         int
	Username   string
	Password   string
	Email      string
	CreateTime time.Time
	UpdateTime time.Time
}

func UserById(id int) *User {
	rows, err := db.Query("SELECT id, username, password, email, create_time, update_time FROM user WHERE id = ?", id)
	user := &User{}
	if err != nil {
		return user
	}
	defer rows.Close()
	for rows.Next() {
		var createTime string
		var updateTime string
		err := rows.Scan(&user.Id, &user.Username, &user.Password, &user.Email, &createTime, &updateTime)
		if err != nil {
			log.Fatal(err)
		}
		user.CreateTime, _ = time.Parse("2006-01-02 15:04:05", createTime)
		user.UpdateTime, _ = time.Parse("2006-01-02 15:04:05", updateTime)
		break
	}
	return user
}

func UserByEmail(email string) *User {
	rows, err := db.Query("SELECT id, username, password, email, create_time, update_time FROM user WHERE email = ?", email)
	user := &User{}
	if err != nil {
		return user
	}
	defer rows.Close()
	for rows.Next() {
		var createTime string
		var updateTime string
		err := rows.Scan(&user.Id, &user.Username, &user.Password, &user.Email, &createTime, &updateTime)
		if err != nil {
			log.Fatal(err)
		}
		user.CreateTime, _ = time.Parse("2006-01-02 15:04:05", createTime)
		user.UpdateTime, _ = time.Parse("2006-01-02 15:04:05", updateTime)
		break
	}
	return user
}
