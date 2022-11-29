package main

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

const version = "0.0.1"

var (
	configFile string
	outputFile string

	rootCmd = &cobra.Command{
		Use:   "excel-proc",
		Short: "excel-pro is a simple excel data pre-processing tool",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				log.Fatal("No data file path set")
			}
			err := Run(args[0])
			if err != nil {
				log.Fatal(err)
			}
		},
	}
)

func init() {
	rootCmd.Version = fmt.Sprintf("Version: %s", version)
	rootCmd.PersistentFlags().StringVar(&configFile, "config", "default.tpl", "Template configuration file")
	rootCmd.PersistentFlags().StringVar(&outputFile, "output", "default.bean", "File for saving results")
}

func main() {
	err := rootCmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}

func Run(dataFile string) error {
	configs, err := ReadConfig(configFile)
	if err != nil {
		return err
	}

	data, err := GetData(dataFile)
	if err != nil {
		return err
	}

	var SourceOrderData [][]interface{}
	for _, cols := range data {
		var rowData []interface{}
		for _, v := range cols {
			if num, ok := ConvertNum(v); ok {
				rowData = append(rowData, num)
			} else {
				rowData = append(rowData, v)
			}

		}

		SourceOrderData = append(SourceOrderData, rowData)
	}

	return Ouput(SourceOrderData, configs, outputFile)
}
