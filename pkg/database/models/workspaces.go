package models

import (
	"errors"

	"github.com/teramono/utilities/pkg/database"
	"gorm.io/gorm"
)

type Workspace struct {
	gorm.Model
	Name string
}

func (workspace *Workspace) Create(db *database.DB) error {
	result := db.Create(workspace)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (workspace *Workspace) GetByName(db *database.DB) (Workspace, error) {
	tempWorkspace := Workspace{}

	result := db.Where("name = ?", workspace.Name).First(&tempWorkspace)
	if result.Error != nil {
		return tempWorkspace, result.Error
	}

	return tempWorkspace, nil
}

func (workspace *Workspace) Exists(db *database.DB) (bool, error) {
	tempWorkspace := Workspace{}

	result := db.First(&tempWorkspace, workspace.ID)
	if err := result.Error; err != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return false, nil
		}

		return false, result.Error
	}

	return true, nil
}
