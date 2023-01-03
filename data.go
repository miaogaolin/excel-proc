package main

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"io"
	"os"
	"path/filepath"

	"github.com/miaogaolin/excel-proc/utils"
	"github.com/pkg/errors"
	"github.com/xuri/excelize/v2"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
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
		b, err := io.ReadAll(readFile)
		if err != nil {
			return nil, err
		}

		e, _, _, err := determineEncodingUtf8OrGBK(bytes.NewReader(b))
		if err != nil {
			return nil, err
		}

		r := transform.NewReader(bytes.NewBuffer(b), e.NewDecoder())

		reader := csv.NewReader(r)
		reader.FieldsPerRecord = -1
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

func determineEncodingUtf8OrGBK(r io.Reader) (e encoding.Encoding, name string, certain bool, err error) {
	b, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		return
	}

	e, name, certain = charset.DetermineEncoding(b, "")
	if name != "utf-8" {
		e = simplifiedchinese.GBK
		name = "gbk"
	}
	return
}
