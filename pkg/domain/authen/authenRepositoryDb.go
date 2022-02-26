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

func (r AuthenRepositoryDb) InsertOTP(user user.User) error {
	err := r.db.QueryRow(`INSERT INTO users(otp) VALUES ($1)RETURNING id;`, user.OTP).Scan(&user.ID)
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

// ReadUserByID obviously to read by id
func (r AuthenRepositoryDb) ReadUserByID(user *user.User) error {
	row := r.db.QueryRow(`SELECT id, name, email, gender, created_at, updated_at, deleted_at FROM users WHERE id=$1`,
		user.ID)
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Gender, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)
	if err != nil {
		return err
	}
	return nil
}

// Update only name and gender could be updated
func (r AuthenRepositoryDb) Update(user *user.User) error {
	var name string
	row := r.db.QueryRow(`SELECT name FROM users WHERE id= $1`, user.ID)
	err := row.Scan(&name)
	if err != nil {
		fmt.Println(err)
		return err
	}
	_, err = r.db.Exec(`UPDATE users SET name=$1, gender=$2 WHERE id=$3`,
		user.Name, user.Gender, user.ID)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

//Delete This function to delete a user by id
func (r AuthenRepositoryDb) Delete(id string) error {
	var usrType string
	row := r.db.QueryRow(`SELECT name FROM users WHERE id= $1`, id)
	err := row.Scan(&usrType)
	if err != nil {
		fmt.Println(err)
		return err
	}
	_, err = r.db.Exec(`DELETE FROM users WHERE id=$1`, id)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func NewAuthenRepositoryDb(db *sql.DB) AuthenRepositoryDb {
	return AuthenRepositoryDb{db}
}
