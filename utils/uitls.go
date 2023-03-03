package utils

import (
	"errors"
	"math"
	"regexp"
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

func GetMoneyNum(money string) (float64, error) {
	money = strings.Trim(money, `"`)

	var sign float64 = 1
	if len(money) >= 1 && money[0] == '-' {
		sign = -1
		money = money[1:]
	}

	// 匹配人民币、美元等等格式：￥123.45
	reCurrency := regexp.MustCompile(`^(\$|¥|€|£)(((\d{1,3}(,\d{3})*)|\d{4,})(\.\d+)?)$`)
	// 匹配普通数字格式：123.45 或 1,234.56
	reNum := regexp.MustCompile(`^((\d{1,3}(,\d{3})*)|\d{4,})(\.\d{1,2})?$`)

	// 按优先级尝试匹配不同的金额格式
	if reCurrency.MatchString(money) {
		m := []rune(money)
		money = string(m[1:])
	}

	if reNum.MatchString(money) {
		numStr := strings.ReplaceAll(money, ",", "") // 去掉逗号分隔符
		num, err := strconv.ParseFloat(numStr, 64)
		if err != nil {
			return 0, err
		}
		return math.Round(num*sign*100) / 100, nil
	}

	// 如果没有匹配成功，则返回错误
	return 0, errors.New("invalid money format")
}
