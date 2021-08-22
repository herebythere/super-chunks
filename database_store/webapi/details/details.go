package details

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type ConfigDetails struct {
	Filepath     string `json:"filepath"`
	FilepathTest string `json:"filepath_test"`
}

type ServerDetails struct {
	HTTPPort     int64 `json:"http_port"`
	IdleTimeout  int64 `json:"idle_timeout"`
	ReadTimeout  int64 `json:"read_timeout"`
	WriteTimeout int64 `json:"write_timeout"`
}

type DBDetails struct {
	Host         string `json:"host"`
	Port         int64  `json:"port"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	DatabaseName string `json:"database_name"`
}

type SuperDBDetails struct {
	ServiceName string        `json:"service_name"`
	Config      ConfigDetails `json:"config"`
	Server      ServerDetails `json:"server"`
	DB          DBDetails     `json:"db"`
}

var (
	detailsPath         = os.Getenv("CONFIG_FILEPATH")
	Details, DetailsErr = ReadDetailsFromFile(detailsPath)
)

func readFile(path string) (*[]byte, error) {
	detailsJSON, errDetiailsJSON := ioutil.ReadFile(path)
	return &detailsJSON, errDetiailsJSON
}

func parseDetails(detailsJSON *[]byte, err error) (*SuperDBDetails, error) {
	if err != nil {
		return nil, err
	}

	var details SuperDBDetails
	errDetails := json.Unmarshal(*detailsJSON, &details)

	return &details, errDetails
}

func ReadDetailsFromFile(path string) (*SuperDBDetails, error) {
	detailsJSON, errDetailsJSON := readFile(path)
	return parseDetails(detailsJSON, errDetailsJSON)
}
