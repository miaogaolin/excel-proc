package main

import (
	"encoding/csv"
	"os"
	"path/filepath"

	"github.com/miaogaolin/excel-proc/utils"
	"github.com/pkg/errors"
	"github.com/xuri/excelize/v2"
)

var dataFileExt = []string{".csv", ".xlam", ".xlsm", ".xlsx", ".xltm ", ".ltx"}

func GetData(filename string) ([][]string, error) {
	// verfiy excel file
	ext := filepath.Ext(filename)
	if !utils.InSliceString(ext, dataFileExt) {
		return nil, errors.Errorf("no excel file, supported file ext %v", dataFileExt)
	}

	readFile, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer readFile.Close()

	var data [][]string
	if ext == ".csv" {
		reader := csv.NewReader(readFile)

		data, err = reader.ReadAll()
		if err != nil {
			return nil, err
		}
	} else {
		f, err := excelize.OpenReader(readFile)
		if err != nil {
			return nil, err
		}

		// get first sheet
		data, err = f.GetRows(f.GetSheetName(0))
		if err != nil {
			return nil, err
		}
	}
	return data, nil
}
