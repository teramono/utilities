package models

import (
	"fmt"

	"github.com/teramono/utilities/pkg/database"
	"gorm.io/gorm"
)

// Workspace ...
type Workspace struct {
	gorm.Model
	Name string
}

// GetByName ...
func (workspace *Workspace) GetByName(db *database.DB) (Workspace, error) {
	tempWorkspace := Workspace{}

	fmt.Println("workspace", workspace)

	db.Where("name = ?", workspace.Name).First(&tempWorkspace)

	return tempWorkspace, nil
}
