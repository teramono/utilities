package setup

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/teramono/utilities/pkg/configs"
	"github.com/teramono/utilities/pkg/database"
)

// Setup ...
type Setup struct {
	Config       configs.GigamonoConfig
	WorkspacesDB database.DB
}

// NewSetup ...
func NewSetup(enableDatabase bool) (Setup, error) {
	// ...
	godotenv.Load()

	// ...
	var db database.DB
	var err error
	if enableDatabase {
		connectionURI := os.Getenv("WORKSPACES_CONNECTION_URI")
		if connectionURI == "" {
			return Setup{}, fmt.Errorf("Could not connect!")
		}

		// ...
		db, err = database.Connect(connectionURI)
		if err != nil {
			return Setup{}, err
		}
	}

	return Setup{WorkspacesDB: db}, nil
}
