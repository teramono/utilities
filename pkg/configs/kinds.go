package configs

import (
	"fmt"
	"strings"
)

// ConfigFormat ...
type ConfigFormat string

// ...
const (
	YAML ConfigFormat = "YAML"
	JSON ConfigFormat = "JSON"
	TOML ConfigFormat = "TOML"
)

// ToConfigFormat ...
func ToConfigFormat(format string) (ConfigFormat, error) {
	switch strings.ToLower(format) {
	case "yaml", "yml":
		return YAML, nil
	case "json":
		return JSON, nil
	case "toml":
		return TOML, nil
	default:
		return "", fmt.Errorf("conversion from string `%s` to ConfigFormat not possible", format)
	}
}
