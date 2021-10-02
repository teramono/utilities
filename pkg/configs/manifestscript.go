package configs

import (
	"strings"

	"github.com/spf13/viper"
)

// ScriptManifest ...
type ScriptManifest struct {
	Meta            Meta     `json:"meta"`
	MiddlewarePaths []string `json:"middlewarePaths"`
}

// NewScriptManifest ...
func NewScriptManifest(manifestString string, format ConfigFormat) (ScriptManifest, error) {
	// TODO: Sec: Validation
	manifest := ScriptManifest{}
	reader := strings.NewReader(manifestString)

	// Set format to parse.
	converter := viper.New()
	converter.SetConfigType(string(format))
	converter.ReadConfig(reader)

	// Unmarshal string into object.
	if err := converter.Unmarshal(&manifest); err != nil {
		return ScriptManifest{}, err
	}

	return manifest, nil
}
