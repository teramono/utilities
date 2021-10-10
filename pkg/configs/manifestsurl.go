package configs

type SurlManifest struct {
	Meta              Meta     `json:"meta"`
	MiddlewareScripts []string `json:"middlewareScripts"`
}

func NewSurlManifest(manifestBytes []byte, format ConfigFormat) (SurlManifest, error) {
	manifest := SurlManifest{}
	if err := UnmarshalConfig(manifestBytes, format, &manifest); err != nil {
		return manifest, err
	}

	return manifest, nil
}
