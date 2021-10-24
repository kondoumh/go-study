package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type Bpmn struct {
	XMLName xml.Name `xml:"definitions"`
	Process BpmnProcess `xml:"process"`
}

type BpmnProcess struct {
	ID string `xml:"id,attr"`
	ProcessDefinition []byte `xml:",innerxml"`
}

type BpmnTask struct {
	ID string `xml:"id,attr"`
	Name string `xml:"name,attr"`
}

func main() {
	file, err := os.Open("diagram.bpmn")
	checkError(err, "File open")

	defer file.Close()
	data, err := ioutil.ReadAll(file)
	checkError(err, "Read")
	var bpmn Bpmn
	err = xml.Unmarshal(data, &bpmn)
	checkError(err, "Unmarshal")

	fmt.Println(bpmn)
}

func checkError(err error, mes string) {
	if err != nil {
		fmt.Printf("%s - error :%v\n", mes, err)
		os.Exit(1)
	}
}
