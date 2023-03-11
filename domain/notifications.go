package domain

import "time"

type Notification struct {
	Id        int32     `json:"id"`
	Title     string    `json:"title"`
	Text      string    `json:"text"`
	UserId    int32     `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

type NotificationsRepository interface {
	All() ([]*Notification, error)
	FetchById(int32) (*Notification, error)
	Insert(notification *Notification) error
	Update(notification *Notification) error
}

type NotificationInteractor interface {
	GetAllNotifications(int32) (GetAllNotificationsResponse, error)
	CreateNotification(title string) (CreateNotificationResponse, error)
}

type GetAllNotificationsResponse struct {
	StatusCode    string
	Notifications []*Notification
}

type CreateNotificationResponse struct {
	StatusCode string
	Id         int32
}
