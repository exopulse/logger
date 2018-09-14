# exopulse logging/debugging support
Golang simple logging framework using informational, error and debugging messages only.

[![CircleCI](https://circleci.com/gh/exopulse/logger.svg?style=svg)](https://circleci.com/gh/exopulse/logger)
[![Build Status](https://travis-ci.org/exopulse/logger.svg?branch=master)](https://travis-ci.org/exopulse/logger)
[![GitHub license](https://img.shields.io/github/license/exopulse/logger.svg)](https://github.com/exopulse/logger/blob/master/LICENSE)

# Overview

This module contains two packages:
- log - logging support
- debug - debugging support

# Using logger package

## Installing package

Use go get to install the latest version of the library.

    $ go get github.com/exopulse/logger
 
Include logger in your application.
```go
import "github.com/exopulse/logger"
import "github.com/exopulse/debug"
```

## Features

### Package log

Package log provides simple logging framework using informational, error and debugging messages only. This package provides a rolling logger also.

This implementation is inspired by Dave Cheney's post:
   
   https://dave.cheney.net/2015/11/05/lets-talk-about-logging

Expected usage is to:
- configure application-wide rolling logger using InstallLogger() method
- log informational and debug messages using "log" package

See "debug" package for setting debug mode on or off.

#### Logging functions

```go
// Info prints to the standard logger.
func Info(v ...interface{})

// Infof prints formatted text to the standard logger.
func Infof(format string, v ...interface{})

// Debug prints to the standard logger, if debug mode is on.
func Debug(v ...interface{})

// Debugf prints formatted text to the standard logger, if debug mode is on.
func Debugf(format string, v ...interface{})

// DebugClose invokes closer function and logs error if any. Typical use for this function is for deferring a close.
func DebugClose(closer func() error)

// Error prints to the standard logger.
func Error(v ...interface{})

// Errorf prints formatted text to the standard logger.
func Errorf(format string, v ...interface{})
```

#### Install logger

    TODO: see debug/debug.go
 
### Package debug

This package provides means to engage/disengage debug mode and query current debug mode.

#### Available functions

```go
// TurnOn engages debug mode.
func TurnOn()

// TurnOff disengages debug mode.
func TurnOff()

// On returns true if debug mode is on.
func On() bool

// Off returns true if debug mode is off.
func Off() bool
```

# About the project

## Contributors

* [exopulse](https://github.com/exopulse)

## License

logger package is released under the MIT license. See
[LICENSE](https://github.com/exopulse/logger/blob/master/LICENSE)
