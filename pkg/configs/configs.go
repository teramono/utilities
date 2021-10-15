package configs

import (
	"bytes"
	"fmt"

	"github.com/spf13/viper"
)

func UnmarshalConfig(configBytes []byte, format ConfigFormat, obj interface{}) error {
	// TODO: SEC: Validation
	reader := bytes.NewReader(configBytes)

	// Set format to parse.
	converter := viper.New()
	converter.SetConfigType(string(format))
	converter.ReadConfig(reader)

	// Unmarshal string into object.
	if err := converter.Unmarshal(obj); err != nil {
		return err
	}

	return nil
}

func GetWrongTypeError(configType string) error {
	return fmt.Errorf("wrong type, expected `%s`", configType)
}
