package utils

import "strconv"

func GenerateContentId(count int) string {
	countStr := strconv.Itoa(count)
	for i := len(countStr); i < 10; i++ {
		countStr = "0" + countStr
	}
	return "ENT-CNT-" + countStr
}

func GenerateEntryParamId(count int) string {
	countStr := strconv.Itoa(count)
	for i := len(countStr); i < 10; i++ {
		countStr = "0" + countStr
	}
	return "ENT-PRM-" + countStr
}
