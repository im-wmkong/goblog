package types

import (
	"goblog/pkg/logger"
	"strconv"
)

func Uint64ToString(num uint64) string {
	return strconv.FormatUint(num, 10)
}

func StringToInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		logger.LogError(err)
	}
	return i
}

func IntToString(i int) string {
	return strconv.Itoa(i)
}
