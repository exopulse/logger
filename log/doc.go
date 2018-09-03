// Package log provides simple logging framework using informational, error and debugging messages only. This package
// provides a rolling logger also.
//
// This implementation is inspired by Dave Cheney's post:
//   https://dave.cheney.net/2015/11/05/lets-talk-about-logging
//
// Expected usage is to:
//  - configure application-wide rolling logger using InstallLogger() method
//  - log informational and debug messages using "log" package
//
// See "debug" package for setting debug mode on or off.
package log
