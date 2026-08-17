package template_engine

import (
	"encoding/json"
	"fmt"

	templates "github.com/manasss0508/notifyx-go/internals/datamodels/template_engine"
)

func JsonStringToTemplateVariables(templateName string, jsonBytes []byte) (templates.TemplateVariables, error) {

	switch templateName {
	case "welcome":
		var v templates.WelcomeVariables
		return v, json.Unmarshal(jsonBytes, &v)

	case "otp":
		var v templates.OtpVariables
		return v, json.Unmarshal(jsonBytes, &v)

	case "password_reset":
		var v templates.PasswordResetVariables
		return v, json.Unmarshal(jsonBytes, &v)

	case "email_verification":
		var v templates.EmailVerificationVariables
		return v, json.Unmarshal(jsonBytes, &v)

	case "login_alert":
		var v templates.LoginAlertVariables
		return v, json.Unmarshal(jsonBytes, &v)

	case "order_confirmation":
		var v templates.OrderConfirmationVariables
		return v, json.Unmarshal(jsonBytes, &v)

	case "order_shipped":
		var v templates.OrderShippedVariables
		return v, json.Unmarshal(jsonBytes, &v)

	case "payment_success":
		var v templates.PaymentSuccessVariables
		return v, json.Unmarshal(jsonBytes, &v)

	case "payment_failed":
		var v templates.PaymentFailedVariables
		return v, json.Unmarshal(jsonBytes, &v)

	case "subscription_renewal":
		var v templates.SubscriptionRenewalVariables
		return v, json.Unmarshal(jsonBytes, &v)

	default:
		return nil, fmt.Errorf("template not valid")
	}
}
