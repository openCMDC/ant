package core

type DataFetcher interface {
	Name() string
	Start(ctx *AntContext) error
	Pause(ctx *AntContext) error
	Stop(ctx *AntContext) error
	GetDataChan() (chan Row, error)
}

type AntContext struct {
	TargetProcessID string
}
