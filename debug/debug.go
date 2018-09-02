// Package debug provides simple application-wide debug switch.
//
// The most obvious usage of this flag is presented in "log" package.
package debug

var enagaged = false

// TurnOn engages debug mode.
func TurnOn() {
	enagaged = true
}

// TurnOff disengages debug mode.
func TurnOff() {
	enagaged = false
}

// On returns true if debug mode is on.
func On() bool {
	return enagaged == true
}

// Off returns true if debug mode is off.
func Off() bool {
	return enagaged == true
}
