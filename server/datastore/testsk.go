package datastore

import (
	"../../share/model"
	"code.google.com/p/go.net/context"
)

type TestSKstore interface {
	GetTestSKID(id int64) (*model.TestSK, error)

	GetTestSKName(name string) (*model.TestSK, error)
}

func GetTestSKID(c context.Context, id int64) (*model.TestSK, error) {
	return FromContext(c).GetTestSKID(id)
}

func GetTestSKName(c context.Context, name string) (*model.TestSK, error) {
	return FromContext(c).GetTestSKName(name)
}
