package setup

import (
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
	var db database.DB

	// ...

	return Setup{WorkspacesDB: db}, nil
}
