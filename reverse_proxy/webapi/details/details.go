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
	HTTPPort                      int64 `json:"http_port"`
	HTTPSPort                     int64 `json:"https_port"`
	IdleTimeout                   int64 `json:"idle_timeout"`
	ReadTimeout                   int64 `json:"read_timeout"`
	RedirectFromHTTPToHTTPS       bool  `json:"redirect_from_http_to_https"`
	SkipSSLVerificationOnForwards bool  `json:"skip_ssl_verification_on_forwards"`
	WriteTimeout                  int64 `json:"write_timeout"`
}

type CertPaths struct {
	Cert       string `json:"cert"`
	PrivateKey string `json:"private_key"`
}

type GatewayDetails struct {
	ServiceName string            `json:"service_name"`
	Config      ConfigDetails     `json:"config"`
	CertPaths   CertPaths         `json:"cert_paths"`
	Routes      map[string]string `json:"routes"`
	Server      ServerDetails     `json:"server"`
}

var (
	detailsPath         = os.Getenv("CONFIG_FILEPATH")
	Details, DetailsErr = ReadDetailsFromFile(detailsPath)
)

func readFile(path string) (*[]byte, error) {
	detailsJSON, errDetiailsJSON := ioutil.ReadFile(path)
	return &detailsJSON, errDetiailsJSON
}

func parseDetails(detailsJSON *[]byte, err error) (*GatewayDetails, error) {
	if err != nil {
		return nil, err
	}

	var details GatewayDetails
	errDetails := json.Unmarshal(*detailsJSON, &details)

	return &details, errDetails
}

func ReadDetailsFromFile(path string) (*GatewayDetails, error) {
	detailsJSON, errDetailsJSON := readFile(path)
	return parseDetails(detailsJSON, errDetailsJSON)
}
