package service

import (
	"context"
	"encoding/json"
	"errors"
	"net/mail"

	"github.com/manasss0508/notifyx-go/internals/datamodels"
	"github.com/manasss0508/notifyx-go/internals/template_engine"
	"go.opentelemetry.io/otel/trace"
)

var (
	ErrInvalidEmail                   = errors.New("invalid email")
	ErrInvalidChannel                 = errors.New("invalid channel")
	ErrInvalidTemplate                = errors.New("invalid template")
	ErrInvalidOrInsufficientVariables = errors.New("invalid or insufficient variables")
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

	// variables validation
	if err := validateVariables((*payload).Template, (*payload).Variables); err != nil {
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
	case "welcome",
		"otp",
		"password_reset",
		"email_verification",
		"login_alert",
		"order_confirmation",
		"order_shipped",
		"payment_success",
		"payment_failed",
		"subscription_renewal":
		return nil

	default:
		return ErrInvalidTemplate
	}

}

func validateVariables(templateName string, jsonObjectMap map[string]string) error {
	bytes, err := json.Marshal(jsonObjectMap)
	if err != nil {
		return ErrInvalidOrInsufficientVariables
	}

	// try to destructure jsonObject into struct ,if failed variables are insufficient
	variables, err := template_engine.JsonStringToTemplateVariables(templateName, bytes)
	if err != nil {
		return ErrInvalidOrInsufficientVariables
	}

	// checking if all variables are available in struct
	if err := variables.Validate(); err != nil {
		return err
	}

	return nil
}
