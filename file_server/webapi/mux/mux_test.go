package mux

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

const (
	testKind       = "test kind"
	statusNotFound = 404
	statusNotOk    = 400
	badPath        = "/39fdibjini4924inorg8onqjp90fjd9os"
)

func TestValidGet(t *testing.T) {
	postReq, errPostReq := http.NewRequest("POST", "/", nil)
	if errPostReq != nil {
		t.Fail()
		t.Logf(errPostReq.Error())
	}

	errValidPost := validGet(postReq)
	if errValidPost == nil {
		t.Fail()
		t.Logf("expected an error but found nil")
	}

	getReq, errGetReq := http.NewRequest("GET", "/", nil)
	if errGetReq != nil {
		t.Fail()
		t.Logf(errGetReq.Error())
	}

	errValidGet := validGet(getReq)
	if errValidGet != nil {
		t.Fail()
		t.Logf(errValidGet.Error())
	}
}

func TestGetFilepath(t *testing.T) {
	expected := baseDir + "/index.html"
	getReq, errGetReq := http.NewRequest("GET", "/", nil)
	if errGetReq != nil {
		t.Fail()
		t.Logf(errGetReq.Error())
	}

	result := getFilepath(getReq)
	if result != nil && *result != expected {
		t.Fail()
		t.Logf(fmt.Sprint("expected: ", expected, ", but found: ", *result))
	}
}

func TestGetBadFilepath(t *testing.T) {
	getReq, errGetReq := http.NewRequest("GET", badPath, nil)
	if errGetReq != nil {
		t.Fail()
		t.Logf(errGetReq.Error())
	}

	filepath := getFilepath(getReq)
	if filepath != nil {
		t.Fail()
		t.Logf(fmt.Sprint("expected nil but did not find a path"))
	}
}

func TestWriteResponse(t *testing.T) {
	testRecorder := httptest.NewRecorder()
	getReq, errGetReq := http.NewRequest("GET", "/", nil)
	if errGetReq != nil {
		t.Fail()
		t.Logf(errGetReq.Error())
		return
	}

	filepath := getFilepath(getReq)
	writeResponse(testRecorder, getReq, filepath, nil)
	if testRecorder.Code != http.StatusOK {
		t.Fail()
		t.Logf(fmt.Sprint("expected: ", http.StatusOK, ", found: ", testRecorder.Code))
	}

	_, errBody := ioutil.ReadAll(testRecorder.Body)
	if errBody != nil {
		t.Fail()
		t.Logf(errBody.Error())
	}
}

func TestWriteBadRequest(t *testing.T) {
	testRecorder := httptest.NewRecorder()
	getReq, errGetReq := http.NewRequest("GET", badPath, nil)
	if errGetReq != nil {
		t.Fail()
		t.Logf(errGetReq.Error())
		return
	}

	writeBadRequest(testRecorder, getReq, nil, nil)
	if testRecorder.Code != http.StatusBadRequest {
		t.Fail()
		t.Logf(fmt.Sprint("expected: ", http.StatusBadRequest, ", found: ", testRecorder.Code))
	}

	_, errBody := ioutil.ReadAll(testRecorder.Body)
	if errBody != nil {
		t.Fail()
		t.Logf(errBody.Error())
	}
}

func TestWriteNotFound(t *testing.T) {
	testRecorder := httptest.NewRecorder()
	getReq, errGetReq := http.NewRequest("GET", badPath, nil)
	if errGetReq != nil {
		t.Fail()
		t.Logf(errGetReq.Error())
		return
	}

	writeNotFound(testRecorder, getReq, nil, nil)
	if testRecorder.Code != http.StatusNotFound {
		t.Fail()
		t.Logf(fmt.Sprint("expected: ", http.StatusNotFound, ", found: ", testRecorder.Code))
	}

	_, errBody := ioutil.ReadAll(testRecorder.Body)
	if errBody != nil {
		t.Fail()
		t.Logf(errBody.Error())
	}
}
