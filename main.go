package main

import (
	"crypto/tls"
	"fmt"
	"net/smtp"
)

func main() {
	// parameters for smtp server
	smtpHost := "smtp.gmail.com"
	smtpPort := "587" // or use "465 as smtpPort
	smtpUser := "your-email@gmail.com"
	smtpPass := "your-email-password"

	// Parameters for email
	from := "christian.ngonoabanda@ynov.com"
	to := "christianmanga384@gmail.com"
	subject := "Test email"
	body := "Hello, this is a test email"

	// Configuration TLS for smtp server connection
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         smtpHost,
	}

	// Connect to the smtp server
	client, err := smtp.Dial(fmt.Sprintf("%s:%d", smtpHost, smtpPort))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Authentification to the smtp server

	auth := smtp.PlainAuth("", smtpUser, smtpPass, smtpHost)
	if err := client.Auth(auth); err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Start TLS
	if err := client.StartTLS(tlsConfig); err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Send the email
	if err := client.Mail(from); err != nil {
		fmt.Println("Error:", err)
		return
	}
	if err := client.Rcpt(to); err != nil {
		fmt.Println("Error, err")
		return
	}
	w, err := client.Data()
	if err != nil {
		fmt.Println("Error:", err)
	}
	_, err = w.Write([]byte("To: " + to + "\r\n" + "Subject: " + subject + "\r\n" + "\r\n" + body))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	err = w.Close()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	client.Quit()
	fmt.Println("Email sent successfully")
}
