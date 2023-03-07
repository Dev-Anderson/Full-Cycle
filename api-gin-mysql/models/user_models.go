package models

type UserID struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
