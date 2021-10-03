package setup

import (
	"github.com/teramono/utilities/pkg/configs"
	"github.com/teramono/utilities/pkg/database"
)

type CommonSetup struct {
	Config configs.GigamonoConfig
}

type APIEngineSetup struct {
	CommonSetup
	WorkspacesDB database.DB
}

type SupervisorEngineSetup struct {
	CommonSetup
	ProcessesDB database.DB
}

func NewCommonSetup() (CommonSetup, error) {
	// Load Gigamono config file.
	config, err := configs.LoadGigamonoConfig()
	if err != nil {
		return CommonSetup{}, err
	}

	return CommonSetup{Config: config}, nil
}

func NewAPIEngineSetup() (APIEngineSetup, error) {
	commonSetup, err := NewCommonSetup()
	if err != nil {
		return APIEngineSetup{}, err
	}

	db, err := database.ConnectDB(commonSetup.Config.Engines.API.WorkspacesDBURI)
	if err != nil {
		return APIEngineSetup{}, err
	}

	return APIEngineSetup{
		CommonSetup:  commonSetup,
		WorkspacesDB: db,
	}, nil
}

func NewProvisionEngineSetup() (SupervisorEngineSetup, error) {
	commonSetup, err := NewCommonSetup()
	if err != nil {
		return SupervisorEngineSetup{}, err
	}

	db, err := database.ConnectDB(commonSetup.Config.Engines.Provision.ProvisionsDBURI)
	if err != nil {
		return SupervisorEngineSetup{}, err
	}

	return SupervisorEngineSetup{
		CommonSetup:  commonSetup,
		ProcessesDB: db,
	}, nil
}
