package main

import (
	"log"
	"net/smtp"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func sendReminderEmail() {
	from := os.Getenv("SMTP_EMAIL")
	password := os.Getenv("APP_PASSWORD")

	recipients := strings.Split(os.Getenv("REMINDER_EMAILS"), ",")

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	subject := "Maganghub Attendance Reminder"
	body := `
		<h3>Maganghub Attendance</h3>
		<p>Please remember to check and submit your attendance for today.</p>

		<p>
		<a href="https://monev.maganghub.kemnaker.go.id/dashboard"
			target="_blank"
			style="color:#1e40af;font-weight:600;">
			Click here to go to Attendance Dashboard
		</a>
		</p>

		<p><small>This reminder is sent automatically.</small></p>
	`

	msg := "From: " + from + "\r\n" +
		"To: " + strings.Join(recipients, ", ") + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"MIME-Version: 1.0\r\n" +
		"Content-Type: text/html; charset=\"UTF-8\"\r\n\r\n" +
		body

	auth := smtp.PlainAuth("", from, password, smtpHost)

	if err := smtp.SendMail(
		smtpHost+":"+smtpPort,
		auth,
		from,
		recipients,
		[]byte(msg),
	); err != nil {
		log.Fatal("Failed to send email:", err)
	}

	log.Println("Email reminder sent to:", recipients)
}

func main() {
	// Optional: load .env locally, Render will use Environment Variables
	if err := godotenv.Load(); err != nil {
		log.Println(".env not found, using Render environment variables")
	}

	sendReminderEmail()
}
