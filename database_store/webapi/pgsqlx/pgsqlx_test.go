package pgsqlx

import (
	"fmt"
	"testing"
)

func TestCreate(t *testing.T) {
	conn, errConn := create(nil)
	if conn != nil {
		t.Error("nil parameters should return nil")
	}
	if errConn == nil {
		t.Error("nil paramters should return error")
	}
}

func TestSetterWasCreated(t *testing.T) {
	if errPool != nil {
		t.Fail()
		t.Logf(errPool.Error())
	}
}

func TestSetterQueries(t *testing.T) {
	expected := "hello world!"
	statement := &Statement{
		Sql:    "SELECT $1",
		Values: []interface{}{expected},
	}
	results, errResults := Query(statement, nil)

	if results == nil {
		t.Fail()
		t.Logf("there should be sql results!")
		return
	}

	if errResults != nil {
		t.Fail()
		t.Logf(errResults.Error())
		return
	}

	result := (*results)[0][0]
	if result != expected {
		t.Fail()
		t.Logf(fmt.Sprint("expected: ", expected, ", found: ", result))
	}
}
