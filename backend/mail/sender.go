package mail

import (
	"fmt"
	"log"
	"net/smtp"
	"os"
)

func SendToMyself(message string, authorName string, authorEmail string) error {
	from := os.Getenv("MAIL_FROM")
	fromPassword := os.Getenv("MAIL_FROM_PASSWORD")
	imapHost := os.Getenv("MAIL_IMAP_HOST")
	smtpHost := os.Getenv("MAIL_SMTP_HOST")
	to := []string{os.Getenv("MAIL_TO")}
	auth := smtp.PlainAuth("", from, fromPassword, imapHost)

	contents := fmt.Sprintf("New blazhko.tech message!\n"+
		"Author: %s (%s)"+
		"Message: %s", authorName, authorEmail, message)
	err := smtp.SendMail(smtpHost, auth, from, to, []byte(contents))
	if err != nil {
		log.Fatal(err)
	}

	return err
}
