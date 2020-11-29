package mail

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/mail"
	"net/smtp"
	"os"
)

func SendToMyself(authorMessage string, authorName string, authorEmail string) error {
	fromAddress := os.Getenv("MAIL_FROM")
	fromPassword := os.Getenv("MAIL_FROM_PASSWORD")
	toAddress := os.Getenv("MAIL_TO")
	smtpHost := os.Getenv("MAIL_SMTP_HOST")

	from := mail.Address{
		Name:    "Common blazhko.tech mailer",
		Address: fromAddress,
	}
	to := mail.Address{
		Name:    "Veronica",
		Address: toAddress,
	}

	subject := "New blazhko.tech Message"
	body := fmt.Sprintf("Author: %s (%s) \nMessage: %s", authorName, authorEmail, authorMessage)

	headers := make(map[string]string)
	headers["From"] = from.String()
	headers["To"] = to.String()
	headers["Subject"] = subject

	message := ""
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + body

	servername := smtpHost

	host, _, _ := net.SplitHostPort(servername)

	auth := smtp.PlainAuth("", fromAddress, fromPassword, host)

	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         host,
	}

	conn, err := tls.Dial("tcp", servername, tlsconfig)
	if err != nil {
		return err
	}

	c, err := smtp.NewClient(conn, host)
	if err != nil {
		return err
	}

	if err = c.Auth(auth); err != nil {
		return err
	}

	if err = c.Mail(from.Address); err != nil {
		return err
	}

	if err = c.Rcpt(to.Address); err != nil {
		return err
	}

	w, err := c.Data()
	if err != nil {
		return err
	}

	_, err = w.Write([]byte(message))
	if err != nil {
		return err
	}

	err = w.Close()
	if err != nil {
		return err
	}

	_ = c.Quit()

	return nil
}
