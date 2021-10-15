package configs

import (
	"strings"

	"github.com/creasty/defaults"
)

type SurlManifest struct {
	Meta              Meta     `json:"meta"`
	MiddlewareScripts []string `json:"middlewareScripts"`
}

func NewSurlManifest(manifestBytes []byte, format ConfigFormat) (SurlManifest, error) {
	manifest := SurlManifest{}

	if err := UnmarshalConfig(manifestBytes, format, &manifest); err != nil {
		return manifest, err
	}

	// Set defaults.
	if err := defaults.Set(&manifest); err != nil {
		return manifest, err
	}

	// Validate manifest type.
	if strings.ToLower(manifest.Meta.Type) != "surl" {
		return manifest, GetWrongTypeError("Surl")
	}

	return manifest, nil
}
