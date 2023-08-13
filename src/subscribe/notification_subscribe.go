package subscribe

import (
	"context"
	"encoding/json"
	"fmt"
	"go-notif/src/model"

	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

type notificationSubscribe struct {
	redis                  *redis.Client
	notificationRepository model.NotificationRepository
	subscription           *redis.PubSub
}

func NewNotificationSubscribe(redis *redis.Client, notificationRepository model.NotificationRepository) model.NotificationSubscribe {
	return &notificationSubscribe{
		redis:                  redis,
		notificationRepository: notificationRepository,
	}
}

func (n *notificationSubscribe) ReceiveNotif() {
	var ctx = context.Background()
	log := logrus.WithFields(logrus.Fields{
		"msg": "subscribe",
	})
	subscriber := n.redis.Subscribe(ctx, "notification")
	defer subscriber.Close()

	ch := subscriber.Channel()
	for msg := range ch {
		log.Info(fmt.Sprintf("Received message from channel notification with msg: %s", msg.Payload))
		var notif *model.Notification
		if err := json.Unmarshal([]byte(msg.Payload), &notif); err != nil {
			log.Error(err)
			continue
		}
		err := n.notificationRepository.Save(ctx, notif)
		if err != nil {
			log.Error(err)
		}
	}

}

func (n *notificationSubscribe) Close() {
	if n.subscription != nil {
		if err := n.subscription.Close(); err != nil {
			logrus.Errorf("Error closing Redis subscription: %v", err)
		}
	}
}
