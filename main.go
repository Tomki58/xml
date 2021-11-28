package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"os"

	"myxml/logic"
	mxml "myxml/models/xml"
)

func readXml(path string) ([]byte, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return data, err
}

func main() {
	data, err := readXml("./instances/AttrWithValues.xml")
	if err != nil {
		log.Fatal(err)
	}

	var instances mxml.Instances
	if err := xml.Unmarshal(data, &instances); err != nil {
		log.Fatal(err)
	}

	query, err := logic.CreateInsertRequest(instances.Instances)
	if err != nil {
		log.Fatal(err)
	}

	repr, err := query.ToDDL()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(repr)
}
