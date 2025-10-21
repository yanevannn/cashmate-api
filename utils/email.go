package utils

import (
	"fmt"
	"net/smtp"
	"os"
)

func SendEmailVerification(toEmail string, otp string, name string, activationLink string) error {
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")
	smtpUser := os.Getenv("SMTP_USER")
	smtpPass := os.Getenv("SMTP_PASS")
	fromName := os.Getenv("SMTP_MAIL_FROM_NAME")
	fromAddress := os.Getenv("SMTP_MAIL_FROM_ADDRESS")

	from := fmt.Sprintf("From: %s <%s>\r\n", fromName, fromAddress)
	to := fmt.Sprintf("To: %s\r\n", toEmail)
	subject := "Subject: Email Verification\r\n"
	mime := "MIME-version: 1.0;\r\nContent-Type: text/html; charset=\"UTF-8\";\r\n\r\n"

	body := fmt.Sprintf(`
						<html>
							<body>
								<h2 style="color:#4CAF50;">Cashmate Verification</h2>
								<p>Hi %s 👋,</p>
								<p>Your verification code is:</p>
								<h3 style="color:#E91E63;">%s</h3>
								<p>Please enter this code in the app to verify your account.</p>
								<button style="background-color:#4CAF50; color:white; padding:10px 20px; text-align:center; text-decoration:none; display:inline-block; font-size:16px; margin:10px 0; border:none; border-radius:5px;">
									<a href="%s" style="color:white; text-decoration:none;">Verify Email</a>
								</button>
								<br>
								<small>⚠️ This code will expire in 15 minutes.</small>
							</body>
						</html>
				`, name, otp, activationLink)

	// Combine all parts of the email
	message := []byte(from + to + subject + mime + body)

	auth := smtp.PlainAuth("", smtpUser, smtpPass, smtpHost)

	return smtp.SendMail(smtpHost+":"+smtpPort, auth, smtpUser, []string{toEmail}, message)
}

func SendEmailResetPassword(toEmail string, otp string, name string) error {
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")
	smtpUser := os.Getenv("SMTP_USER")
	smtpPass := os.Getenv("SMTP_PASS")
	fromName := os.Getenv("SMTP_MAIL_FROM_NAME")
	fromAddress := os.Getenv("SMTP_MAIL_FROM_ADDRESS")

	from := fmt.Sprintf("From: %s <%s>\r\n", fromName, fromAddress)
	to := fmt.Sprintf("To: %s\r\n", toEmail)
	subject := "Subject: Reset Password Code\r\n"
	mime := "MIME-version: 1.0;\r\nContent-Type: text/html; charset=\"UTF-8\";\r\n\r\n"

	body := fmt.Sprintf(`
						<html>
							<body>
								<h2 style="color:#4CAF50;">Cashmate Password Reset</h2>
								<p>Hi %s 👋,</p>
								<p>Your password reset code is:</p>
								<h3 style="color:#E91E63;">%s</h3>
								<p>Please enter this code in the app to reset your password.</p>
								<br>
								<small>⚠️ This code will expire in 15 minutes.</small>
							</body>
						</html>
				`, name, otp)

	// Combine all parts of the email
	message := []byte(from + to + subject + mime + body)

	auth := smtp.PlainAuth("", smtpUser, smtpPass, smtpHost)

	return smtp.SendMail(smtpHost+":"+smtpPort, auth, smtpUser, []string{toEmail}, message)
}
