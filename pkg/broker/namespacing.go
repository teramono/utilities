package broker

import (
	"fmt"

	"github.com/teramono/utilities/pkg/configs"
)

func Namespace(version uint, action string, object string) string {
	return fmt.Sprintf("v%d.%s.%s", version, action, object)
}

func GetSubjectWithId(version uint, action string, object string, id interface{}) string {
	return fmt.Sprintf("%s.%v", Namespace(version, action, object), id)
}

func GetWorkspacesSubjectWithId(config *configs.GigamonoConfig, action string, id interface{}) string {
	version := config.Broker.Subscriptions.Workspaces.Version
	return GetSubjectWithId(version, action, "workspaces", id)
}

func GetWorkspacesSubjectByEngine(engine EngineKind, config *configs.GigamonoConfig, action string) string {
	version := config.Broker.Subscriptions.Workspaces.Version
	return GetSubjectWithId(version, action, "workspaces", GetIDByEngine(engine, config, action))
}

func GetWorkspacesResponderGroupByEngine(engine EngineKind, config *configs.GigamonoConfig, action string) string {
	version := config.Broker.Subscriptions.Workspaces.Version
	return GetSubjectWithId(version, action, "workspacesResponder", GetIDByEngine(engine, config, action))
}

func GetIDByEngine(engine EngineKind, config *configs.GigamonoConfig, action string) string {
	var id = "*"

	switch engine {
	case EngineBackend:
		switch action {
		case "run":
			id = config.Engines.Backend.Subscriptions.Workspaces.Run.SubscribedID
		}
	case EngineDB:
		switch action {
		case "query":
			id = config.Engines.Backend.Subscriptions.Workspaces.Run.SubscribedID
		}
	}

	return id
}
