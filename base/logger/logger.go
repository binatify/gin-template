package logger

import (
	"github.com/sirupsen/logrus"
)

var (
	ReqIDFieldName = "ReqID"
)

func NewLogger(cfg *Config) (*logrus.Logger, error) {
	level, err := logrus.ParseLevel(cfg.Level)
	if err != nil {
		return nil, err
	}

	logger := logrus.New()
	logger.Level = level

	// set output
	out, err := cfg.GetOutputWriter()
	if err != nil {
		return nil, err
	}

	logger.SetOutput(out)

	// set fomatter
	logger.SetFormatter(cfg.GetFormater())

	// set calller printer
	logger.SetReportCaller(!cfg.DisableCaller)

	return logger, nil
}

func NewAppLogger(logger *logrus.Logger, reqId string) *logrus.Entry {
	return logger.WithFields(logrus.Fields{
		ReqIDFieldName: reqId,
	})
}

func GetReqId(entry *logrus.Entry) string {
	id, ok := entry.Data[ReqIDFieldName].(string)
	if ok {
		return id
	}

	return ""
}
