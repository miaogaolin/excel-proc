package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type ConfigSection struct {
	Template  string
	Condition string
}
type Config struct {
	Fields   map[string]string
	Sections []ConfigSection
}

func ReadConfig(path string) (*Config, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var (
		section        []string
		configSections []ConfigSection
		line           int
	)

	fields := map[string]string{}
	for {
		if !scanner.Scan() {
			// 跳出前把 section 中的配置数据解析
			if len(section) > 0 {
				tmp, err := getSectionConfig(section)
				if err != nil {
					return nil, fmt.Errorf("line num %d, %v", line, err)
				}
				configSections = append(configSections, *tmp)
			}
			break
		}
		line++
		text := scanner.Text()
		trimText := strings.TrimSpace(text)
		// read fields
		if len(section) == 0 && trimText != "" {
			reg := regexp.MustCompile(`(\w+):\s*(?:"([^"]*)"|'([^']*)'|(\S+))`)
			fieldsText := reg.FindSubmatch([]byte(trimText))
			if len(fieldsText) == 5 {
				fields[string(fieldsText[1])] = string(fieldsText[2]) + string(fieldsText[3]) + string(fieldsText[4])
				continue
			}

		}

		// 按照空行划分
		if trimText == "" && len(section) > 0 {
			tmp, err := getSectionConfig(section)
			if err != nil {
				return nil, fmt.Errorf("line num %d, %v", line, err)
			}
			configSections = append(configSections, *tmp)
			section = nil
		} else if trimText != "" {
			section = append(section, text)
		}
	}

	return &Config{
		Sections: configSections,
		Fields:   fields,
	}, nil
}

func getSectionConfig(section []string) (*ConfigSection, error) {
	if len(section) == 0 {
		return nil, errors.New("config is empty")
	}

	var cfg ConfigSection
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
