package logic

import (
	medge "myxml/models/edgedb"
	mxml "myxml/models/xml"
)

func New(instance mxml.Instance) medge.InsertQuery {
	var query medge.InsertQuery

	query.Domain = instance.Domain
	query.Name = instance.Name

	for _, generic := range instance.Generics {
		if generic.XMLName.Local != "" {
			query.UID = generic.UID
			query.Etype = generic.Etype
			query.Owner = generic.Owner
		}

		for _, attr := range generic.Attrs {
			query.Properties = append(query.Properties, medge.Property{
				Value:      GetValue(attr),
				Annotation: createAnnotation(attr),
			})
		}
	}

	return query
}

func createAnnotation(attr mxml.Attr) string {
	return ""
}
