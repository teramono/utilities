package configs

import (
	"strings"

	"github.com/spf13/viper"
)

// SubappManifest ...
type SubappManifest struct {
	Meta Meta `json:"meta"`
}

// NewSubappManifest ...
func NewSubappManifest(manifestString string, format ConfigFormat) (SubappManifest, error) {
	// TODO: Sec: Validation
	manifest := SubappManifest{}
	reader := strings.NewReader(manifestString)

	// Set format to parse.
	converter := viper.New()
	converter.SetConfigType(string(format))
	converter.ReadConfig(reader)

	// Unmarshal string into object.
	if err := converter.Unmarshal(&manifest); err != nil {
		return SubappManifest{}, err
	}

	return manifest, nil
}
