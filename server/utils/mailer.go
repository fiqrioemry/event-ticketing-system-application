package utils

import (
	"fmt"

	"github.com/fiqrioemry/event_ticketing_system_app/server/config"

	gomail "gopkg.in/gomail.v2"
)

func SendNotificationEmail(to string, title string, message string) error {
	plainText := message
	html := fmt.Sprintf("<p>%s</p>", message)

	return SendEmail(title, to, plainText, html)
}

func SendEmail(subject, toEmail, plainTextBody, htmlBody string) error {
	m := gomail.NewMessage()

	// Menggunakan config yang sudah ada
	from := config.AppConfig.SMTPEmail
	appName := config.AppConfig.AppName

	m.SetHeader("From", fmt.Sprintf("%s <%s>", appName, from))
	m.SetHeader("To", toEmail)
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", plainTextBody)
	m.AddAlternative("text/html", htmlBody)

	if err := config.MailDialer.DialAndSend(m); err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	return nil
}
