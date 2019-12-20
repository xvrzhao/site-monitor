package parser

import (
	"flag"
	"log"
	"strconv"
	"time"
)

var (
	monitorInterval, monitorTotalTimeout, monitorRespHeaderTimeout int
	debug                                                          bool
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
	flag.BoolVar(&debug, "debug", false, "turn on debug mode")

	flag.StringVar(&MonitorURL, "url", "", "HTTP URL to monitor")
	flag.IntVar(&monitorInterval, "cycle", 20000, "detection cycle in milliseconds")
	flag.IntVar(&monitorTotalTimeout, "timeout", 5000, "timeout for waiting for HTTP response, in milliseconds")
	flag.IntVar(&monitorRespHeaderTimeout, "header-timeout", 2000, "timeout for waiting for HTTP response headers after establishing a connection, in milliseconds")

	flag.StringVar(&MailFromName, "mail-fname", "Site Monitor", "mail sender's name")
	flag.StringVar(&MailFromAddr, "mail-faddr", "", "mail sender's address")
	flag.StringVar(&MailFromPwd, "mail-fpwd", "", "mail sender's SMTP password")
	flag.StringVar(&MailToAddr, "mail-taddr", "", "mail recipient address")
	flag.StringVar(&MailSMTPAuthHost, "mail-auth", "", "SMTP authentication host address")
	flag.StringVar(&MailSMTPServerAddr, "mail-server", "", "SMTP server address")

	//testing.Init() // uncomment when executing unit tests
	flag.Parse()

	convertDuration(monitorInterval, &MonitorInterval)
	convertDuration(monitorTotalTimeout, &MonitorTotalTimeout)
	convertDuration(monitorRespHeaderTimeout, &MonitorRespHeaderTimeout)

	var miss string

	if MonitorURL == "" {
		miss += "-url, "
	}
	if MailFromName == "" {
		miss += "-mail-fname, "
	}
	if MailFromAddr == "" {
		miss += "-mail-faddr, "
	}
	if MailFromPwd == "" {
		miss += "-mail-fpwd, "
	}
	if MailToAddr == "" {
		miss += "-mail-taddr, "
	}
	if MailSMTPAuthHost == "" {
		miss += "-mail-auth, "
	}
	if MailSMTPServerAddr == "" {
		miss += "-mail-server, "
	}
	if miss != "" {
		log.Fatalf("missing the necessary flag(s): %srun `site-monitor -h` for more information about each flag.", miss)
	}
}

func convertDuration(n int, du *time.Duration) {
	d, err := time.ParseDuration(strconv.Itoa(n) + "ms")
	if err != nil {
		log.Fatalln("the input time format is not available")
	}
	*du = d
}
