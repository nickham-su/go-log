package logger

import (
	"io"
	"log"
	"os"
	"time"
)

type logLevel int

const (
	levelDebug logLevel = iota
	levelInfo
	levelWarning
	levelError
)

var (
	Debug   *logger
	Info    *logger
	Warning *logger
	Error   *logger
	dateStr string
	dirPath string
	writers []io.Writer
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
	Debug = newLogger(levelDebug, dirPath+dateStr+".debug.log")
	Info = newLogger(levelInfo, dirPath+dateStr+".info.log")
	Warning = newLogger(levelWarning, dirPath+dateStr+".warning.log")
	Error = newLogger(levelError, dirPath+dateStr+".error.log")
}

func AppendWriter(writer ...io.Writer) {
	writers = append(writers, writer...)
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

func newLogger(level logLevel, fileName string) *logger {
	return &logger{
		logger:   nil,
		fileName: fileName,
		level:    level,
	}
}

type logger struct {
	logger   *log.Logger
	fileName string
	level    logLevel
}

func (l *logger) Println(v ...interface{}) {
	if l.logger == nil {
		file, err := os.OpenFile(l.fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalln("打开日志文件失败：", err)
		}
		w := append(writers, file)
		l.logger = log.New(io.MultiWriter(w...), "", log.Ldate|log.Lmicroseconds)
	}
	switch l.level {
	case levelDebug:
		v = append([]interface{}{"DEBUG"}, v...)
	case levelInfo:
		v = append([]interface{}{"INFO"}, v...)
	case levelWarning:
		v = append([]interface{}{"WARNING"}, v...)
	case levelError:
		v = append([]interface{}{"ERROR"}, v...)
	}
	l.logger.Println(v...)
}

func (l *logger) Printf(format string, v ...interface{}) {
	if l.logger == nil {
		file, err := os.OpenFile(l.fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalln("打开日志文件失败：", err)
		}
		w := append(writers, file)
		l.logger = log.New(io.MultiWriter(w...), "", log.Ldate|log.Lmicroseconds)
	}
	switch l.level {
	case levelDebug:
		format = "DEBUG " + format
	case levelInfo:
		format = "INFO " + format
	case levelWarning:
		format = "WARNING " + format
	case levelError:
		format = "ERROR " + format
	}
	l.logger.Printf(format, v...)
}
