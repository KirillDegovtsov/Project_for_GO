package domain

import "time"

type Note struct {
	Title       string    `json:"title"`
	Text        string    `json:"text"`
	UserId      string    `json:"user_id"`
	CreatedTime time.Time `json:"created_time"`
	LastChange  time.Time `json:"last_change"`
}

type User struct {
	Id       string `json:"user_id"`
	Login    string `json:"login"`
	Password string `json:"password"`
}
