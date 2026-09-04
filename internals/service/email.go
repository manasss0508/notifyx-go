package service

import (
	"fmt"
	"net/smtp"
)

type EmailService struct {
	Host     string
	Port     string
	Username string
	Password string
	From     string
}

func NewEmailService(
	host string,
	port string,
	username string,
	password string,
	from string,
) *EmailService {
	return &EmailService{
		Host:     host,
		Port:     port,
		Username: username,
		Password: password,
		From:     from,
	}
}

func (s *EmailService) Send(
	to string,
	subject string,
	body string,
) error {

	// creating credentials
	auth := smtp.PlainAuth(
		"",
		s.Username,
		s.Password,
		s.Host,
	)

	// creating mail and converting it to bytes
	message := []byte(
		"From: " + s.From + "\r\n" +
			"To: " + to + "\r\n" +
			"Subject: " + subject + "\r\n" +
			"Content-Type: text/plain; charset=UTF-8\r\n" +
			"\r\n" +
			body + "\r\n",
	)
	fmt.Println("mail : ", message)

	// smtp address
	address := s.Host + ":" + s.Port

	// sending mail
	err := smtp.SendMail(
		address,
		auth,
		s.From,
		[]string{to},
		message,
	)

	if err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	return nil
}
