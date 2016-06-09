# Verbose

A library to display verbose messages per package.

Read the doc at [godoc](https://godoc.org/github.com/mh-cbon/verbose)

![demo](https://raw.githubusercontent.com/mh-cbon/verbose/master/demo.jpg)

# Install

```sh
glide get github.com/mh-cbon/verbose
```

# Usage

```go
package main

import (
  "github.com/mh-cbon/verbose"
)

var logger = verbose.Auto()

func main () {
  logger.Println("something you want to tell about...")
}
```

# Command line

to display the log messages from your binary you will have to set `VERBOSE` env variable,

```sh
VERBOSE=package goprogram
```

Where `VERBOSE` is a comma separated list of package patterns to display,

```
VERBOSE=package*,package2 goprogram
```

when `VERBOSE=*` all log messages will display

```
VERBOSE=* goprogram
```

# Configuring the printer

Using `verbose.SetPrinter(p printer.Printer)` you can define the printer at runtime.

To avoid collision, its preferable to configure it into `main`

The default printer is `LogPrinter` a wrapper of `go/log` package.

```go
package main

import (
  "fmt"
  "github.com/mh-cbon/verbose"
  "github.com/mh-cbon/verbose/printer"
  "test"
)

var logger = verbose.Auto()

func main () {

  verbose.SetPrinter(printer.FmtPrinter{}) // changed to a wrapper of go/fmt

  fmt.Println("Hello from main")
  logger.Println("message to log from main github.com/mh-cbon/verbose")
  a.Hello()
}
```

#### printer.Printer interface

```go
type Printer interface {
	Printf(logger *Logger, format string, a ...interface{})
	Print(logger *Logger, v ...interface{})
	Println(logger *Logger, v ...interface{})
}
```
