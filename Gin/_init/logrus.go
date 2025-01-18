package logr

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"path/filepath"
)

// CustomHook是自定义的Logrus Hook,用于根据日志级别输出到不同目录
type CustomHook struct {
	//为了避免每次写入时重复打开文件，定义文件句柄
	fileHandles map[logrus.Level]*os.File
}

func (hook *CustomHook) Levels() []logrus.Level {
	//这里指定这个hook会处理哪些级别的日志
	return []logrus.Level{
		logrus.InfoLevel,
		logrus.ErrorLevel,
		logrus.DebugLevel,
		logrus.WarnLevel,
	}
}

func (hook *CustomHook) Fire(entry *logrus.Entry) error {
	if hook.fileHandles == nil {
		hook.fileHandles = make(map[logrus.Level]*os.File)
	}
	//根据日志级别决定输出的目录文件
	var logDir string
	switch entry.Level {
	case logrus.DebugLevel:
		logDir = "../Logs/Debug/"
	case logrus.InfoLevel:
		logDir = "../Logs/Info/"
	case logrus.ErrorLevel:
		logDir = "../Logs/Error/"
	case logrus.WarnLevel:
		logDir = "../Logs/Warn/"
	default:
		logDir = "../Logs/other"
	}
	//确保目标目录存在
	if err := os.MkdirAll(logDir, os.ModePerm); err != nil {
		return fmt.Errorf(err.Error())
	}
	//根据级别生成日志文件的完整路径
	logfile := filepath.Join(logDir, "logfile.txt")
	//如果该级别的文件句柄已打开，直接复用，否则打开新文件
	file, exist := hook.fileHandles[entry.Level]
	if !exist {
		//打开或创建
		var err error
		file, err = os.OpenFile(logfile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			return fmt.Errorf(err.Error())
		}
		//存储文件句柄
		hook.fileHandles[entry.Level] = file
	}

	//将日志输出到文件
	entry.Logger.SetOutput(file)
	return nil
}
