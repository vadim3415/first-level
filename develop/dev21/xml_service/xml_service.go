package xml_service

import "fmt"

// внешний сервис

type AnalyticDataService interface{
	SendXmlData()
	}
	
	type XmlDocument struct{
		XmlData string
	}
	
	func (doc XmlDocument) SendXmlData(){
		fmt.Println("Отправка xml документа", doc.XmlData)
	}