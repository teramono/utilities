package init

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
func NewSetup() (Setup, error) {
	// ...
	godotenv.Load()

	// ...
	connectionURI := os.Getenv("WORKSPACES_CONNECTION_URI")
	if connectionURI != "" {
		return Setup{}, fmt.Errorf("???")
	}

	// ...
	db, err := database.Connect(connectionURI)
	if err != nil {
		return Setup{}, err
	}

	return Setup{WorkspacesDB: db}, err
}

