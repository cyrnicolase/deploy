package models

import (
	"time"
)

// Date 日期格式 Y-m-d
type Date time.Time

// MarshalJSON 实现Marshaler 接口
func (d Date) MarshalJSON() ([]byte, error) {
	if time.Time(d).IsZero() {
		return []byte(`""`), nil
	}
	return []byte(`"` + time.Time(d).Format("2006-01-02") + `"`), nil
}

// DateTime 时间格式 Y-m-d H:i:s
type DateTime time.Time

// MarshalJSON 实现Marshaler 接口
func (datetime DateTime) MarshalJSON() ([]byte, error) {
	if time.Time(datetime).IsZero() {
		return []byte(`""`), nil
	}
	return []byte(`"` + time.Time(datetime).Format("2006-01-02 15:04:05") + `"`), nil
}
