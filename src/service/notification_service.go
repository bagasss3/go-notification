package service

import (
	"context"
	"encoding/json"
	"go-notif/src/constant"
	"go-notif/src/helper"
	"go-notif/src/model"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

type notificationService struct {
	redis *redis.Client
}

func NewNotificationService(redis *redis.Client) model.NotificationService {
	return &notificationService{
		redis: redis,
	}
}

func (n *notificationService) Create(ctx context.Context, req model.CreateRequest) (bool, error) {
	log := logrus.WithFields(logrus.Fields{
		"msg":     "service",
		"request": req,
	})

	if err := req.Validate(); err != nil {
		log.Error(err)
		return false, constant.HttpValidationOrInternalErr(err)
	}

	notif := &model.Notification{
		Id:        helper.GenerateID(),
		Channel:   "notification",
		Message:   req.Message,
		CreatedAt: time.Now(),
	}

	jsonData, err := json.Marshal(notif)
	if err != nil {
		log.Error(err)
		return false, err
	}

	if err := n.redis.Publish(ctx, notif.Channel, jsonData).Err(); err != nil {
		log.Error(err)
		return false, err
	}

	return true, nil
}
