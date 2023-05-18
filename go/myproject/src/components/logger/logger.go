package logger

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/sirupsen/logrus"
)

type Config struct {
	Level  string
	Output string
	Format string
}

type Log struct {
	log    *logrus.Logger
	config Config
}

func New(config Config) Log {
	// log = logrus.NewEntry(logrus.New())
	log := logrus.New()
	out := config.Output
	if out == "" || out == "stdout" {
		log.SetOutput(os.Stdout)
	} else {
		f, err := os.OpenFile(out, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)
		if err != nil {
			panic(err)
		}
		log.SetOutput(f)
	}

	var formatter logrus.Formatter
	format := config.Format
	if format == "" || format == "text" {
		formatter = &logrus.TextFormatter{}
	}

	if format == "json" {
		formatter = &logrus.JSONFormatter{
			FieldMap: logrus.FieldMap{
				logrus.FieldKeyTime:  "time",
				logrus.FieldKeyLevel: "level",
				logrus.FieldKeyMsg:   "message",
				logrus.FieldKeyFunc:  "caller",
			},
		}
	}

	level := getLogLevel(config)
	log.SetLevel(level)

	log.SetFormatter(formatter)
	return Log{
		log:    log,
		config: config,
	}
}

func (l Log) Error(v ...interface{}) {
	message := l.toString(v)
	l.log.WithField("caller", getCaller()).Error(message)
}

func (l Log) Info(v ...interface{}) {
	message := l.toString(v)
	l.log.Info(message)
}

func (l Log) Debug(v ...interface{}) {
	message := l.toString(v)
	l.log.Debug(message)
}

func (l Log) Fatal(v ...interface{}) {
	message := l.toString(v)
	l.log.WithField("func", getCaller()).Fatal(message)
}

func getCaller() string {
	pc, _, _, ok := runtime.Caller(2)
	details := runtime.FuncForPC(pc)
	if ok && details != nil {
		return details.Name()
	}
	return ""
}

func (l Log) toString(v ...interface{}) string {
	var buf bytes.Buffer
	for _, s := range v {
		buf.WriteString(fmt.Sprintf("%+v", s))
	}
	value := strings.Replace(buf.String(), "[", "", 1)
	value = strings.Replace(value, "]", "", 1)

	level := getLogLevel(l.config)
	if level.String() == "debug" {
		value = concatDebugInfo(value)
	}

	return value
}

func concatDebugInfo(value string) string {
	pc, file, line, ok := runtime.Caller(3)
	if !ok {
		file = "?"
		line = 0
	}

	fn := runtime.FuncForPC(pc)
	var fnName string
	if fn == nil {
		fnName = "?"
	} else {
		dotName := filepath.Ext(fn.Name())
		fnName = strings.TrimLeft(dotName, ".")
	}

	return fmt.Sprintf("%s:%d %s: %s", filepath.Base(file), line, fnName, value)
}

func getLogLevel(config Config) logrus.Level {
	level := config.Level
	if lvl, err := logrus.ParseLevel(level); err == nil {
		return lvl
	}
	return logrus.ErrorLevel
}
