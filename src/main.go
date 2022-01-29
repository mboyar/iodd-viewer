package main

func main() {

	var data Table
	xmlIODD := parseXml("../resources/sample-iodd-2.xml")

	//println("XML IODD:")
	//println("\txmlIODD.ProfileBody.DeviceIdentity.VendorName:", xmlIODD.ProfileBody.DeviceIdentity.VendorName)
	//println("\txmlIODD.ProfileBody.DeviceIdentity.DeviceName.TextId:", xmlIODD.ProfileBody.DeviceIdentity.DeviceName.TextId)

	lenExtTextCollection := len(xmlIODD.ExternalTextCollection.PrimaryLanguage.Text)

	for i := 0; i < lenExtTextCollection; i++ {

		//if xmlIODD.ExternalTextCollection.PrimaryLanguage.Text[i].ID == xmlIODD.ProfileBody.DeviceIdentity.DeviceName.TextId {
		//	println("\txmlIODD.ExternalTextCollection.PrimaryLanguage.Text[", i, "].Value:", xmlIODD.ExternalTextCollection.PrimaryLanguage.Text[i].Value)
		//}

		data.Rows = append(data.Rows, Column{TextId: xmlIODD.ExternalTextCollection.PrimaryLanguage.Text[i].ID,
			Value: xmlIODD.ExternalTextCollection.PrimaryLanguage.Text[i].Value})
	}

	generateHtml(data, "../html/template.html")
}
