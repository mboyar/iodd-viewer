package main

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"mime/multipart"
)

type Column struct {
	TextId string
	Value  string
}

type Table struct {
	Rows []Column
}

//Parse IODD XML file
func parseIODDbyPath(filename string) IODevice {

	var data IODevice

	buff, err := ioutil.ReadFile(filename)

	if err != nil {
		log.Fatal(err)
	}

	err = xml.Unmarshal(buff, &data)

	if err != nil {
		log.Fatal(err)
	}

	return data
}

func parseIODDbyFile(file multipart.File) IODevice {

	var data IODevice

	buff, err := ioutil.ReadAll(file)

	if err != nil {
		log.Fatal(err)
	}

	err = xml.Unmarshal(buff, &data)

	if err != nil {
		log.Fatal(err)
	}

	return data
}

func generateHtmlData(file multipart.File) Table {

	var dataHtml Table
	xmlIODD := parseIODDbyFile(file)

	//println("XML IODD:")
	//println("\txmlIODD.ProfileBody.DeviceIdentity.VendorName:", xmlIODD.ProfileBody.DeviceIdentity.VendorName)
	//println("\txmlIODD.ProfileBody.DeviceIdentity.DeviceName.TextId:", xmlIODD.ProfileBody.DeviceIdentity.DeviceName.TextId)

	lenExtTextCollection := len(xmlIODD.ExternalTextCollection.PrimaryLanguage.Text)

	for i := 0; i < lenExtTextCollection; i++ {

		//if xmlIODD.ExternalTextCollection.PrimaryLanguage.Text[i].ID == xmlIODD.ProfileBody.DeviceIdentity.DeviceName.TextId {
		//	println("\txmlIODD.ExternalTextCollection.PrimaryLanguage.Text[", i, "].Value:", xmlIODD.ExternalTextCollection.PrimaryLanguage.Text[i].Value)
		//}

		dataHtml.Rows = append(dataHtml.Rows, Column{TextId: xmlIODD.ExternalTextCollection.PrimaryLanguage.Text[i].ID,
			Value: xmlIODD.ExternalTextCollection.PrimaryLanguage.Text[i].Value})
	}

	return dataHtml
}
