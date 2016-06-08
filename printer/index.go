package printer

import (
  "fmt"
  "log"
)

type Printer interface {
  Printf(pkg string, format string, a ...interface{})
  Print(pkg string, v ...interface{})
  Println(pkg string, v ...interface{})
}

type FmtPrinter struct{}

func (p FmtPrinter) Printf (pkg string, format string, a ...interface{}) {
  format = "%s: " + format
  b := make([]interface{}, 1)
  b[0] = pkg
  a = append(b, a...)
  fmt.Printf(format, a...)
}

func (p FmtPrinter) Print (pkg string, a ...interface{}) {
  pkg = pkg + ":"
  b := make([]interface{}, 1)
  b[0] = pkg
  a = append(b, a...)
  fmt.Print(a...)
}

func (p FmtPrinter) Println (pkg string, a ...interface{}) {
  pkg = pkg + ":"
  b := make([]interface{}, 1)
  b[0] = pkg
  a = append(b, a...)
  fmt.Println(a...)
}

type LogPrinter struct{}

func (p LogPrinter) Printf (pkg string, format string, a ...interface{}) {
  format = "%s: " + format
  b := make([]interface{}, 1)
  b[0] = pkg
  a = append(b, a...)
  log.Printf(format, a...)
}

func (p LogPrinter) Print (pkg string, a ...interface{}) {
  pkg = pkg + ":"
  b := make([]interface{}, 1)
  b[0] = pkg
  a = append(b, a...)
  log.Print(a...)
}

func (p LogPrinter) Println (pkg string, a ...interface{}) {
  pkg = pkg + ":"
  b := make([]interface{}, 1)
  b[0] = pkg
  a = append(b, a...)
  log.Println(a...)
}
