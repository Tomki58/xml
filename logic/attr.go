package logic

import (
	mxml "myxml/models/xml"
)

// GetValue gets the value from attr.
func GetValue(a mxml.Attr) interface{} {
	// check string
	if a.String.XMLName.Local != "" {
		return a.String.Value
	}

	if a.Float.XMLName.Local != "" {
		return a.Float.Value
	}

	if a.Integer.XMLName.Local != "" {
		return a.Integer.Value
	}

	if a.Boolean.XMLName.Local != "" {
		return a.Integer.Value
	}

	if a.Datetime.XMLName.Local != "" {
		return a.Datetime.Value
	}

	return nil
}
