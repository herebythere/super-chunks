// brian taylor vann
// details

package details

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"time"
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

type CacheDetails struct {
	Host        string        `json:"host"`
	IdleTimeout time.Duration `json:"idle_timeout"`
	MaxActive   int64         `json:"max_active"`
	MaxIdle     int64         `json:"max_idle"`
	MaxSamples  int64         `json:"max_samples"`
	MaxSizeInMB string        `json:"max_size_in_mb"`
	Port        int64         `json:"port"`
	Protocol    string        `json:"protocol"`
}

type SuperCacheDetails struct {
	ServiceName string        `json:"service_name"`
	Config      ConfigDetails `json:"config"`
	Server      ServerDetails `json:"server"`
	Cache       CacheDetails  `json:"cache"`
}

var (
	detailsPath         = os.Getenv("CONFIG_FILEPATH")
	Details, DetailsErr = ReadDetailsFromFile(detailsPath)
)

func readFile(path string) (*[]byte, error) {
	detailsJSON, errDetiailsJSON := ioutil.ReadFile(path)
	return &detailsJSON, errDetiailsJSON
}

func parseDetails(detailsJSON *[]byte, err error) (*SuperCacheDetails, error) {
	if err != nil {
		return nil, err
	}

	var details SuperCacheDetails
	errDetails := json.Unmarshal(*detailsJSON, &details)

	return &details, errDetails
}

func ReadDetailsFromFile(path string) (*SuperCacheDetails, error) {
	detailsJSON, errDetailsJSON := readFile(path)
	return parseDetails(detailsJSON, errDetailsJSON)
}
