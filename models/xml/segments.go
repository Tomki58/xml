package models

import (
	"encoding/xml"
	"time"
)

type String struct {
	XMLName xml.Name `xml:"string"`
	Value   string   `xml:"value"`
}

type Float struct {
	XMLName xml.Name `xml:"float"`
	Value   float64  `xml:"value,omitempty"`
}

type Integer struct {
	XMLName xml.Name `xml:"integer"`
	Value   int32    `xml:"value,omitempty"`
}

type Boolean struct {
	XMLName xml.Name `xml:"boolean"`
	Value   bool     `xml:"value,omitempty"`
}

type Datetime struct {
	XMLName xml.Name  `xml:"datetime"`
	Value   time.Time `xml:"value,omitempty"`
}
