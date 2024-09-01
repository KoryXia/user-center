package initialize

import (
	"bytes"
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	log "github.com/sirupsen/logrus"
	"path"
	"strings"
	"time"
)

type MyFormatter struct {
	log.TextFormatter
}

func (c *MyFormatter) Format(entry *log.Entry) ([]byte, error) {
	var logBuffer bytes.Buffer
	if entry.Buffer != nil {
		logBuffer = *entry.Buffer
	}

	logBuffer.WriteString(
		fmt.Sprintf(
			"%s [%s]",
			strings.ToUpper(entry.Level.String()),
			entry.Time.Format(c.TimestampFormat),
		),
	)
	if entry.HasCaller() {
		logBuffer.WriteString(fmt.Sprintf("[%s:%d]", path.Base(entry.Caller.File), entry.Caller.Line))
	}
	logBuffer.WriteString(fmt.Sprintf(" %s\n", entry.Message))
	return logBuffer.Bytes(), nil
}

func Log() {
	output, err := rotatelogs.New(
		"./log/server.log"+".%Y%m%d",
		rotatelogs.WithRotationCount(3),
		rotatelogs.WithRotationTime(time.Hour*24),
	)
	if err != nil {
		panic(err.Error())
	}
	log.SetFormatter(&MyFormatter{
		log.TextFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
			FullTimestamp:   true,
		},
	})
	log.SetReportCaller(true)
	log.SetOutput(output)
	log.SetLevel(log.InfoLevel)
}
