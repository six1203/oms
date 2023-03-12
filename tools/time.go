package tools

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

func GetNowTime() time.Time {
	return time.Now()
}

func GetNowTimeAddMinute(minutes int8) time.Time {
	return time.Now().Add(time.Duration(minutes) * time.Minute)
}

func GetUnixEpoch() time.Time {
	return time.Unix(0, 0)
}

// TimeToTimestamp time.Time 转换为 protobuf的Timestamp
func TimeToTimestamp(time time.Time) *timestamppb.Timestamp {
	return timestamppb.New(time)
}
