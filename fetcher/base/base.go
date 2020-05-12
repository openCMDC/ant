package base

import (
	"ant/core"
	"ant/db"
)

type DataFetcher interface {
	Name() string
	Start() error
	Pause() error
	Stop() error
}

type FetcherCtx struct {
	Storage db.Interface
	AntCtx  core.AntContext
}
