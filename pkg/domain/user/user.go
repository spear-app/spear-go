package user

import (
	"time"
)

//user model

type User struct {
	ID            uint       `json:"id"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
	DeletedAt     *time.Time `json:"deleted_at"`
	Name          string     `json:"name"`
	Email         string     `json:"email,omitempty" validate:"required,email"`
	Password      string     `json:"password,omitempty" validate:"required,gte=8"`
	OTP           string     `json:"OTP,omitempty"`
	EmailVerified bool       `json:"email_verified"`
	Gender        string     `json:"gender"`
}

//interface that contains functions of CRUD operations
type UserRepository interface {
	Index() ([]User, error)
	Create(User) error
	Update(User, string) error
	Delete(string) error
	ReadUserInfo(string) (User, error)
}
