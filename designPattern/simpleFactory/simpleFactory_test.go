package simpleFactory

import "testing"

func TestSimpleFactory(t *testing.T) {
	loggerFactory := new(LoggerFactory)
	sysLog, err := loggerFactory.GetLogger("sysLog")
	if err != nil {
		t.Errorf("get SysLog error: %s", err.Error())
	}
	if sysLog.Name() != "sysLog" {
		t.Errorf("get SysLog name error: %s", sysLog.Name())
	}

	fileLog, err := loggerFactory.GetLogger("fileLog")

	if err != nil {
		t.Errorf("get FileLog error: %s", err.Error())
	}

	if fileLog.Name() != "fileLog" {
		t.Errorf("get FileLog name error: %s", fileLog.Name())
	}
}
