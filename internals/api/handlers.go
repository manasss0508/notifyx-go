package api

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/google/uuid"
	"github.com/manasss0508/notifyx-go/internals/datamodels"
	"github.com/manasss0508/notifyx-go/internals/repository"
	"github.com/manasss0508/notifyx-go/internals/service"
)

func (state AppState) createNotificationHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	// otel span for tracing
	ctx := r.Context()
	ctx, span := state.OtelTracer.Start(ctx, "creating notification") // parents span of handle
	defer span.End()
	slog.InfoContext(ctx, "Creating Notification")

	// deserializing into r.body into struct
	decoder := json.NewDecoder(r.Body) // reading body into memory
	var req datamodels.CreateNotificationRequest
	req.Priority = "LOW" // setting default value of this field

	if err := decoder.Decode(&req); err != nil { // decoding r.body and handling error
		sendResponse(w, http.StatusInternalServerError,
			datamodels.ErrorNotification{
				Message: "internal server error",
			}) // sending an error response
		return
	}

	// validation
	if err := service.Validation(ctx, state.OtelTracer, &req); err != nil {
		sendResponse(w, http.StatusBadRequest,
			datamodels.ErrorNotification{
				Message: err.Error(),
			},
		) // sending an error response
		return
	}
	slog.InfoContext(ctx, "Notification validation success")

	//creating notification id
	notificationID := uuid.New()
	//notificationIDString := notificationID.String()
	slog.InfoContext(ctx, "Notification id generated")

	// save to database
	createdNotif, err := repository.DbCreateNotifcation(ctx, state.OtelTracer, state.Db, notificationID, req, "PENDING")
	if err != nil {
		sendResponse(w, http.StatusInternalServerError,
			datamodels.ErrorNotification{
				Message: err.Error(),
			},
		) // sending an error response
		return
	}
	slog.InfoContext(ctx, "Notification saved to database")

	//publish to queue
	err = state.Rbmq.Publish(ctx, state.OtelTracer, createdNotif.ID, createdNotif.Channel)
	if err != nil {
		sendResponse(w, http.StatusInternalServerError,
			datamodels.ErrorNotification{
				Message: err.Error(),
			},
		) // sending an error response
		return
	}

	// send response, notification is created successfully
	sendResponse(w, http.StatusCreated,
		datamodels.CreateNotificationResponse{
			NotificationId: createdNotif.ID,
			Status:         createdNotif.Status,
		},
	)
}

func (state AppState) getNotificationHandler(w http.ResponseWriter, r *http.Request) {
	// otel span for tracing
	ctx := r.Context()
	ctx, span := state.OtelTracer.Start(ctx, "creating notification") // parents span of handle
	defer span.End()

	// parsing uuid from string
	uuid, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		sendResponse(w, http.StatusInternalServerError,
			datamodels.ErrorNotification{
				Message: err.Error(),
			},
		)
		return
	}

	// get notification from db
	notif, err := repository.DbGetNotification(ctx, state.OtelTracer, state.Db, uuid)

	// if notification not exits
	if err != nil {
		sendResponse(w, http.StatusInternalServerError,
			datamodels.GetNotificationResponse{
				Message: "notification does not exits",
			},
		)
		return
	}

	// notification exits
	sendResponse(w, http.StatusOK,
		datamodels.GetNotificationResponse{
			Notification: notif,
			Message:      "",
		},
	)

	return
}

func sendResponse[T any](w http.ResponseWriter, sc int, response T) {

	// response content
	w.Header().Set("Content-Type", "application/json")

	// response status code
	w.WriteHeader(sc)

	// marshaling and sending response using encoder
	if err := json.NewEncoder(w).Encode(response); err != nil {
		fmt.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError) //
	}

}
