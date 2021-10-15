package configs

import (
	"strings"

	"github.com/creasty/defaults"
)

type SubappManifest struct {
	Meta Meta `json:"meta"`
}

func NewSubappManifest(manifestBytes []byte, format ConfigFormat) (SubappManifest, error) {
	manifest := SubappManifest{}

	if err := UnmarshalConfig(manifestBytes, format, &manifest); err != nil {
		return manifest, err
	}

	// Set defaults.
	if err := defaults.Set(&manifest); err != nil {
		return manifest, err
	}

	// Validate manifest type.
	if strings.ToLower(manifest.Meta.Type) != "subapp" {
		return manifest, GetWrongTypeError("Subapp")
	}

	return manifest, nil
}
