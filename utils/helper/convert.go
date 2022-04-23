package helper

import (
	"strconv"
	"time"
)

func BoolToInt(bValue bool) int {
	value := 0
	if bValue {
		value = 1
	}
	return value
}

func StringToInt64(sValue string, defValue int64) int64 {
	value := defValue
	iValue, err := strconv.ParseInt(sValue, 0, 64)
	if err == nil {
		value = iValue
	}
	return value
}

func PointDateToString(ptValue *time.Time, layout string) string {
	if ptValue == nil {
		return ""
	}
	tValue := *ptValue
	return tValue.Format(layout)
}

func StringToDate(sValue string, timelayout string) time.Time {
	var value time.Time
	t, err := time.Parse(timelayout, sValue)
	if err != nil {
		return value
	}
	return t
}
