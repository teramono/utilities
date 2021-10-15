package configs

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/creasty/defaults"
)

// GigamonoConfig ...
type GigamonoConfig struct {
	Meta   Meta `json:"meta"`
	Broker struct {
		URL           string `json:"url" default:"nats://127.0.0.1:4222"`
		Subscriptions struct {
			Workspaces EnabledSubscription `json:"workspaces"`
			Logs       EnabledSubscription `json:"logs"`
		} `json:"subscriptions"`
	} `json:"broker"`
	Engines struct {
		API struct {
			Port         uint   `json:"port"`
			DBURL        string `json:"dbURL"`
			ReplyTimeout uint   `json:"replyTimeout" default:"1"`
		} `json:"api"`
		Backend struct {
			RootPath      string `json:"rootPath"`
			Subscriptions struct {
				Workspaces struct {
					Run EngineSubscription `json:"run" default:"1"`
				} `json:"workspaces"`
			} `json:"subscriptions"`
		} `json:"backend"`
		DB struct {
			DBURL         string `json:"dbURL"`
			Subscriptions struct {
				Workspaces struct {
					Query EngineSubscription `json:"query"`
				} `json:"workspaces"`
			} `json:"subscriptions"`
		} `json:"db"`
	} `json:"engines"`
	UI struct {
		Dir string `json:"dir"`
	} `json:"ui"`
	Logs struct {
		File        string `json:"file"`
		IsPublished string `json:"isPublished"`
	} `json:"logs"`
}

type EnabledSubscription struct {
	Version uint `json:"version"`
}

type EngineSubscription struct {
	SubscribedID string `json:"subscribedId" default:"*"`
}

// NewGigamonoConfig ...
func NewGigamonoConfig(configBytes []byte, format ConfigFormat) (GigamonoConfig, error) {
	config := GigamonoConfig{}

	if err := UnmarshalConfig(configBytes, format, &config); err != nil {
		return config, err
	}

	// Set defaults.
	if err := defaults.Set(&config); err != nil {
		return config, err
	}

	// Validate config type.
	if strings.ToLower(config.Meta.Type) != "gigamono" {
		return config, GetWrongTypeError("Gigamono")
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

	return NewGigamonoConfig(fileContent, format)
}
