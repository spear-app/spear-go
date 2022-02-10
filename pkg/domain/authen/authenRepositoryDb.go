package authen

import (
	"database/sql"
	"fmt"

	"github.com/spear-app/spear-go/pkg/domain/user"
)

type AuthenRepositoryDb struct {
	db *sql.DB
}

func (r AuthenRepositoryDb) Signup(user *user.User) error {
	err := r.db.QueryRow(`INSERT INTO users(email, password, name, gender) VALUES ($1,$2,$3,$4)RETURNING id;`, user.Email, user.Password, user.Name, user.Gender).Scan(&user.ID)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
func (r AuthenRepositoryDb) Login(user *user.User) error {
	row := r.db.QueryRow(`SELECT id, name, email, password, gender FROM users WHERE email=$1`, user.Email)
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Gender)
	if err != nil {
		return err
	}
	return nil
}

func NewAuthenRepositoryDb(db *sql.DB) AuthenRepositoryDb {
	return AuthenRepositoryDb{db}
}
