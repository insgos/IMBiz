package logger

import (
	"io"
	"log"
	"os"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
)

func SetLogger(path string) {
	rl, _ := rotatelogs.New(path + "/log.%Y-%m-%d")
	log.SetOutput(io.MultiWriter(rl, os.Stdout))
}
