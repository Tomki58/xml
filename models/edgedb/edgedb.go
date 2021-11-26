package models

import (
	"time"

	"github.com/google/uuid"
)

type baseAggr struct {
	Domain string
	Name   string
	Etype  string
	UID    uuid.UUID
	Owner  uuid.UUID

	To   time.Time
	From time.Time
}

type InsertQuery struct {
	baseAggr

	Properties []Property
	Links      []Link
}

type Property struct {
	Value      interface{}
	Annotation string
}

type Link struct{}
