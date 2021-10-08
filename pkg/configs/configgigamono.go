package configs

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

// GigamonoConfig ...
type GigamonoConfig struct {
	Meta    Meta `json:"meta"`
	Engines struct {
		API struct {
			Port  uint   `json:"port"`
			DBURL string `json:"dbURL"`
		} `json:"api"`
	} `json:"engines"`
	UI struct {
		Dir string `json:"dir"`
	} `json:"ui"`
	Broker struct {
		URL           string `json:"url"`
		Subscriptions struct {
			Workspaces Subscription `json:"workspaces"`
			Logs       Subscription `json:"logs"`
		} `json:"subscriptions"`
	} `json:"broker"`
	Logs struct {
		File        string `json:"file"`
		IsPublished string `json:"isPublished"`
	} `json:"logs"`
}

type Subscription struct {
	Version uint `json:"version"`
}

// NewGigamonoConfig ...
func NewGigamonoConfig(gigamonoString string, format ConfigFormat) (GigamonoConfig, error) {
	// TODO: Sec: Validation
	config := GigamonoConfig{}
	reader := strings.NewReader(gigamonoString)

	// Set format to parse.
	converter := viper.New()
	converter.SetConfigType(string(format))
	converter.ReadConfig(reader)

	// Unmarshal string into object.
	if err := converter.Unmarshal(&config); err != nil {
		return GigamonoConfig{}, err
	}

	return config, nil
}

func LoadGigamonoConfig() (GigamonoConfig, error) {
	// Get config file path from env.
	path := os.Getenv(GigamonoConfigPathEnvVar)
	if path == "" {
		return GigamonoConfig{}, fmt.Errorf(
			"cannot find environment variable `%s` needed for loading config",
			GigamonoConfigPathEnvVar,
		)
	}

	// Get file extension and use it to determine config format.
	fileExtension := filepath.Ext(path)
	format, err := ToConfigFormat(fileExtension[1:])
	if err != nil {
		return GigamonoConfig{}, fmt.Errorf("file extension `%s` for config not supported", fileExtension)
	}

	// Read file.
	fileContent, err := ioutil.ReadFile(path)
	if err != nil {
		return GigamonoConfig{}, err
	}

	return NewGigamonoConfig(string(fileContent), format)
}
