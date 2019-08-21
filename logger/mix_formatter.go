package logger

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

// 默认时间格式
var defaultTimetampFormat = "2006-01-02 15:04:05"

// MixFormatter 混合结构
type MixFormatter struct {
	// TimestampFormat sets the format used for marshaling timestamps
	TimestampFormat string

	// Filename sets the log filename
	Filename string
}

// Format 格式化
// 2019-08-20 15:18:21 INFO [message=] json
func (f *MixFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	timestampFormat := f.TimestampFormat
	if "" == timestampFormat {
		timestampFormat = defaultTimetampFormat
	}
	t := time.Now().Format(timestampFormat)
	message := entry.Message
	level := entry.Level.String()

	b := &bytes.Buffer{}
	encoder := json.NewEncoder(b)
	if err := encoder.Encode(entry.Data); nil != err {
		return nil, fmt.Errorf("fail to marshal fields to JSON, %v", err)
	}

	output := fmt.Sprintf("%s %s.%s: [%s] %s\n", t, f.Filename, level, message, strings.Trim(string(b.Bytes()), "\n"))

	return []byte(output), nil
}
