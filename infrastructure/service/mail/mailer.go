package mail

import (
	"fmt"
	"log"
	"net/smtp"
	"os"
)

func SendMail(body string, receiver []string) {
	// Konfigurasi SMTP
	host := os.Getenv("MAIL_ADDRESS")
	port := os.Getenv("MAIL_PORT")
	username := os.Getenv("MAIL_USERNAME")
	password := os.Getenv("MAIL_PASSWORD")

	subject := "Alert: Middleware Critical Issue"

	// Menggabungkan Subject dan Body
	message := []byte(subject + "\r\n" + body)

	// Konfigurasi SMTP
	auth := smtp.PlainAuth("", username, password, host)

	// Kirim email menggunakan SMTP
	address := fmt.Sprintf("%s:%v", host, port)
	err := smtp.SendMail(address, auth, username, receiver, message)
	if err != nil {
		log.Fatal("Error sending mail", err)
	}

	log.Println("Mail sent!")
}
