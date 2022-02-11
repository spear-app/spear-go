package driver

import (
	"database/sql"
	"fmt"
	"github.com/spear-app/spear-go/pkg/domain/user"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func Seed(db *sql.DB) {
	usr1 := user.User{
		ID:            0,
		CreatedAt:     time.Time{},
		UpdatedAt:     time.Time{},
		DeletedAt:     nil,
		Name:          "Addison",
		Email:         "Addison@gmail.com",
		Password:      "password1",
		OTP:           "",
		EmailVerified: false,
		Gender:        "",
	}
	usr2 := user.User{
		ID:            0,
		CreatedAt:     time.Time{},
		UpdatedAt:     time.Time{},
		DeletedAt:     nil,
		Name:          "Micheal",
		Email:         "Micheal@gmail.com",
		Password:      "password2",
		OTP:           "",
		EmailVerified: false,
		Gender:        "MALE",
	}
	usr3 := user.User{
		ID:            0,
		CreatedAt:     time.Time{},
		UpdatedAt:     time.Time{},
		DeletedAt:     nil,
		Name:          "Shayna",
		Email:         "Shayna@gmail.com",
		Password:      "password3",
		OTP:           "",
		EmailVerified: false,
		Gender:        "FEMALE",
	}
	usr4 := user.User{
		ID:            0,
		CreatedAt:     time.Time{},
		UpdatedAt:     time.Time{},
		DeletedAt:     nil,
		Name:          "Selena",
		Email:         "Selena@gmail.com",
		Password:      "password4",
		OTP:           "",
		EmailVerified: false,
		Gender:        "FEMALE",
	}
	usr5 := user.User{
		ID:            0,
		CreatedAt:     time.Time{},
		UpdatedAt:     time.Time{},
		DeletedAt:     nil,
		Name:          "Smith",
		Email:         "Smith@gmail.com",
		Password:      "password5",
		OTP:           "",
		EmailVerified: false,
		Gender:        "MALE",
	}
	usr6 := user.User{
		ID:            0,
		CreatedAt:     time.Time{},
		UpdatedAt:     time.Time{},
		DeletedAt:     nil,
		Name:          "Anisa",
		Email:         "Anisa@gmail.com",
		Password:      "password6",
		OTP:           "",
		EmailVerified: false,
		Gender:        "FEMALE",
	}
	usr7 := user.User{
		ID:            0,
		CreatedAt:     time.Time{},
		UpdatedAt:     time.Time{},
		DeletedAt:     nil,
		Name:          "Moody",
		Email:         "Moody@gmail.com",
		Password:      "password7",
		OTP:           "",
		EmailVerified: false,
		Gender:        "MALE",
	}
	usr8 := user.User{
		ID:            0,
		CreatedAt:     time.Time{},
		UpdatedAt:     time.Time{},
		DeletedAt:     nil,
		Name:          "Elis",
		Email:         "Elis@gmail.com",
		Password:      "password8",
		OTP:           "",
		EmailVerified: false,
		Gender:        "FEMALE",
	}
	usr9 := user.User{
		ID:            0,
		CreatedAt:     time.Time{},
		UpdatedAt:     time.Time{},
		DeletedAt:     nil,
		Name:          "Jane",
		Email:         "Jane@gmail.com",
		Password:      "password9",
		OTP:           "",
		EmailVerified: false,
		Gender:        "FEMALE",
	}
	usr10 := user.User{
		ID:            0,
		CreatedAt:     time.Time{},
		UpdatedAt:     time.Time{},
		DeletedAt:     nil,
		Name:          "Brayan",
		Email:         "Brayan@gmail.com",
		Password:      "password10",
		OTP:           "",
		EmailVerified: false,
		Gender:        "MALE",
	}
	var arr [10]user.User
	arr[0] = usr1
	arr[1] = usr2
	arr[2] = usr3
	arr[3] = usr4
	arr[4] = usr5
	arr[5] = usr6
	arr[6] = usr7
	arr[7] = usr8
	arr[8] = usr9
	arr[9] = usr10
	for i := 0; i < 10; i++ {
		hash, err := bcrypt.GenerateFromPassword([]byte(arr[i].Password), 10)
		err = db.QueryRow(`INSERT INTO users(email, password, name, gender) VALUES ($1,$2,$3,$4)RETURNING id;`, arr[i].Email, hash, arr[i].Name, arr[i].Gender).Scan(&arr[i].ID)
		if err != nil {
			fmt.Println(err)
		}
	}
}
