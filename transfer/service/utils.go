package service

import (
	"fmt"
)

// Log a text to docker context
func Log(text string, a ...interface{}) {
	formatted := fmt.Sprintf("%s\n", text)
	fmt.Fprintf(cli.Out(), formatted, a...)
}

// Debug a text to docker context
func Debug(text string, a ...interface{}) {
	if debug {
		formatted := fmt.Sprintf("%s\n", text)
		fmt.Fprintf(cli.Out(), formatted, a...)
	}
}

// Error a text to docker context
func Error(text string, a ...interface{}) {
	formatted := fmt.Sprintf("%s\n", text)
	fmt.Fprintf(cli.Err(), formatted, a...)
}
