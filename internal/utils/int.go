package utils

import "strconv"

func GetUint(s string, defaultValue int) int {

	a := defaultValue
	if s != "" {
		sInt, _ := strconv.Atoi(s)
		if sInt > 0 {
			a = sInt
		}
	}

	return a

}
