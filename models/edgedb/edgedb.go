package models

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
)

type baseAggr struct {
	Domain string
	Name   string
	Etype  string
	UID    uuid.UUID
	Owner  uuid.UUID

	From time.Time
	To   time.Time
}

type Property struct {
	Name  string
	Value interface{}
}

type InsertQuery struct {
	baseAggr

	Properties []Property
}

func (ir InsertQuery) ToDDL() (string, error) {
	builder := strings.Builder{}

	builder.WriteString(fmt.Sprintf(
		"INSERT %s::%s {\n", ir.Domain, ir.Name,
	))

	var counter int
	for _, property := range ir.Properties {
		switch val := property.Value.(type) {
		case string:
			builder.WriteString(fmt.Sprintf(
				"\t%s := <std::str>$%d,\n", property.Name, counter,
			))
			counter++

		case InsertQuery:
			ir, err := val.ToDDL()
			if err != nil {
				return "", err
			}
			builder.WriteString(fmt.Sprintf(
				"\t%s := (%s),\n", property.Name, ir,
			))

		case SelectQuery:
			sr, err := val.ToDDL()
			if err != nil {
				return "", nil
			}
			builder.WriteString(fmt.Sprintf(
				"\t%s := (%s),\n", property.Name, sr,
			))
		}
	}

	builder.WriteString("};\n")

	return builder.String(), nil
}

type SelectQuery struct {
	baseAggr
}

func (sr SelectQuery) ToDDL() (string, error) {
	builder := strings.Builder{}

	builder.WriteString(fmt.Sprintf(
		"SELECT %s::%s\n", sr.Domain, sr.Name,
	))
	builder.WriteString(fmt.Sprintf(
		"FILTER .uuid = <std::uuid>%q", sr.UID,
	))

	return builder.String(), nil
}
