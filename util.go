package sfcc_export_parser

import (
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"os"
)

func readFile(filename string) ([]byte, error) {
	return ioutil.ReadFile(filename)
}

func ParseXML(filename string) ([]Order, error) {
	data, err := readFile(filename)

	if err != nil {
		return nil, err
	}

	var orders Orders

	err = xml.Unmarshal(data, &orders)

	if err != nil {
		return nil, err
	}

	return orders.Orders, nil
}

func XMLtoJSON(filename, exportFile string) error {
	orders, err := ParseXML(filename)

	if err != nil {
		return err
	}

	b, err := json.MarshalIndent(orders, "", "\t")

	if err != nil {
		return err
	}

	err = ioutil.WriteFile(exportFile, b, os.ModePerm)

	if err != nil {
		return err
	}

	return nil
}
