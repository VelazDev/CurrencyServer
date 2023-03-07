package utils

import "strconv"

func ConvertTofloat(str string, bit int) float64 {
	tmp, err := strconv.ParseFloat(str, bit)
	if err != nil {
		return 0
	}
	return tmp
}
