package simpleFactory

import (
	"fmt"
)

type LoggerFactory struct {
}

func (l *LoggerFactory) GetLogger(name string) (Logger, error) {
	switch name {
	case "sysLog":
		return new(SysLog), nil
	case "fileLog":
		return new(FileLog), nil
	default:
		return nil, fmt.Errorf("logger %s not found", name)
	}
}

// 日志接口
type Logger interface {
	Name() string
}

// 系统日志
type SysLog struct {
}

func (s *SysLog) Name() string {
	return "sysLog"
}

// 文件日志
type FileLog struct {
}

func (f *FileLog) Name() string {
	return "fileLog"
}
