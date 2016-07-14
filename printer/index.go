// defines a Printer interface and provide two default implementation
package printer

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/mh-cbon/verbose/color"
)

var currentPrinter Printer

// Configure current Printer
func SetPrinter(p Printer) {
	currentPrinter = p
}

// instance of a logger
type Logger struct {
	Name    string
	Enabled bool
	Color   color.ColorFunc
}

// Methods to display messages
func (l *Logger) Printf(format string, a ...interface{}) {
	if l.Enabled {
		currentPrinter.Printf(l, format, a...)
	}
}
func (l *Logger) Print(a ...interface{}) {
	if l.Enabled {
		currentPrinter.Print(l, a...)
	}
}
func (l *Logger) Println(a ...interface{}) {
	if l.Enabled {
		currentPrinter.Println(l, a...)
	}
}

// Printer is the interface to implement to log message via verbose package
type Printer interface {
	Printf(logger *Logger, format string, a ...interface{})
	Print(logger *Logger, v ...interface{})
	Println(logger *Logger, v ...interface{})
}

// A wrapper of go/fmt package
type FmtPrinter struct{}

// Print with a format
func (p FmtPrinter) Printf(logger *Logger, format string, a ...interface{}) {
	format = "%s " + format
	b := make([]interface{}, 1)
	b[0] = logger.Color(logger.Name)
	a = append(b, a...)
	fmt.Fprintf(os.Stderr, format, a...)
}

// Print given arguments
func (p FmtPrinter) Print(logger *Logger, a ...interface{}) {
	b := make([]interface{}, 1)
	b[0] = logger.Color(logger.Name)
	a = append(b, a...)
	fmt.Fprint(os.Stderr, a...)
}

// Print given arguments with an ending line
func (p FmtPrinter) Println(logger *Logger, a ...interface{}) {
	b := make([]interface{}, 1)
	b[0] = logger.Color(logger.Name)
	a = append(b, a...)
	format := strings.Repeat("%s ", len(b))
	fmt.Fprintf(os.Stderr, format, a...)
}

// A wrapper of go/log package
type LogPrinter struct{}

// Print with a format
func (p LogPrinter) Printf(logger *Logger, format string, a ...interface{}) {
	format = "%s " + format
	b := make([]interface{}, 1)
	b[0] = logger.Color(logger.Name)
	a = append(b, a...)
	log.Printf(format, a...)
}

// Print given arguments
func (p LogPrinter) Print(logger *Logger, a ...interface{}) {
	b := make([]interface{}, 1)
	b[0] = logger.Color(logger.Name)
	a = append(b, a...)
	log.Print(a...)
}

// Print given arguments with an ending line
func (p LogPrinter) Println(logger *Logger, a ...interface{}) {
	b := make([]interface{}, 1)
	b[0] = logger.Color(logger.Name)
	a = append(b, a...)
	log.Println(a...)
}
