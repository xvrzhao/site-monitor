package main

import (
	"flag"
	"log"
	"strconv"
	"time"
)

var (
	debugFlag         bool
	urlFlag           string
	cycleFlag         time.Duration
	timeoutFlag       time.Duration
	headerTimeoutFlag time.Duration

	mailFromNameFlag   string
	mailFromAddrFlag   string
	mailFromPwdFlag    string
	mailToAddrFlag     string
	mailAuthHostFlag   string
	mailServerAddrFlag string
)

func init() {
	var cycle, timeout, headerTimeout int

	flag.BoolVar(&debugFlag, "debug", false, "turn on debug mode")

	flag.StringVar(&urlFlag, "url", "", "HTTP URL to monitor")
	flag.IntVar(&cycle, "cycle", 20000, "detection cycle in milliseconds")
	flag.IntVar(&timeout, "timeout", 5000, "timeout for waiting for HTTP response, in milliseconds")
	flag.IntVar(&headerTimeout, "header-timeout", 2000, "timeout for waiting for HTTP response headers after establishing a connection, in milliseconds")

	flag.StringVar(&mailFromNameFlag, "mail-fname", "Site Monitor", "mail sender's name")
	flag.StringVar(&mailFromAddrFlag, "mail-faddr", "", "mail sender's address")
	flag.StringVar(&mailFromPwdFlag, "mail-fpwd", "", "mail sender's SMTP password")
	flag.StringVar(&mailToAddrFlag, "mail-taddr", "", "mail recipient address")
	flag.StringVar(&mailAuthHostFlag, "mail-auth", "", "SMTP authentication host address")
	flag.StringVar(&mailServerAddrFlag, "mail-server", "", "SMTP server address")

	//testing.Init() // uncomment when executing unit tests
	flag.Parse()

	var miss string

	if urlFlag == "" {
		miss += "-url, "
	}
	if mailFromNameFlag == "" {
		miss += "-mail-fname, "
	}
	if mailFromAddrFlag == "" {
		miss += "-mail-faddr, "
	}
	if mailFromPwdFlag == "" {
		miss += "-mail-fpwd, "
	}
	if mailToAddrFlag == "" {
		miss += "-mail-taddr, "
	}
	if mailAuthHostFlag == "" {
		miss += "-mail-auth, "
	}
	if mailServerAddrFlag == "" {
		miss += "-mail-server, "
	}
	if miss != "" {
		log.Fatalf("missing the necessary flag(s): %srun \"site-monitor -h\" for more information about each flag.", miss)
	}

	convertDuration(cycle, &cycleFlag)
	convertDuration(timeout, &timeoutFlag)
	convertDuration(headerTimeout, &headerTimeoutFlag)
}

func convertDuration(n int, du *time.Duration) {
	d, err := time.ParseDuration(strconv.Itoa(n) + "ms")
	if err != nil {
		log.Fatal("the input time format is not available")
	}
	*du = d
}
