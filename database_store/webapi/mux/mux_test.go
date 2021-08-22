package mux

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"webapi/pgsqlx"
)

const (
	testKind      = "test kind"
	testBodyError = "failed to execute hello world"
	expectedHello = "hello world!"
	statusOk      = 200
	statusNotOk   = 400
)

var (
	testBody = pgsqlx.Statement{
		Sql:    fmt.Sprint("SELECT $1"),
		Values: []interface{}{expectedHello},
	}

	expectedResponse = [][]interface{}{[]interface{}{"hello world!"}}
)

func TestCreateMux(t *testing.T) {
	proxyMux := CreateMux()
	if proxyMux == nil {
		t.Fail()
		t.Logf("proxyMux was not created")
	}
}

func TestWriteError(t *testing.T) {
	testRecorder := httptest.NewRecorder()
	writeError(testRecorder, testKind, testBodyError)

	if testRecorder.Code != statusNotOk {
		t.Fail()
		t.Logf(fmt.Sprint("expected: ", statusNotOk, ", found: ", testRecorder.Code))
	}

	var errors ErrorDeclarations
	json.NewDecoder(testRecorder.Body).Decode(&errors)

	if len(errors) == 0 {
		t.Fail()
		t.Logf("error array has a length of zero")
		return
	}

	if errors[0].Kind != testKind {
		t.Fail()
		t.Logf(fmt.Sprint("expected: ", testKind, ", found: ", errors[0].Kind))
	}

	if errors[0].Message != testBodyError {
		t.Fail()
		t.Logf(fmt.Sprint("expected: ", testBodyError, ", found: ", errors[0].Message))
	}
}

func TestWriteResponse(t *testing.T) {
	testRecorder := httptest.NewRecorder()
	writeResponse(testRecorder, expectedResponse, nil)

	if testRecorder.Code != statusOk {
		t.Fail()
		t.Logf(fmt.Sprint("expected: ", statusOk, ", found: ", testRecorder.Code))
	}

	var result [][]interface{}
	json.NewDecoder(testRecorder.Body).Decode(&result)

	resultsLength := len(result)
	if resultsLength != 1 {
		t.Fail()
		t.Logf(fmt.Sprint("expected a length of ", 1, ", instead found a length of ", resultsLength))
	}

	resultEntry := result[0][0]
	if resultEntry != expectedHello {
		t.Fail()
		t.Logf(fmt.Sprint("expected: ", expectedHello, ", found: ", resultEntry))
	}
}

func TestGetBody(t *testing.T) {
	bodyBytes := new(bytes.Buffer)
	errJson := json.NewEncoder(bodyBytes).Encode(&testBody)
	if errJson != nil {
		t.Fail()
		t.Logf(errJson.Error())
		return
	}

	req, errReq := http.NewRequest("POST", "/", bodyBytes)
	if errReq != nil {
		t.Fail()
		t.Logf(errReq.Error())
		return
	}

	reqBody, errReqBody := getBody(req, nil)
	if errReqBody != nil {
		t.Fail()
		t.Logf(fmt.Sprint("expected an array, ", errReqBody.Error()))
		return
	}
	if reqBody == nil {
		t.Fail()
		t.Logf(fmt.Sprint("request body is nil"))
		return
	}
}

func TestValidPost(t *testing.T) {
	bodyBytes := new(bytes.Buffer)
	errJson := json.NewEncoder(bodyBytes).Encode(&testBody)
	if errJson != nil {
		t.Fail()
		t.Logf(errJson.Error())
		return
	}

	req, errReq := http.NewRequest("POST", "/", bodyBytes)
	if errReq != nil {
		t.Fail()
		t.Logf(errReq.Error())
		return
	}

	errValid := validPost(req)
	if errValid != nil {
		t.Fail()
		t.Logf(fmt.Sprint("expected an array, ", errValid.Error()))
		return
	}
}

func TestQuery(t *testing.T) {
	bodyBytes := new(bytes.Buffer)
	errJson := json.NewEncoder(bodyBytes).Encode(testBody)
	if errJson != nil {
		t.Fail()
		t.Logf(errJson.Error())
		return
	}

	req, errReq := http.NewRequest("POST", "/", bodyBytes)
	if errReq != nil {
		t.Fail()
		t.Logf(errReq.Error())
		return
	}

	testRecorder := httptest.NewRecorder()
	query(testRecorder, req)

	if testRecorder.Code != statusOk {
		t.Fail()
		t.Logf(fmt.Sprint("expected: ", statusOk, ", found: ", testRecorder.Code))
	}

	var result [][]interface{}
	errJsonDecode := json.NewDecoder(testRecorder.Body).Decode(&result)
	if errJsonDecode != nil {
		t.Fail()
		t.Logf(errJsonDecode.Error())
		return
	}

	resultsLength := len(result)
	if resultsLength != 1 {
		t.Fail()
		t.Logf(fmt.Sprint("expected a length of ", 1, ", instead found a length of ", resultsLength))
	}

	resultEntry := result[0][0]
	if resultEntry != expectedHello {
		t.Fail()
		t.Logf(fmt.Sprint("expected: ", expectedHello, ", found: ", resultEntry))
	}
}
