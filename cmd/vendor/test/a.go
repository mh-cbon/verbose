package a

import (
  "fmt"

  "github.com/mh-cbon/verbose"
)


var logger = verbose.Auto()

func Hello () {
  fmt.Println("hello from test/a")
  logger.Println("Logged message from vendored test/a")
}
