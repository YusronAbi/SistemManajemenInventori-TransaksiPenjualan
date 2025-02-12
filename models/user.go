package models

import "time"

type User struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserRepository interface {
	GetByID(id int64) (*User, error)
	GetByEmail(email string) (*User, error)
	Store(user *User) error
	Update(user *User) error
	Delete(id int64) error
}

type UserUsecase interface {
	GetByID(id int64) (*User, error)
	Register(user *User) error
	Login(email, password string) (string, error)
	Update(user *User) error
}
