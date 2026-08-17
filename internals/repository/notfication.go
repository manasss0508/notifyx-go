package repository

import (
	"context"
	"encoding/json"

	"github.com/google/uuid"
	"github.com/manasss0508/notifyx-go/internals/datamodels"
	"github.com/manasss0508/notifyx-go/internals/repository/queries"
	"go.opentelemetry.io/otel/trace"
)

func DbCreateNotifcation(
	ctx context.Context,
	otelTracer trace.Tracer,
	q *queries.Queries,
	notifId uuid.UUID,
	notif datamodels.CreateNotificationRequest, status string) (queries.Notification, error) {

	// otel
	_, span := otelTracer.Start(ctx, "saving notification in database")
	defer span.End()

	//varibles
	bytes, err := json.Marshal(notif.Variables)
	if err != nil {
		return queries.Notification{}, err
	}
	variblesJsonString := string(bytes)

	// preparing arguments
	args := queries.DbCreateNotificationParams{
		ID:        notifId,
		Channel:   notif.Channel,
		Recipient: notif.Recipient,
		Template:  notif.Template,
		Variables: variblesJsonString,
		Priority:  notif.Priority,
		Status:    status,
	}

	// executing query
	return q.DbCreateNotification(context.Background(), args)
}
func DbGetNotification(
	ctx context.Context,
	otelTracer trace.Tracer,
	q *queries.Queries,
	notifId uuid.UUID) (queries.Notification, error) {

	// otel
	_, span := otelTracer.Start(ctx, "getting notification from database")
	defer span.End()

	// executing query
	return q.DbGetNotificationById(context.Background(), notifId)

}
