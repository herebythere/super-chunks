package mux

import (
	"encoding/json"
	"errors"
	"net/http"

	"webapi/redisx"
)

const (
	contentType     = "Content-Type"
	applicationJson = "application/json"
	execRoute       = "/"
	post            = "POST"
)

type ErrorEntity struct {
	Kind    string `json:"kind"`
	Message string `json:"message"`
}

type ErrorDeclarations = []ErrorEntity

const (
	failedToExec = "failed to exec command"
)

var (
	errMethodMessage  = errors.New("request method is not GET")
	errNilRequestBody = errors.New("request body is nil")
)

func validPost(r *http.Request) error {
	if r.Method == post {
		return nil
	}

	return errMethodMessage
}

func getBody(r *http.Request, err error) (*[]interface{}, error) {
	if err != nil {
		return nil, err
	}

	if r.Body == nil {
		return nil, errNilRequestBody
	}

	var rBody []interface{}
	errRBody := json.NewDecoder(r.Body).Decode(&rBody)

	return &rBody, errRBody
}

func writeError(w http.ResponseWriter, kind string, message string) {
	setErrors := ErrorDeclarations{
		ErrorEntity{
			Kind:    kind,
			Message: message,
		},
	}

	w.WriteHeader(http.StatusBadRequest)
	w.Header().Set(contentType, applicationJson)
	json.NewEncoder(w).Encode(setErrors)
}

func writeResponse(w http.ResponseWriter, entry interface{}, err error) {
	if err != nil {
		writeError(w, failedToExec, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set(contentType, applicationJson)
	json.NewEncoder(w).Encode(entry)
}

func exec(w http.ResponseWriter, r *http.Request) {
	errPost := validPost(r)
	rBody, errRBody := getBody(r, errPost)
	result, errResult := redisx.Exec(rBody, errRBody)

	writeResponse(w, result, errResult)
}

func CreateMux() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc(execRoute, exec)

	return mux
}
