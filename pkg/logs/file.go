package logs

import (
	"fmt"
	"log"

	"github.com/teramono/utilities/pkg/file"
)

func SetLogFile(filename string) error {
	file, err := file.OpenOrCreateFile(filename, true)
	if err != nil {
		return fmt.Errorf("cannot create log file `%s`: %w", filename, err)
	}

	log.SetPrefix("\n")
	log.SetOutput(file)
	return nil
}
