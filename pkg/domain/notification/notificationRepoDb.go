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
	err := r.db.QueryRow(`INSERT INTO notifications(title, body, user_id) VALUES ($1,$2,$3)RETURNING id;`, notificationObj.Title, notificationObj.Body, notificationObj.UserUID).Scan(&notificationObj.ID)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (r NotificationRepositoryDb) ReadByNotificationID(id int) (Notification,error){
	var notiObj Notification
	sqlStatement := `SELECT id,title,body,user_id,created_at,updated_at,deleted_at FROM notifications WHERE id=$1;`
	row := r.db.QueryRow(sqlStatement, id)
	err := row.Scan(&notiObj.ID,&notiObj.Title,&notiObj.Body,&notiObj.UserUID,&notiObj.CreatedAt,&notiObj.UpdatedAt,&notiObj.DeletedAt) 
	//case sql.ErrNoRows:
  	//	return notiObj,errors.New("notification not found")
	//case nil:
  	//	return notiObj,nil
	//default:
  	return notiObj,err
	
}


func NewNotificationRepositoryDb(db *sql.DB) NotificationRepositoryDb {
	return NotificationRepositoryDb{db}
}