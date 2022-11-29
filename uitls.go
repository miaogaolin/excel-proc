package main

import (
	"runtime"
	"strconv"
	"strings"
)

func LineSeperator() string {
	if runtime.GOOS == "windows" {
		return "\r\n"
	}
	return "\n"
}

func InSliceString(ele string, slices []string) bool {
	for _, v := range slices {
		ele = strings.ToUpper(ele)
		v = strings.ToUpper(v)
		if ele == v {
			return true
		}
	}
	return false
}

func ConvertNum(txt string) (interface{}, bool) {
	txt = strings.Replace(txt, ",", "", -1)
	txt = strings.Trim(txt, `"`)
	// int64
	if !strings.Contains(txt, ".") {
		num, err := strconv.ParseInt(txt, 10, 64)
		if err != nil {
			return 0, false
		}
		return num, true
	}

	// float64
	num, err := strconv.ParseFloat(txt, 64)
	if err != nil {
		return 0, false
	}
	return num, true
}
