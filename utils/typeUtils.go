package utils

func BoolToInt(flag bool) int {
	if flag {
		return 1
	} else {
		return 0
	}
}

func IntToBool(flag int) bool {
	if flag == 0 {
		return false
	} else {
		return true
	}
}
