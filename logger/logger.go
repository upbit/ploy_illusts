package logger

import (
	"fmt"
	"path"
	"runtime"
)

// Logger Define which logger system can use
type Logger interface {
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
}

var log Logger = &StdLogger{}

// SetLogger Set logger backend
func SetLogger(l Logger) {
	log = l
}

// Debugf Call Debugf
func Debugf(format string, args ...interface{}) {
	if log != nil {
		log.Debugf(format, args...)
	}
}

// Infof Call Infof
func Infof(format string, args ...interface{}) {
	if log != nil {
		log.Infof(format, args...)
	}
}

// Warnf Call Warnf
func Warnf(format string, args ...interface{}) {
	if log != nil {
		log.Warnf(format, args...)
	}
}

// Errorf Call Errorf
func Errorf(format string, args ...interface{}) {
	if log != nil {
		log.Errorf(format, args...)
	}
}

// StdLogger logger to stdout
type StdLogger struct{}

// GetCaller 获取调用位置信息
func GetCaller() string {
	_, file, line, _ := runtime.Caller(1)
	return fmt.Sprintf("%v:%v", path.Base(file), line)
}

// Debugf Debugf
func (l *StdLogger) Debugf(format string, args ...interface{}) {
	fmt.Printf(fmt.Sprintf("%s [DEBUG] ", GetCaller())+format+"\n", args...)
}

// Infof Infof
func (l *StdLogger) Infof(format string, args ...interface{}) {
	fmt.Printf(fmt.Sprintf("%s [INFO] ", GetCaller())+format+"\n", args...)
}

// Warnf Warnf
func (l *StdLogger) Warnf(format string, args ...interface{}) {
	fmt.Printf(fmt.Sprintf("%s [WARN] ", GetCaller())+format+"\n", args...)
}

// Errorf Errorf
func (l *StdLogger) Errorf(format string, args ...interface{}) {
	fmt.Printf(fmt.Sprintf("%s [ERROR] ", GetCaller())+format+"\n", args...)
}
