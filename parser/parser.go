package parser

import (
	"flag"
	"log"
	"time"
)

var (
	monitorIntervalStr, monitorTotalTimeoutStr, monitorRespHeaderTimeoutStr string
	debug                                                                   bool
)

var (
	MonitorURL               string
	MonitorInterval          time.Duration
	MonitorRespHeaderTimeout time.Duration
	MonitorTotalTimeout      time.Duration

	MailFromName       string
	MailFromAddr       string
	MailFromPwd        string
	MailToAddr         string
	MailSMTPAuthHost   string
	MailSMTPServerAddr string
)

func Debug(format string, v ...interface{}) {
	if debug {
		log.Printf(format, v...)
	}
}

func init() {
	flag.BoolVar(&debug, "debug-mode", false, "turn on debug mode")

	flag.StringVar(&MonitorURL, "monitor-url", "", "the url for monitor")
	flag.StringVar(&monitorIntervalStr, "monitor-interval", "20000", "milli second, testing cycle")
	flag.StringVar(&monitorTotalTimeoutStr, "monitor-total-timeout", "5000", "milli second, timeout for waiting url response")
	flag.StringVar(&monitorRespHeaderTimeoutStr, "monitor-respheader-timeout", "2000", "milli second, timeout for waiting response header")

	flag.StringVar(&MailFromName, "mail-from-name", "", "sender's name")
	flag.StringVar(&MailFromAddr, "mail-from-addr", "", "sender's mail address")
	flag.StringVar(&MailFromPwd, "mail-from-password", "", "sender's smtp auth code")
	flag.StringVar(&MailToAddr, "mail-to-addr", "", "receiver's mail address")
	flag.StringVar(&MailSMTPAuthHost, "mail-smtp-auth-host", "", "smtp auth host")
	flag.StringVar(&MailSMTPServerAddr, "mail-smtp-server-addr", "", "smtp server addr")

	flag.Parse()

	if MonitorURL == "" {
		log.Fatalf("missing monitor param: --monitor-url\n")
	}
	convertDuration(&monitorIntervalStr, &MonitorInterval)
	convertDuration(&monitorTotalTimeoutStr, &MonitorTotalTimeout)
	convertDuration(&monitorRespHeaderTimeoutStr, &MonitorRespHeaderTimeout)

	var miss string
	if MailFromName == "" {
		miss += "--mail-from-name, "
	}
	if MailFromAddr == "" {
		miss += "--mail-from-addr, "
	}
	if MailFromPwd == "" {
		miss += "--mail-from-password, "
	}
	if MailToAddr == "" {
		miss += "--mail-to-addr, "
	}
	if MailSMTPAuthHost == "" {
		miss += "--mail-smtp-auth-host, "
	}
	if MailSMTPServerAddr == "" {
		miss += "--mail-smtp-server-addr, "
	}
	if miss != "" {
		log.Fatalf("missing mail param(s): %srun `site-monitor --help` for more information.", miss)
	}
}

func convertDuration(str *string, du *time.Duration) {
	d, err := time.ParseDuration(*str + "ms")
	if err != nil {
		log.Fatalln("input monitor time param error")
	}
	*du = d
}
