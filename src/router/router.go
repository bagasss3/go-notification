package router

import (
	"go-notif/src/model"

	"github.com/labstack/echo/v4"
)

type route struct {
	group                  *echo.Group
	notificationController model.NotificationController
}

func RouteService(group *echo.Group, notificationController model.NotificationController) {
	rt := &route{
		group:                  group,
		notificationController: notificationController,
	}
	rt.routerInit()
}

func (r *route) routerInit() {
	r.group.POST("/notif", r.notificationController.HandleCreate())
}
