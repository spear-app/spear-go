package notification

import "time"
type Notification struct{
	ID				int 		`json:"id"` 
	UserUID			int  		`json:"user_id" validate:"required"`
	CreatedAt		time.Time  	`json:"created_at"`
	UpdatedAt		time.Time 	`json:"updated_at"`
	DeletedAt		*time.Time	`json:"deleted_at"`
	Title			string		`json:"title" validate:"required"`
	Body			string		`json:"body" validate:"required"`
}

type NotificationRepository interface {
	Create(*Notification) error
	ReadByNotificationID(int) (Notification,error)
}