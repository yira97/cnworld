package format

import "strconv"

func ToInt64(s string, defaultValue int64) int64 {
	n, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return defaultValue
	}
	return n
}

func ToInt(s string, dafaultValue int) int {
	return int(ToInt64(s, int64(dafaultValue)))
}
