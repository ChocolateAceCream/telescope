package utils

import (
	"github.com/ChocolateAceCream/telescope/backend/singleton"
	"go.uber.org/zap"
	"gopkg.in/gomail.v2"
)

func SendMail(email, subject, body string) (err error) {
	msg := gomail.NewMessage()
	config := singleton.Config.Email
	msg.SetHeader("From", config.Username)
	msg.SetHeader("To", email)
	msg.SetHeader("Subject", subject)
	msg.SetBody("text/html", body)

	n := gomail.NewDialer(config.Host, config.Port, config.Username, config.Password)

	// Send the email
	err = n.DialAndSend(msg)
	if err != nil {
		singleton.Logger.Error("Failed to send email", zap.Error(err))
	}
	return
}
