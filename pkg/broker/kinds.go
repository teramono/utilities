package broker

type EngineKind uint

const (
	EngineAPI     EngineKind = iota
	EngineBackend EngineKind = iota
	EngineDB      EngineKind = iota
)
