package repository

import (
	"context"
	"go-notif/src/model"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type notificationRepository struct {
	db *gorm.DB
}

func NewNotificationRepository(db *gorm.DB) model.NotificationRepository {
	return &notificationRepository{
		db: db,
	}
}

func (n *notificationRepository) Save(ctx context.Context, notification *model.Notification) error {
	log := logrus.WithFields(logrus.Fields{
		"msg":          "repository",
		"notification": notification,
	})

	err := n.db.WithContext(ctx).Create(notification).Error
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}
