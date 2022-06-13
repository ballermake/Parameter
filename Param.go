package gage

import "strings"

func CheckNum(nums ...int) string {
	var result = make([]string, len(nums))

	for index, num := range nums {

		if num%2 == 0 {
			result[index] = "GENAP"
		} else {
			result[index] = "GANJIL"
		}
	}
	return strings.Join(result, ", ")
}
