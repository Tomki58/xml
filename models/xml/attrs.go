package models

import "encoding/xml"

type Attr struct {
	StdAttributes
	XMLName xml.Name `xml:"attr"`
	Name    string   `xml:"name,attr"`

	// semgents
	String   String   `xml:"string,omitempty"`
	Float    Float    `xml:"float,omitempty"`
	Integer  Integer  `xml:"integer,omitempty"`
	Boolean  Boolean  `xml:"boolean,omitempty"`
	Datetime Datetime `xml:"datetime,omitempty"`

	// instances
	Instance Instance `xml:"instance"`
}
