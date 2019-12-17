package mailer

import (
	"fmt"
	"github.com/xvrzhao/site-monitor/parser"
	"net/smtp"
	"sync"
	"time"
)

var (
	lastSentMux sync.Mutex
	lastSent    time.Time
)

func Send(mailBody string) {
	lastSentMux.Lock()
	defer lastSentMux.Unlock()

	if d, _ := time.ParseDuration("10m"); time.Now().Sub(lastSent) >= d {

		parser.Debug("* Sending warning mail ...\n")
		from := fmt.Sprintf("%s <%s>", parser.MailFromName, parser.MailFromAddr)
		subject := "site-monitor warning"
		contentType := "text/plain; charset=utf-8"
		body := mailBody

		msg := fmt.Sprintf("To: %s\r\nFrom: %s\r\nSubject: %s\r\nContent-Type: %s\r\n\r\n%s",
			parser.MailToAddr, from, subject, contentType, body)
		auth := smtp.PlainAuth("", parser.MailFromAddr, parser.MailFromPwd, parser.MailSMTPAuthHost)
		err := smtp.SendMail(parser.MailSMTPServerAddr, auth, parser.MailFromAddr, []string{parser.MailToAddr}, []byte(msg))
		if err != nil {
			parser.Debug("* Sending mail failed: %s\n", err.Error())
		} else {
			lastSent = time.Now()
			parser.Debug("* Warning mail has been sent to %s\n", parser.MailToAddr)
		}
	}
}
