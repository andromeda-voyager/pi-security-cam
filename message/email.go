package message

import (
	"log"
	"net/smtp"
	"piSecurityCam/config"
)

func sendEmail(msg []byte) {
	hostname := "smtp.gmail.com"
	auth := smtp.PlainAuth("", config.EmailAccount(),
		config.Password(), hostname)

	err := smtp.SendMail(hostname+":587", auth, config.EmailAccount(), []string{config.EmailAccount()}, msg)
	if err != nil {
		log.Fatal(err)
	}
}
