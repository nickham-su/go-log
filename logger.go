package logger

import (
	"io"
	"log"
	"os"
	"time"
)

var (
	Debug   *logger
	Info    *logger
	Warning *logger
	Error   *logger
	dateStr string
	dirPath string
)

func init() {
	createLogger()
	ticker := time.NewTicker(time.Second)
	go func() {
		for t := range ticker.C {
			date := t.Format("2006-01-02")
			if dateStr != date {
				createLogger()
			}
		}
	}()
}

func createLogger() {
	dateStr = time.Now().Format("2006-01-02")
	Debug = newLogger(dirPath + dateStr + ".debug.log")
	Info = newLogger(dirPath + dateStr + ".info.log")
	Warning = newLogger(dirPath + dateStr + ".warning.log")
	Error = newLogger(dirPath + dateStr + ".error.log")
}

func SetDir(path string) {
	if path == "" {
		return
	}
	if path[len(path)-1:] != "/" {
		path += "/"
	}
	os.Mkdir(path, os.ModePerm)
	dirPath = path
	createLogger()
}

func newLogger(fileName string) *logger {
	return &logger{
		logger:   nil,
		fileName: fileName,
	}
}

type logger struct {
	logger   *log.Logger
	fileName string
}

func (l *logger) Println(v ...interface{}) {
	if l.logger == nil {
		file, err := os.OpenFile(l.fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalln("打开日志文件失败：", err)
		}
		l.logger = log.New(io.MultiWriter(os.Stdout, file), "", log.Ldate|log.Lmicroseconds)
	}
	l.logger.Println(v...)
}

func (l *logger) Printf(format string, v ...interface{}) {
	if l.logger == nil {
		file, err := os.OpenFile(l.fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalln("打开日志文件失败：", err)
		}
		l.logger = log.New(io.MultiWriter(os.Stdout, file), "", log.Ldate|log.Lmicroseconds)
	}
	l.logger.Printf(format, v...)
}
