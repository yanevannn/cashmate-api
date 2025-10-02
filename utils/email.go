package utils

import (
	"fmt"
	"net/smtp"
	"os"
)

func SendEmailVerification(toEmail string, otp string) error {
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")
	smtpUser := os.Getenv("SMTP_USER")
	smtpPass := os.Getenv("SMTP_PASS")

	subject := "Email Verification\r\n"
	mime := "MIME-version: 1.0;\r\nContent-Type: text/html; charset=\"UTF-8\";\r\n\r\n"

	body := fmt.Sprintf(`
						<html>
							<body>
								<h2 style="color:#4CAF50;">Cashmate Verification</h2>
								<p>Hi 👋,</p>
								<p>Your verification code is:</p>
								<h3 style="color:#E91E63;">%s</h3>
								<p>Please enter this code in the app to verify your account.</p>
								<br>
								<small>⚠️ This code will expire in 10 minutes.</small>
							</body>
						</html>
				`, otp)

	message := []byte(subject + mime + body)

	auth := smtp.PlainAuth("", smtpUser, smtpPass, smtpHost)

	return smtp.SendMail(smtpHost+":"+smtpPort, auth, smtpUser, []string{toEmail}, message)
}
