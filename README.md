# Verbose

A library to display verbose messages per package.

__wip__

__wip__

__wip__


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
