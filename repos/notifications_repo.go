package repos

import (
	"database/sql"
	"microservice_clean_design/app/core"
	"microservice_clean_design/domain"
)

type NotificationDBRepo struct {
	log core.Logger
	db  *sql.DB
}

func NewNotificationsDBRepo(log core.Logger, db *sql.DB) *NotificationDBRepo {
	return &NotificationDBRepo{
		log: log,
		db:  db,
	}
}

func (r *NotificationDBRepo) All() ([]*domain.Notification, error) {
	var notifications []*domain.Notification

	query := "SELECT id, title, created_at FROM notifications ORDER BY id;"
	raws, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	for raws.Next() {
		notification := &domain.Notification{}
		err = raws.Scan(&notification.Id, &notification.Title, &notification.CreatedAt)
		if err != nil {
			return nil, err
		}
		notifications = append(notifications, notification)
	}
	return notifications, nil
}

func (r *NotificationDBRepo) FetchById(id int32) (*domain.Notification, error) {
	notification := &domain.Notification{
		Id: id,
	}
	query := "SELECT id, title, created_at FROM notifications WHERE id=$1"
	err := r.db.QueryRow(query, id).Scan(&notification.Id, &notification.Title, &notification.CreatedAt)
	if err != nil {
		return notification, err
	}
	return notification, nil

}

func (r *NotificationDBRepo) Insert(notification *domain.Notification) error {
	var id int32
	query := "INSERT INTO notifications (title) VALUES ($1) returning id"
	err := r.db.QueryRow(query, notification.Title).Scan(&id)
	if err != nil {
		return err
	}
	notification.Id = id
	return nil
}

func (r *NotificationDBRepo) Update(notification *domain.Notification) error {
	query := "UPDATE notifications SET title=$2, updated_at=now() WHERE id=$1"
	_, err := r.db.Exec(query, notification.Id, notification.Title)
	if err != nil {
		return err
	}
	return nil
}
