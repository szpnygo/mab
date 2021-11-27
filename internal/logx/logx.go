package logx

import (
	"log"

	"go.uber.org/atomic"
)

var (
	logEnabled = atomic.NewBool(true)
)

func SetLogEnabled(enabled bool) {
	logEnabled.Store(enabled)
}

func Print(msg string) {
	if logEnabled.Load() {
		log.Println("[MAB] ", msg)
	}
}
