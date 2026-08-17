package datamodels

import (
	"time"

	"github.com/google/uuid"
	"github.com/manasss0508/notifyx-go/internals/repository/queries"
)

// =========================== create notification
type CreateNotificationRequest struct {
	Channel    string            `json:"channel"`
	Recipient  string            `json:"recipient"`
	Template   string            `json:"template"`
	Variables  map[string]string `json:"variables"`
	Priority   string            `json:"priority"`
	ScheduleAt time.Time         `json:"schedule_at"`
}

type CreateNotificationResponse struct {
	NotificationId uuid.UUID `json:"notification_id"`
	Status         string    `json:"status"`
}

// =========================== get notification
type GetNotificationResponse struct {
	Notification queries.Notification `json:"notification"`
	Message      string               `json:"message"`
}

type ErrorNotification struct {
	Message string `json:"message"`
}
