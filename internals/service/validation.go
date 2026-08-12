package service

import (
	"context"
	"errors"
	"net/mail"

	"github.com/manasss0508/notifyx-go/internals/datamodels"
	"go.opentelemetry.io/otel/trace"
)

var (
	ErrInvalidEmail    = errors.New("invalid email")
	ErrInvalidChannel  = errors.New("invalid channel")
	ErrInvalidTemplate = errors.New("invalid template")
)

// Validation : it validate create notification req.body
func Validation(ctx context.Context, otelTracer trace.Tracer, payload *datamodels.CreateNotificationRequest) error {
	// otel
	_, span := otelTracer.Start(ctx, " validating notification")
	defer span.End()

	// channel and recipient validation
	if err := validateChannelAndRecipient((*payload).Channel, (*payload).Recipient); err != nil {
		return err
	}

	// template validation
	if err := validateTemplate((*payload).Template); err != nil {
		return err
	}

	return nil
}

func validateChannelAndRecipient(ch, recipient string) error {
	switch ch {
	case "MAIL":
		return validateMail(recipient)
	default:
		return ErrInvalidChannel
	}
}

func validateMail(email string) error {
	_, err := mail.ParseAddress(email)

	if err != nil {
		return ErrInvalidEmail
	}

	return nil
}

func validateTemplate(temp string) error {

	switch temp {
	case "WELCOME":
		return nil
	default:
		return ErrInvalidTemplate
	}

}
