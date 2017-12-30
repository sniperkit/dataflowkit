package log

import (
	"path"
	"runtime"
	"strings"

	"github.com/sirupsen/logrus"
)

type ContextHook struct{}

//Levels return Levels of ContextHook
func (hook ContextHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

//Fire the hook.
func (hook ContextHook) Fire(entry *logrus.Entry) error {
	pc := make([]uintptr, 3, 3)
	cnt := runtime.Callers(6, pc)

	for i := 0; i < cnt; i++ {
		fu := runtime.FuncForPC(pc[i] - 1)
		name := fu.Name()
		if !strings.Contains(name, "github.com/sirupsen/logrus") {
			file, line := fu.FileLine(pc[i] - 1)
			entry.Data["file"] = path.Base(file)
			entry.Data["func"] = path.Base(name)
			entry.Data["line"] = line
			break
		}
	}
	return nil
}

//NewLogger creates New Logger instance.
func NewLogger() *logrus.Logger {
	logger := logrus.New()
	logger.AddHook(ContextHook{})
	return logger
}
