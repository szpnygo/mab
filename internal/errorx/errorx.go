package errorx

import (
	"errors"
	"fmt"

	"github.com/szpnygo/mab/internal/logx"
)

func Log(err error) {
	logx.Print("error: " + err.Error())
}

func Error(err error) error {
	logx.Print("error: " + err.Error())
	return err
}

func New(text string) error {
	logx.Print("error: " + text)
	return errors.New(text)
}

func Errorf(format string, a ...interface{}) error {
	text := fmt.Sprintf(format, a...)
	logx.Print("error: " + text)
	return errors.New(text)
}
