package templates

import "errors"

type TemplateVariables interface {
	Validate() error
}

type WelcomeVariables struct {
	Name string `json:"name"`
}

func (v WelcomeVariables) Validate() error {
	if v.Name == "" {
		return errors.New("template variables 'name' not present ")
	}
	return nil
}

type OtpVariables struct {
	Name          string `json:"name"`
	OTP           string `json:"otp"`
	ExpiryMinutes string `json:"expiry_minutes"`
}

func (v OtpVariables) Validate() error {
	if v.Name == "" {
		return errors.New("template variables 'name' not present ")
	}
	if v.OTP == "" {
		return errors.New("template variables 'otp' not present ")
	}
	if v.ExpiryMinutes == "" {
		return errors.New("template variables 'expiry_minutes' not present ")
	}

	return nil
}

type PasswordResetVariables struct {
	Name          string `json:"name"`
	ResetLink     string `json:"reset_link"`
	ExpiryMinutes string `json:"expiry_minutes"`
}

func (v PasswordResetVariables) Validate() error {
	if v.Name == "" {
		return errors.New("template variables 'name' not present ")
	}
	if v.ResetLink == "" {
		return errors.New("template variables 'reset_link' not present ")
	}
	if v.ExpiryMinutes == "" {
		return errors.New("template variables 'expiry_minutes' not present ")
	}
	return nil
}

type EmailVerificationVariables struct {
	Name             string `json:"name"`
	VerificationLink string `json:"verification_link"`
}

func (v EmailVerificationVariables) Validate() error {
	if v.Name == "" {
		return errors.New("template variables 'name' not present ")
	}
	if v.VerificationLink == "" {
		return errors.New("template variables 'verification_link' not present ")
	}
	return nil
}

type LoginAlertVariables struct {
	Name      string `json:"name"`
	LoginTime string `json:"login_time"`
	Location  string `json:"location"`
	Device    string `json:"device"`
}

func (v LoginAlertVariables) Validate() error {
	if v.Name == "" {
		return errors.New("template variables 'name' not present ")
	}
	if v.LoginTime == "" {
		return errors.New("template variables 'login_time' not present ")
	}
	if v.Location == "" {
		return errors.New("template variables 'location' not present ")
	}
	if v.Device == "" {
		return errors.New("template variables 'device' not present ")
	}
	return nil
}

type OrderConfirmationVariables struct {
	OrderID  string `json:"order_id"`
	Name     string `json:"name"`
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
}

func (v OrderConfirmationVariables) Validate() error {
	if v.OrderID == "" {
		return errors.New("template variables 'order_id' not present ")
	}
	if v.Name == "" {
		return errors.New("template variables 'name' not present ")
	}
	if v.Amount == "" {
		return errors.New("template variables 'amount' not present ")
	}
	if v.Currency == "" {
		return errors.New("template variables 'currency' not present ")
	}
	return nil
}

type OrderShippedVariables struct {
	Name              string `json:"name"`
	OrderID           string `json:"order_id"`
	TrackingNumber    string `json:"tracking_number"`
	EstimatedDelivery string `json:"estimated_delivery"`
}

func (v OrderShippedVariables) Validate() error {
	if v.Name == "" {
		return errors.New("template variables 'name' not present ")
	}
	if v.OrderID == "" {
		return errors.New("template variables 'order_id' not present ")
	}
	if v.TrackingNumber == "" {
		return errors.New("template variables 'tracking_number' not present ")
	}
	if v.EstimatedDelivery == "" {
		return errors.New("template variables 'estimated_delivery' not present ")
	}
	return nil
}

type PaymentSuccessVariables struct {
	Name          string `json:"name"`
	TransactionID string `json:"transaction_id"`
	Amount        string `json:"amount"`
	Currency      string `json:"currency"`
}

func (v PaymentSuccessVariables) Validate() error {
	if v.Name == "" {
		return errors.New("template variables 'name' not present ")
	}
	if v.TransactionID == "" {
		return errors.New("template variables 'transaction_id' not present ")
	}
	if v.Amount == "" {
		return errors.New("template variables 'amount' not present ")
	}
	if v.Currency == "" {
		return errors.New("template variables 'currency' not present ")
	}
	return nil

}

type PaymentFailedVariables struct {
	Name     string `json:"name"`
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
	Reason   string `json:"reason"`
}

func (v PaymentFailedVariables) Validate() error {
	if v.Name == "" {
		return errors.New("template variables 'name' not present ")
	}
	if v.Amount == "" {
		return errors.New("template variables 'amount' not present ")
	}
	if v.Currency == "" {
		return errors.New("template variables 'currency' not present ")
	}
	if v.Reason == "" {
		return errors.New("template variables 'reason' not present ")
	}
	return nil
}

type SubscriptionRenewalVariables struct {
	Name            string `json:"name"`
	PlanName        string `json:"plan_name"`
	RenewalDate     string `json:"renewal_date"`
	NextBillingDate string `json:"next_billing_date"`
}

func (v SubscriptionRenewalVariables) Validate() error {
	if v.Name == "" {
		return errors.New("template variables 'name' not present ")
	}
	if v.PlanName == "" {
		return errors.New("template variables 'plan_name' not present ")
	}
	if v.RenewalDate == "" {
		return errors.New("template variables 'renewal_date' not present ")
	}
	if v.NextBillingDate == "" {
		return errors.New("template variables 'next_billing_date' not present ")
	}
	return nil
}
