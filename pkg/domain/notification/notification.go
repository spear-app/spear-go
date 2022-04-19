package notification

import "time"

type Notification struct{
	ID				uint 		`json:"id"`
	UserUID			uint  		`json:"user_id"`
	CreatedAt		time.Time  	`json:"created_at"`
	UpdatedAt		time.Time 	`json:"updated_at"`
	DeletedAt		*time.Time	`json:"deleted_at"`
	Title			string		`json:"title"`
	Body			string		`json:"body"`
}

type NotificationRepository interface {
	Create(*Notification) error
}