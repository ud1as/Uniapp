package model

// struct for auth
type User struct {
	Id       int    `json:"id" db:"id" binding:"required"`
	Name     string `json:"name" db:"name" binding:"required"`
	Username string `json:"username" db:"username" binding:"required"`
	Password string `json:"password" db:"password" binding:"required"`
}
