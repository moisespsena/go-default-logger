package defaultlogger

import (
	"os"

	"github.com/op/go-logging"
)

var Format = logging.MustStringFormatter(
	`%{color}%{time:15:04:05.000} %{level:.4s} [%{module}] %{shortfunc}: %{id:03x}%{color:reset} %{message}`,
)

var Backend = logging.NewLogBackend(os.Stderr, "", 0)

var backendFormatter logging.Backend

func Init() {
	backendFormatter := logging.NewBackendFormatter(Backend, Format)
	logging.SetBackend(backendFormatter)
}

func NewLogger(module string) *logging.Logger {
	if backendFormatter == nil {
		Init()
	}

	return logging.MustGetLogger(module)
}
