package logger

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/sirupsen/logrus"

	"{{ .ProjectName }}/src/shared/conf"
)

var log = logrus.NewEntry(logrus.New())

func Configure() {
	config := conf.Get()
	l := logrus.New()
	out := config.Log.Output
	if out == "" || out == "stdout" {
		l.SetOutput(os.Stdout)
	} else {
		f, err := os.OpenFile(out, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)
		if err != nil {
			panic(err)
		}
		l.SetOutput(f)
	}

	var formatter logrus.Formatter
	format := config.Log.Format
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

	level := getLogLevel()
	l.SetLevel(level)

	l.SetFormatter(formatter)
	log = logrus.NewEntry(l)
}

func Error(v ...interface{}) {
	message := toString(v)
	log.WithField("caller", getCaller()).Error(message)
}

func Info(v ...interface{}) {
	message := toString(v)
	log.Info(message)
}

func Debug(v ...interface{}) {
	message := toString(v)
	log.Debug(message)
}

func Fatal(v ...interface{}) {
	message := toString(v)
	log.WithField("func", getCaller()).Fatal(message)
}

func getCaller() string {
	pc, _, _, ok := runtime.Caller(2)
	details := runtime.FuncForPC(pc)
	if ok && details != nil {
		return details.Name()
	}
	return ""
}

func toString(v ...interface{}) string {
	var buf bytes.Buffer
	for _, s := range v {
		buf.WriteString(fmt.Sprintf("%+v", s))
	}
	value := strings.Replace(buf.String(), "[", "", 1)
	value = strings.Replace(value, "]", "", 1)

	level := getLogLevel()
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

func getLogLevel() logrus.Level {
	config := conf.Get()
	level := config.Log.Level
	if lvl, err := logrus.ParseLevel(level); err == nil {
		return lvl
	} else {
		return logrus.ErrorLevel
	}
}
