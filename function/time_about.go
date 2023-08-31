package function

import (
	"time"
)

func GetCurrentTimeStr() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// GetCurrentTimeUnix 获取当前时间戳
func GetCurrentTimeUnix() int64 {
	return time.Now().Unix()
}

func ParseDateStrToLocTime(timeStr string) time.Time {
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		return time.Time{}
	}

	timeObj, err := time.ParseInLocation("2006-01-02", timeStr, loc)
	if err != nil {
		return time.Time{}
	}

	return timeObj
}
