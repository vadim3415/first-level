package json_service

import (
	"fmt"
	"strconv"
)

// мой сервис

type JsonDocument struct {
	JsonData int
}

func (doc JsonDocument) ConvertToXml() string {
	convert := strconv.Itoa(doc.JsonData)
	return convert
}
///////адаптр для работы с внешним сервисом//////////////
type JsonDocumentAdapter struct {
	JsonDocument *JsonDocument
}

func (adapter JsonDocumentAdapter) SendXmlData() {
	xmlDoc:= adapter.JsonDocument.ConvertToXml()
	fmt.Println("отправка Xml данных", xmlDoc)
}