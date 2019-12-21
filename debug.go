package main

import "log"

func debug(format string, v ...interface{}) {
	if debugFlag {
		log.Printf(format, v...)
	}
}
