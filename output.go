package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"os"
	"strconv"

	"github.com/Masterminds/sprig/v3"
	"github.com/miaogaolin/condition"
	"github.com/miaogaolin/excel-proc/config"
	"github.com/miaogaolin/excel-proc/utils"
	"github.com/pkg/errors"
)

// Ouput result  to file
func Ouput(orders [][]interface{}, config *config.Config, output string) error {
	if config == nil {
		return errors.New("config file is empty")
	}

	var ignore int
	res := bytes.NewBuffer(nil)

	var maxColumn int
	for _, v := range orders {
		if len(v) > maxColumn {
			maxColumn = len(v)
		}
	}

	for _, v := range orders {
		if len(v) != maxColumn {
			ignore++
			continue
		}
		data := map[string]interface{}{}
		for i, col := range v {
			colName := "col" + strconv.FormatInt(int64(i+1), 10)
			data[colName] = col
		}

		// render fields data
		for k, v := range config.Fields {
			fieldRes := bytes.NewBuffer(nil)
			tpl := template.Must(
				template.New("fields template").Funcs(sprig.FuncMap()).Parse(v))
			err := tpl.Execute(fieldRes, data)
			if err != nil {
				return errors.Wrapf(err, "template excute, template=%v, data=%v", v, data)
			}

			data[k] = template.HTML(fieldRes.String())
		}

		// 验证每条数据和配置文件中哪个条件匹配
		var isSuccess bool
		for _, c := range config.Sections {
			ok, err := validateConditon(data, c.Condition)
			if err != nil {
				if errors.Is(err, condition.ErrNotFoundCol) ||
					errors.Is(err, condition.ErrParseData) {
					d, _ := json.Marshal(data)
					fmt.Printf("[Warning] config file, line:%v, %v\n condition: %v\n data: %v\n", c.ConditionLine, err, c.Condition, string(d))
					break
				}
				return fmt.Errorf(`[Error] config file, line:%v, %v, condition=%v`, c.ConditionLine, err, c.Condition)
			}
			if ok {
				// 渲染模版数据
				tpl := template.Must(
					template.New("base").Funcs(sprig.FuncMap()).Parse(c.Template))
				err := tpl.Execute(res, data)
				if err != nil {
					return errors.Wrapf(err, "template excute, template=%v, data=%v", c.Template, data)
				}
				isSuccess = true
				res.Write([]byte(utils.LineSeperator()))
				break
			}
		}
		if !isSuccess {
			ignore++
		}
	}

	fmt.Printf(`-------------------
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
		return false, err
	}
	return r, nil
}
