package main

import (
	"github.com/nickham-su/go-logger"
	"os"
)

func main() {
	logger.AppendWriter(os.Stdout)
	logger.SetDir("logs")
	logger.Debug.Println("debug")
	logger.Info.Println("info")
	logger.Warning.Println("warning")
	logger.Error.Println("error")
}
