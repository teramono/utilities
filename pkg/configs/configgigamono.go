package configs

import (
	"strings"

	"github.com/spf13/viper"
)

// GigamonoConfig ...
type GigamonoConfig struct {
	Meta Meta `json:"meta"`
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
