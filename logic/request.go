package logic

import (
	medge "myxml/models/edgedb"
	mxml "myxml/models/xml"
)

func CreateInsertRequest(instance mxml.Instance) (*medge.InsertQuery, error) {
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
			value, err := GetValue(attr)
			if err != nil {
				return nil, err
			}
			query.Properties = append(query.Properties, medge.Property{
				Name:  attr.Name,
				Value: value,
			})
		}
	}

	return &query, nil
}

func CreateSelectRequest(instance mxml.Instance) medge.SelectQuery {
	var query medge.SelectQuery

	query.Domain = instance.Domain
	query.Name = instance.Name
	for _, generic := range instance.Generics {
		query.UID = generic.UID
	}

	return query
}
