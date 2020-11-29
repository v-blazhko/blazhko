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
	auth := smtp.PlainAuth("", fromAddress, fromPassword, smtpHost)

	from := mail.Address{
		Name:    "Common blazhko.tech mailer",
		Address: fromAddress,
	}
	to := mail.Address{
		Address: toAddress,
	}

	subject := "New blazhko.tech authorMessage"
	body := fmt.Sprintf("Author: %s (%s) \n Message: %s", authorName, authorEmail, authorMessage)

	headers := make(map[string]string)
	headers["From"] = from.String()
	headers["To"] = to.String()
	headers["Subject"] = subject

	message := ""
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + body

	serverName := smtpHost
	host, _, _ := net.SplitHostPort(serverName)

	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         host,
	}

	conn, err := tls.Dial("tcp", serverName, tlsConfig)
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

	_, err = w.Write([]byte(authorMessage))
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
