package models

import (
	"encoding/xml"

	"github.com/google/uuid"
)

type Generic struct {
	XMLName xml.Name  `xml:"generic"`
	UID     uuid.UUID `xml:"uid,attr"`
	Etype   string    `xml:"etype,attr"`
	Owner   uuid.UUID `xml:"owner,attr"`
	From    string    `xml:"from,attr"`

	Attrs []Attr `xml:"attr"`
}
