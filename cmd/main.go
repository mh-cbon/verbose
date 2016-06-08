package main

import (
  "fmt"
  "github.com/mh-cbon/verbose"
  "test"
)

var logger = verbose.Auto()

func main () {
  fmt.Println("Hello from main")
  logger.Println("message to log from main github.com/mh-cbon/verbose")
  a.Hello()
}
