package configs

type SubappManifest struct {
	Meta Meta `json:"meta"`
}

func NewSubappManifest(manifestBytes []byte, format ConfigFormat) (SubappManifest, error) {
	manifest := SubappManifest{}
	if err := UnmarshalConfig(manifestBytes, format, &manifest); err != nil {
		return manifest, err
	}

	return manifest, nil
}
