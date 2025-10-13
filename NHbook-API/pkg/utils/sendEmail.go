package utils

import (
	"gopkg.in/gomail.v2"
)

func SendEmail(to string, subject string, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", "your-email@gmail.com")
	m.SetHeader("To", "recipient@example.com")
	m.SetHeader("Subject", "Hello from Go!")
	m.SetBody("text/plain", "This is the plain text body.")

	d := gomail.NewDialer("smtp.gmail.com", 587, "your-email@gmail.com", "your-app-password")

	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}
