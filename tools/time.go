package tools

import "time"

func GetNowTime() time.Time {
	return time.Now()
}

func GetNowTimeAddMinute(minutes int8) time.Time {
	return time.Now().Add(time.Duration(minutes) * time.Minute)
}

func GetUnixEpoch() time.Time {
	return time.Unix(0, 0)
}
