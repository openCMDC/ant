package base

import (
	"ant/core"
	"ant/storage"
)

type DataFetcher interface {
	Name() string
	Start() error
	Pause() error
	Stop() error
}

type FetcherCtx struct {
	Storage storage.Interface
	AntCtx  core.AntContext
}
