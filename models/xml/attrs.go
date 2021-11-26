package models

import "encoding/xml"

type Attr struct {
	StdAttributes
	XMLName xml.Name `xml:"attr"`
	Name    string   `xml:"name,attr"`

	String   String   `xml:"string,omitempty"`
	Float    Float    `xml:"float,omitempty"`
	Integer  Integer  `xml:"integer,omitempty"`
	Boolean  Boolean  `xml:"boolean,omitempty"`
	Datetime Datetime `xml:"datetime,omitempty"`

	Instance Instance `xml:"instance"`
}
