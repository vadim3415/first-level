package main

import (
	"firstLevel/develop/dev21/json_service"
	"firstLevel/develop/dev21/xml_service"
)

// внешний сервис, который работает с форматом xml. (string)
//  Все документы, которые передаются и принимаются в xml формате
// у меня все данные в формате json (int)

func main() {
	myDocXml := xml_service.XmlDocument{XmlData: "hello"}
	myDocXml.SendXmlData() 

	myDocJson:= json_service.JsonDocument{JsonData: 123}
	convertMyDocJson:= json_service.JsonDocumentAdapter{JsonDocument: &myDocJson}
	convertMyDocJson.SendXmlData()
}