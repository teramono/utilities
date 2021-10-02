package configs

import (
	"strings"

	"github.com/spf13/viper"
)

// ExtensionManifest ...
type ExtensionManifest struct {
	Meta            Meta     `json:"meta"`
}

// NewExtensionManifest ...
func NewExtensionManifest(manifestString string, format ConfigFormat) (ExtensionManifest, error) {
	// TODO: Sec: Validation
	manifest := ExtensionManifest{}
	reader := strings.NewReader(manifestString)

	// Set format to parse.
	converter := viper.New()
	converter.SetConfigType(string(format))
	converter.ReadConfig(reader)

	// Unmarshal string into object.
	if err := converter.Unmarshal(&manifest); err != nil {
		return ExtensionManifest{}, err
	}

	return manifest, nil
}
