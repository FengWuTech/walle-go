package strutil

import (
	"github.com/unknwon/com"
	"strconv"
	"strings"
)

func IntSliceToStringSlice(ints []int) []string {
	var strSlice []string
	for _, v := range ints {
		strSlice = append(strSlice, strconv.Itoa(v))
	}
	return strSlice
}

func StringToIntSlice(str string, separator string) []int {
	var arrIntValues []int
	arrStrValues := strings.Split(str, separator)
	for _, v := range arrStrValues {
		arrIntValues = append(arrIntValues, com.StrTo(v).MustInt())
	}
	return arrIntValues
}

func IntSliceToString(ints []int, split string) string {
	var strs []string
	for _, v := range ints {
		strs = append(strs, strconv.Itoa(v))
	}
	return strings.Join(strs, split)
}
