package main

import (
  "fmt"
  "github.com/mh-cbon/verbose"
  "github.com/mh-cbon/verbose/printer"
  "test"
)

var logger = verbose.Auto()

func main () {

  verbose.SetPrinter(printer.LogPrinter{})

  fmt.Println("Hello from main")
  logger.Println("message to log from main github.com/mh-cbon/verbose")
  a.Hello()
}
