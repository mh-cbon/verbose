// defines a Printer interface and provide two default implementation
package printer

import (
	"fmt"
	"log"
)

// Printer is the interface to implement to log message via verbose package
type Printer interface {
	Printf(pkg string, format string, a ...interface{})
	Print(pkg string, v ...interface{})
	Println(pkg string, v ...interface{})
}

// A wrapper of go/fmt package
type FmtPrinter struct{}

// Print with a format
func (p FmtPrinter) Printf(pkg string, format string, a ...interface{}) {
	format = "%s: " + format
	b := make([]interface{}, 1)
	b[0] = pkg
	a = append(b, a...)
	fmt.Printf(format, a...)
}

// Print given arguments
func (p FmtPrinter) Print(pkg string, a ...interface{}) {
	pkg = pkg + ":"
	b := make([]interface{}, 1)
	b[0] = pkg
	a = append(b, a...)
	fmt.Print(a...)
}

// Print given arguments with an ending line
func (p FmtPrinter) Println(pkg string, a ...interface{}) {
	pkg = pkg + ":"
	b := make([]interface{}, 1)
	b[0] = pkg
	a = append(b, a...)
	fmt.Println(a...)
}

// A wrapper of go/log package
type LogPrinter struct{}

// Print with a format
func (p LogPrinter) Printf(pkg string, format string, a ...interface{}) {
	format = "%s: " + format
	b := make([]interface{}, 1)
	b[0] = pkg
	a = append(b, a...)
	log.Printf(format, a...)
}

// Print given arguments
func (p LogPrinter) Print(pkg string, a ...interface{}) {
	pkg = pkg + ":"
	b := make([]interface{}, 1)
	b[0] = pkg
	a = append(b, a...)
	log.Print(a...)
}

// Print given arguments with an ending line
func (p LogPrinter) Println(pkg string, a ...interface{}) {
	pkg = pkg + ":"
	b := make([]interface{}, 1)
	b[0] = pkg
	a = append(b, a...)
	log.Println(a...)
}
