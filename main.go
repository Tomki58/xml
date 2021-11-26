package main

import (
	"encoding/xml"
	"log"
	"os"

	"myxml/logic"
	m "myxml/models/xml"
)

func readXml(path string) ([]byte, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return data, err
}

func main() {
	d, err := readXml("/home/argabidullin/nestorclient/instances/BaseKind.xml")
	if err != nil {
		log.Fatal(err)
	}

	var instances m.Instances
	xml.Unmarshal(d, &instances)

	query := logic.New(instances.Instances)
	log.Print(query)
}
