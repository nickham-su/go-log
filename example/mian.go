package main

import "github.com/nickham-su/go-logger"

func main() {
	logger.SetDir("logs")
	logger.Info.Println("info")
	logger.Warning.Println("warning")
	logger.Error.Println("error")
}
