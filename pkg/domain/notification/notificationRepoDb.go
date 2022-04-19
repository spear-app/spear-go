package notification

import (
	"database/sql"
	"fmt"
)

type NotificationRepositoryDb struct {
	db *sql.DB
}

// TODO user not found error
func (r NotificationRepositoryDb) Create(notificationObj *Notification) error{
	err := r.db.QueryRow(`INSERT INTO notifications(title, body, user_id) VALUES ($1,$2,$3)RETURNING id;`, notificationObj.Title, notificationObj.Body, notificationObj.UserUID, user.Gender).Scan(&user.ID)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}