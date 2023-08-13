package model

import (
	"context"
	"time"

	"github.com/labstack/echo/v4"
)

type CreateRequest struct {
	Message string `json:"message" validate:"required"`
}

func (c *CreateRequest) Validate() error {
	return validate.Struct(c)
}

type Notification struct {
	Id        string    `json:"id"`
	Channel   string    `json:"channel"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
}

type NotificationRepository interface {
	Save(ctx context.Context, notif *Notification) error
}

type NotificationSubscribe interface {
	ReceiveNotif()
	Close()
}

type NotificationService interface {
	Create(ctx context.Context, req CreateRequest) (bool, error)
}

type NotificationController interface {
	HandleCreate() echo.HandlerFunc
}
