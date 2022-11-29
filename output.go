package main

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
	"strconv"
	"strings"

	"github.com/Masterminds/sprig"
	"github.com/zhouzhuojie/conditions"
)

// Ouput result  to file
func Ouput(orders [][]interface{}, configs []Config, output string) error {
	res := bytes.NewBuffer(nil)
	for _, v := range orders {
		data := map[string]interface{}{}
		for i, col := range v {
			colName := "col" + strconv.FormatInt(int64(i+1), 10)
			data[colName] = col
		}
		// 验证每条数据和配置文件中哪个条件匹配
		for _, c := range configs {
			ok, err := validateConditon(data, c.Condition)
			if err != nil {
				fmt.Printf("[Warning] %v, data = %v, condition = %v\n", err, data, c.Condition)
				continue
			}
			if ok {
				// 渲染模版数据
				tpl := template.Must(
					template.New("base").Funcs(sprig.FuncMap()).Parse(c.Template))
				err := tpl.Execute(res, data)
				if err != nil {
					fmt.Printf("template excute, %v", err)
				}
				break
			}
		}

	}

	return os.WriteFile(output, res.Bytes(), 0777)
}

func validateConditon(data map[string]interface{}, condition string) (bool, error) {
	// 没有条件代表，代表所有匹配
	if condition == "" {
		return true, nil
	}
	// Parse the condition language and get expression
	p := conditions.NewParser(strings.NewReader(condition))
	expr, err := p.Parse()
	if err != nil {
		fmt.Printf("condition=%s,%v\n", condition, err)
		return false, err
	}

	// Evaluate expression passing data for $vars
	r, err := conditions.Evaluate(expr, data)
	if err != nil {
		return false, err
	}
	return r, nil
}
