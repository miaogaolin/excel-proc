package main

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
	"strconv"

	"github.com/Masterminds/sprig"
	"github.com/miaogaolin/excel-proc/condition"
	"github.com/pkg/errors"
)

// Ouput result  to file
func Ouput(orders [][]interface{}, configs []Config, output string) error {
	var ignore int
	res := bytes.NewBuffer(nil)
	for _, v := range orders {
		data := map[string]interface{}{}
		for i, col := range v {
			colName := "col" + strconv.FormatInt(int64(i+1), 10)
			data[colName] = col
		}
		// 验证每条数据和配置文件中哪个条件匹配
		var isSuccess bool
		for _, c := range configs {
			ok, err := validateConditon(data, c.Condition)
			if err != nil {
				if errors.Is(err, condition.ErrNotFoundCol) ||
					errors.Is(err, condition.ErrParseData) {
					fmt.Printf(`[Warning] %v, condition="%v", data=%v`, err, c.Condition, v)
					break
				}
				return fmt.Errorf(`[Error] %v, condition="%v"`, err, c.Condition)
			}
			if ok {
				// 渲染模版数据
				tpl := template.Must(
					template.New("base").Funcs(sprig.FuncMap()).Parse(c.Template))
				err := tpl.Execute(res, data)
				if err != nil {
					return errors.Wrap(err, "template excute")
				}
				isSuccess = true
				break
			}
		}

		if !isSuccess {
			ignore++
		}
	}

	fmt.Printf(`
total count: %d
ignore count: %d
success count: %d
	`, len(orders), ignore, len(orders)-ignore)
	return os.WriteFile(output, res.Bytes(), 0777)
}

func validateConditon(data map[string]interface{}, conditionExpr string) (bool, error) {
	// 没有条件代表，代表所有匹配
	if conditionExpr == "" {
		return true, nil
	}

	// Evaluate expression passing data for $vars
	r, err := condition.Validate(data, conditionExpr)
	if err != nil {
		fmt.Println(data, "\n", conditionExpr)
		return false, err
	}
	return r, nil
}
