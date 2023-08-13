package controller

import (
	"go-notif/src/constant"
	"go-notif/src/dto"
	"go-notif/src/model"
	"net/http"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

type notificationController struct {
	notificationService model.NotificationService
}

func NewNotificationController(notificationService model.NotificationService) model.NotificationController {
	return &notificationController{
		notificationService: notificationService,
	}
}

func (n *notificationController) HandleCreate() echo.HandlerFunc {
	return func(c echo.Context) error {
		req := model.CreateRequest{}
		if err := c.Bind(&req); err != nil {
			log.Error(err)
			return constant.ErrInternal
		}

		_, err := n.notificationService.Create(c.Request().Context(), req)
		if err != nil {
			log.Error(err)
			return err
		}

		return c.JSON(http.StatusOK, dto.ResponseSuccess{
			Success: true,
		})
	}
}
