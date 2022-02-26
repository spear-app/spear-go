package authen

import (
	"github.com/spear-app/spear-go/pkg/domain/user"
)

type Authen struct {
	User user.User `json:"user"`
	//array contains ids of skills that user want to add them to himself
}
type AuthenRepository interface {
	Signup(user *user.User) error
	Login(user *user.User) error
	ReadUserByID(user *user.User) error
	Update(user *user.User) error
	Delete(id string) error
	InsertOTP(user *user.User) error
	VerifyEmail(user *user.User) error
	ReadOTP(user *user.User) error
}
