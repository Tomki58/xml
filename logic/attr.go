package logic

import (
	mxml "myxml/models/xml"
)

// GetValue gets the value from attr.
func GetValue(a mxml.Attr) (interface{}, error) {
	// check string
	if a.String.XMLName.Local != "" {
		return a.String.Value, nil
	}

	// check float
	if a.Float.XMLName.Local != "" {
		return a.Float.Value, nil
	}

	// check integer
	if a.Integer.XMLName.Local != "" {
		return a.Integer.Value, nil
	}

	// check bool
	if a.Boolean.XMLName.Local != "" {
		return a.Integer.Value, nil
	}

	// check datetime
	if a.Datetime.XMLName.Local != "" {
		return a.Datetime.Value, nil
	}

	// check instance
	if a.Instance.XMLName.Local != "" {
		sr := CreateSelectRequest(a.Instance)
		return sr, nil
	}

	if i := a.Instance; i.XMLName.Local != "" {

	}

	return nil, nil
}
