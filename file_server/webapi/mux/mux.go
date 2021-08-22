package mux

import (
	"errors"
	"net/http"
	"os"
	"strings"
)

const (
	get_method = "GET"
)

type ErrorEntity struct {
	Kind    string `json:"kind"`
	Message string `json:"message"`
}

type ErrorDeclarations = []ErrorEntity

const (
	fileErrorMessage = "failed to find file"
	textHTML         = "text/html"
	applicationJson  = "application/json"
	indexHTML        = "index.html"
	dirPath          = "/"
)

var (
	osBaseDir          = os.Getenv("FILESERVER_FILEPATH")
	baseDir            = strings.TrimSuffix(osBaseDir, dirPath)
	notFoundFilepath   = os.Getenv("NOT_FOUND_FILEPATH")
	badRequestFilepath = os.Getenv("BAD_REQUEST_FILEPATH")
	errMethod          = errors.New("request method is not GET")
)

func validGet(r *http.Request) error {
	if r.Method == get_method {
		return nil
	}

	return errMethod
}

func getFilepath(r *http.Request) *string {
	filepath := baseDir + r.URL.Path
	stats, errStats := os.Stat(filepath)
	if os.IsNotExist(errStats) {
		return nil
	}

	if !stats.IsDir() {
		return &filepath
	}

	filepathWithIndex := baseDir + r.URL.Path + indexHTML
	_, errIndexStats := os.Stat(filepathWithIndex)
	if !os.IsNotExist(errIndexStats) {
		return &filepathWithIndex
	}

	return nil
}

func writeBadRequest(w http.ResponseWriter, r *http.Request, filepath *string, err error) error {
	if err != nil || filepath != nil {
		return err
	}

	w.WriteHeader(http.StatusBadRequest)
	http.ServeFile(w, r, badRequestFilepath)

	return err
}

func writeNotFound(w http.ResponseWriter, r *http.Request, filepath *string, err error) error {
	if err != nil || filepath != nil {
		return err
	}

	w.WriteHeader(http.StatusNotFound)
	http.ServeFile(w, r, notFoundFilepath)

	return err
}

func writeResponse(w http.ResponseWriter, r *http.Request, filepath *string, err error) error {
	if err != nil || filepath == nil {
		return err
	}

	http.ServeFile(w, r, *filepath)

	return err
}

func exec(w http.ResponseWriter, r *http.Request) {
	filepath := getFilepath(r)
	errGet := validGet(r)
	errResponse := writeResponse(w, r, filepath, errGet)
	errNotFound := writeNotFound(w, r, filepath, errResponse)
	writeBadRequest(w, r, filepath, errNotFound)
}

func CreateMux() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc(dirPath, exec)

	return mux
}
