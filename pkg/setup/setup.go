package setup

import (
	"log"

	"github.com/teramono/utilities/pkg/broker"
	"github.com/teramono/utilities/pkg/configs"
	"github.com/teramono/utilities/pkg/database"
	"github.com/teramono/utilities/pkg/logs"
)

type CommonSetup struct {
	broker.BrokerClient
	Config configs.GigamonoConfig
}

type APIEngineSetup struct {
	CommonSetup
	DB database.DB
}

func NewCommonSetup() (CommonSetup, error) {
	// Load gigamono config file.
	config, err := configs.LoadGigamonoConfig()
	if err != nil {
		log.Printf("unable to load config: %v", err)
	}

	// Set log file.
	if err = logs.SetLogFile(config.Logs.File); err != nil {
		log.Println(err)
	}

	// Connect to broker server.
	brokerClient, err := broker.NewBrokerClient(config.Broker.URL)
	if err != nil {
		return CommonSetup{}, err
	}

	return CommonSetup{
		BrokerClient: brokerClient,
		Config:       config,
	}, nil
}

func NewAPIEngineSetup() (APIEngineSetup, error) {
	commonSetup, err := NewCommonSetup()
	if err != nil {
		return APIEngineSetup{}, err
	}

	// Connect to api.db.
	db, err := database.ConnectDB(commonSetup.Config.Engines.API.DBURL)
	if err != nil {
		return APIEngineSetup{}, err
	}

	return APIEngineSetup{
		CommonSetup: commonSetup,
		DB:          db,
	}, nil
}
