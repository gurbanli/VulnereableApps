package util

import (
	"fmt"
	"net/smtp"
	"os"
)

func SendMail(toMail []string, message []byte) bool{
	fromMail := os.Getenv("from_mail")
	mailPassword := os.Getenv("mail_password")
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"
	auth := smtp.PlainAuth("", fromMail, mailPassword, smtpHost )
	err := smtp.SendMail(smtpHost + ":" + smtpPort, auth, fromMail, toMail, message)
	if err != nil{
		fmt.Println(err)
		return false
	}
	return true
}

func SendOTP(otpCode string, email string) bool{
	toMail := []string{
		email,
	}
	message := fmt.Sprintf("To: %s\r\n" +     "Subject: OTP\r\n" +     "\r\n" +     "OTP Code:%s\r\n", email, otpCode)
	return SendMail(toMail, []byte(message))
}

func SendForgotPasswordToken(resetToken string, email string) bool{
	toMail := []string{
		email,
	}
	message := fmt.Sprintf("To: %s\r\n" +     "Subject: Reset Password\r\n" +     "\r\n" +     "Reset Password Token:%s\r\n", email, resetToken)
	return SendMail(toMail, []byte(message))
}