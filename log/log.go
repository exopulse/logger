package log

import (
	"log"

	"exobyte.org/pulse/logger/debug"
)

// Info prints to the standard logger.
func Info(v ...interface{}) {
	log.Println(v...)
}

// Infof prints formatted text to the standard logger.
func Infof(format string, v ...interface{}) {
	log.Printf(format+"\n", v...)
}

// Debug prints to the standard logger, if debug mode is on.
func Debug(v ...interface{}) {
	if debug.On() {
		log.Println(v...)
	}
}

// Debugf prints formatted text to the standard logger, if debug mode is on.
func Debugf(format string, v ...interface{}) {
	if debug.On() {
		log.Printf(format+"\n", v...)
	}
}
