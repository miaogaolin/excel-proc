package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type Config struct {
	Condition string
	Template  string
}

func ReadConfig(path string) ([]Config, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var (
		section []string
		res     []Config
		line    int
	)

	for {
		if !scanner.Scan() {
			// 跳出前把 section 中的配置数据解析
			if len(section) > 0 {
				tmp, err := getSectionConfig(section)
				if err != nil {
					return nil, fmt.Errorf("line num %d, %v", line, err)
				}
				res = append(res, *tmp)
			}
			break
		}

		line++
		text := scanner.Text()
		trimText := strings.TrimSpace(text)
		// 按照空行划分
		if trimText == "" && len(section) > 0 {
			tmp, err := getSectionConfig(section)
			if err != nil {
				return nil, fmt.Errorf("line num %d, %v", line, err)
			}
			res = append(res, *tmp)
			section = nil
		} else if trimText != "" {
			section = append(section, text)
		}
	}
	return res, nil
}

func getSectionConfig(section []string) (*Config, error) {
	if len(section) == 0 {
		return nil, errors.New("config is empty")
	}

	var cfg Config
	// have condition staement
	if strings.HasPrefix(section[0], ";") {
		cfg.Condition = strings.TrimLeft(section[0], ";")
		if len(section) == 1 {
			return nil, errors.New("template is empty")
		}
		section = section[1:]
	}
	for _, v := range section {
		cfg.Template += v + LineSeperator()
	}
	return &cfg, nil
}
