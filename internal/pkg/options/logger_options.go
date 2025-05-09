package options

import (
	"strings"

	"gin-study/pkg/logger"
)

type LoggerOption struct {
	Level   string   `json:"level"    mapstructure:"level"`
	File    string   `json:"file"     mapstructure:"file"`
	TraceID string   `json:"trace_id" mapstructure:"-"`
	Fields  []string `json:"fields"   mapstructure:"-"`
}

func (o *LoggerOption) NewClient() *logger.ZapLogger {
	// 初始化日志
	levelOpt := logger.WithInfoLevel()

	switch strings.ToLower(o.Level) {
	case "debug":
		levelOpt = logger.WithDebugLevel()
	case "warn":
		levelOpt = logger.WithWarnLevel()
	case "error":
		levelOpt = logger.WithErrorLevel()
	case "panic":
		levelOpt = logger.WithPanicLevel()
	case "fatal":
		levelOpt = logger.WithFatalLevel()
	}

	return logger.NewLogger(
		levelOpt,
		logger.WithFileRotationP(o.File),
		logger.WithCtxFields(o.Fields),
		logger.WithTraceID(o.TraceID),
	)
}
