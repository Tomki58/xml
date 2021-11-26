package models

import (
	"encoding/xml"

	"github.com/google/uuid"
)

type StdAttributes struct {
	UID       uuid.UUID `xml:"uid,attr"`
	Owner     uuid.UUID `xml:"owner,attr"`
	Via       string    `xml:"via,attr"`
	Many      bool      `xml:"isMany,attr"`
	Reversed  bool      `xml:"reversed,attr"`
	Reference bool      `xml:"reference,attr"`
	Required  bool      `xml:"required,attr"`
	Readonly  bool      `xml:"readonly,attr"`
}

type Instances struct {
	XMLName   xml.Name `xml:"instances"`
	Instances Instance `xml:"instance"`
}

type Instance struct {
	XMLName  xml.Name  `xml:"instance"`
	Domain   string    `xml:"domain,attr"`
	Name     string    `xml:"model,attr"`
	Generics []Generic `xml:"generic"`
}
