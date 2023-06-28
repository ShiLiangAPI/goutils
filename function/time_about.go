package function

import "time"

func GetTimeStr() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// GetTimeUnix 获取当前时间戳
//func GetTimeUnix() int64 {
//	return time.Now().Unix()
//}
