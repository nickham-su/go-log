package main

import (
	"github.com/nickham-su/go_logger"
	"os"
)

func main() {
	logger.AppendWriter(os.Stdout)
	logger.SetDir("logs")
	logger.Debug.Println("debug")
	logger.Info.Println("info")
	logger.Warning.Println("warning")
	logger.Error.Println("error")
	logger.Debug.Printf("debug %v", 1)
	logger.Info.Printf("info %v", 2)
	logger.Warning.Printf("warning %v", 3)
	logger.Error.Printf("error %v", 4)
	//logger.Error.Fatalln("Fatalln")
	logger.Error.Fatalf("Fatalf %v", "exit")
}
