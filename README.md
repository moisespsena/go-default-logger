# go-default-logger
My Default Logger for GoLang projects

## Installation

```bash
go get -u github.com/moisespsena-go/default-logger
```

## Usage

```go
package main

import "github.com/moisespsena-go/default-logger"

var log = defaultlogger.NewLogger("main")

func main() {
	log.Info("Message!")
}
```

## Change Format

```go
package pkg

import (
	"os"
	
	"github.com/moisespsena-go/default-logger"
	"github.com/op/go-logging"
)

func init() {
	defaultlogger.Format = logging.MustStringFormatter(
		`%{color}%{time:15:04:05.000} %{module} %{shortfunc} -> %{level:.4s} %{id:03x}%{color:reset} %{message}`)
}
```

## Change Backend

```go
package pkg

import (
	"os"
	
	"github.com/moisespsena-go/default-logger"
	"github.com/op/go-logging"
)

func init() {
	defaultlogger.Backend = logging.NewLogBackend(os.Stderr, "", 0)
}
```

## Thank You!

By [Moises P. Sena](https://github.com/moisespsena)