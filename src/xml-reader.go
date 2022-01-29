package main

import (
	"encoding/xml"
	"io/ioutil"
	"log"
)

//Parse XML file
func parseXml(filename string) IODevice {
	var data IODevice
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	err = xml.Unmarshal(file, &data)
	if err != nil {
		log.Fatal(err)
	}
	return data
}
