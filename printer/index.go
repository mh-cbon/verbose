package printer

import (
  "fmt"
  "log"
)

type Printer interface {
  Printf(format string, a ...interface{})
  Print(v ...interface{})
  Println(v ...interface{})
}

type FmtPrinter struct{}

func (p FmtPrinter) Printf (format string, a ...interface{}) {
  fmt.Printf(format, a...)
}

func (p FmtPrinter) Print (a ...interface{}) {
  fmt.Print(a...)
}

func (p FmtPrinter) Println (a ...interface{}) {
  fmt.Println(a...)
}

type LogPrinter struct{}

func (p LogPrinter) Printf (format string, a ...interface{}) {
  log.Printf(format, a...)
}

func (p LogPrinter) Print (a ...interface{}) {
  log.Print(a...)
}

func (p LogPrinter) Println (a ...interface{}) {
  log.Println(a...)
}
