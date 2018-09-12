package dbg

import (
	"fmt"

	"github.com/mgutz/ansi"
)

// Color is an enum for debug colors
type Color string

const (
	// ColorDefault debug
	ColorDefault Color = "default"
	// ColorGreen debug
	ColorGreen = "green"
	// ColorYellow debug
	ColorYellow = "yellow"
	// ColorRed debug
	ColorRed = "red"
	// ColorBlue debug
	ColorBlue = "blue"
	// ColorWhite debug
	ColorWhite = "white"
	// ColorBlack debug
	ColorBlack = "black"
	// ColorCyan debug
	ColorCyan = "cyan"
	// ColorMagenta debug
	ColorMagenta = "magenta"
)

//Sprint formats your string with colors
func Sprint(color Color, msg ...interface{}) string {
	return sprint(color, false, msg...)
}

//Sprintf formats your string with colors
func Sprintf(color Color, format string, msg ...interface{}) string {
	return sprint(color, false, fmt.Sprintf(format, msg...))
}

func sprint(color Color, isMap bool, msg ...interface{}) string {
	if color != ColorDefault {
		c := ansi.ColorCode(fmt.Sprintf("%s+h:black", color))
		reset := ansi.ColorCode("reset")

		msg = append([]interface{}{c}, msg...)
		msg = append(msg, reset)
	}

	if isMap {
		return fmt.Sprintf("%+v\n", msg...)
	}
	return fmt.Sprintln(msg...)
}

func dbg(color Color, isMap bool, msg ...interface{}) {
	fmt.Print(sprint(color, isMap, msg...))
}

// Warn prints the given params in Yellow
func Warn(msg ...interface{}) {
	dbg(ColorYellow, false, msg...)
}

// Error prints the given params in Red
func Error(msg ...interface{}) {
	dbg(ColorRed, false, msg...)
}

// ColorDebug prints a single param in given color
func ColorDebug(msg interface{}, color Color) {
	dbg(color, false, msg)
}

// Debug prints the given params without color (wraps fmt.Println)
func Debug(msg ...interface{}) {
	fmt.Println(msg...)
}

// DebugMap prints the given params as key:value maps
func DebugMap(msg ...interface{}) {
	dbg(ColorDefault, true, msg...)
}

// Green prints the given params in Green
func Green(msg ...interface{}) {
	dbg(ColorGreen, false, msg...)
}

// Yellow prints the given params in Yellow
func Yellow(msg ...interface{}) {
	dbg(ColorYellow, false, msg...)
}

// Red prints the given params in Red
func Red(msg ...interface{}) {
	dbg(ColorRed, false, msg...)
}

// Blue prints the given params in Blue
func Blue(msg ...interface{}) {
	dbg(ColorBlue, false, msg...)
}

// Cyan prints the given params in Cyan
func Cyan(msg ...interface{}) {
	dbg(ColorCyan, false, msg...)
}

// Magenta prints the given params in Magenta
func Magenta(msg ...interface{}) {
	dbg(ColorMagenta, false, msg...)
}
