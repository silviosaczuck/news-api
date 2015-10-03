package models

/*
User model
*/
type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string
}

/*
Users Multiple users
*/
type Users []*User