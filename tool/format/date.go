package format

import "time"

// 时间字符串 -> 时间戳
func StringToTimestamp(param string) int64 {
	var retTimestamp int64
	if timeData, err := time.Parse("2006-01-02", param); err == nil {
		retTimestamp = timeData.Unix()
	} else {
		panic(err)
	}
	return retTimestamp
}

// 时间戳 -> 时间字符串
func TimestampToString(param int64) string {
	var retString string
	timeData := time.Unix(param, 0)
	retString = timeData.Format("2006-01-02")
	return retString
}

// 字符串 -> 时间
func StringToTime(param string) time.Time {
	var retTime time.Time

	if timeData, err := time.Parse("2006-01-02", param); err == nil {
		retTime = timeData
	} else {
		panic(err)
	}
	return retTime
}

// 时间 -> 字符串
func TimeToString(param time.Time) string {
	var retString string
	retString = param.Format("2006-01-02")
	return retString
}
