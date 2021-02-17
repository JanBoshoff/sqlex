package sqlex_test

import (
	"reflect"
	"testing"
)

func NewColumnTypeMock(typeHolder interface{}) ColumnTypeMock {
	return ColumnTypeMock{
		test: typeHolder,
	}
}

type ColumnTypeMock struct {
	test interface{}
}

func (ct ColumnTypeMock) ScanType() reflect.Type {
	return reflect.TypeOf(ct.test)
}

func failTest(t *testing.T, exptected interface{}, received interface{}) {
	t.Log("Expected: ", exptected)
	t.Log("Received: ", received)
	t.Fail()
}
