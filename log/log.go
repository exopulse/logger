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

// DebugClose invokes closer function and logs error if any. Typical use for this function is for deferring a close.
func DebugClose(closer func() error) {
	if err := closer(); err != nil {
		Debug(err)
	}
}

// Error prints to the standard logger.
func Error(v ...interface{}) {
	var vs []interface{}

	vs = append(vs, "*ERROR*")
	vs = append(vs, v...)

	log.Println(vs...)
}

// Errorf prints formatted text to the standard logger.
func Errorf(format string, v ...interface{}) {
	log.Printf("*ERROR* "+format+"\n", v...)
}
