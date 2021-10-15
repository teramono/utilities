package configs

import (
	"strings"

	"github.com/creasty/defaults"
)

// ExtensionManifest ...
type ExtensionManifest struct {
	Meta Meta `json:"meta"`
}

// NewExtensionManifest ...
func NewExtensionManifest(manifestBytes []byte, format ConfigFormat) (ExtensionManifest, error) {
	manifest := ExtensionManifest{}

	if err := UnmarshalConfig(manifestBytes, format, &manifest); err != nil {
		return manifest, err
	}

	// Set defaults.
	if err := defaults.Set(&manifest); err != nil {
		return manifest, err
	}

	// Validate manifest type.
	if strings.ToLower(manifest.Meta.Type) != "extension" {
		return manifest, GetWrongTypeError("Extension")
	}

	return manifest, nil
}
