package main

import (
	"fmt"
	"net/smtp"
	"strings"
	"sync"
	"time"
)

var (
	mailMux      sync.Mutex
	mailLastTime time.Time
)

func sendMail(mailBody string) {
	mailMux.Lock()
	defer mailMux.Unlock()

	if d, _ := time.ParseDuration("10m"); time.Now().Sub(mailLastTime) >= d {

		debug("* Sending warning mail ...")
		from := fmt.Sprintf("%s <%s>", mailFromNameFlag, mailFromAddrFlag)
		subject := "site-monitor warning"
		contentType := "text/plain; charset=utf-8"
		body := mailBody

		to := strings.Join(mailToAddrFlag, ", ")
		msg := fmt.Sprintf("To: %s\r\nFrom: %s\r\nSubject: %s\r\nContent-Type: %s\r\n\r\n%s",
			to, from, subject, contentType, body)
		auth := smtp.PlainAuth("", mailFromAddrFlag, mailFromPwdFlag, mailAuthHostFlag)
		err := smtp.SendMail(mailServerAddrFlag, auth, mailFromAddrFlag, mailToAddrFlag, []byte(msg))
		if err != nil {
			debug("* Sending mail failed: %v", err)
		} else {
			debug("* Mail has been sent to %s", to)
		}
		mailLastTime = time.Now()
	}
}
