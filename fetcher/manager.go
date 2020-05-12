package fetcher

import (
	"ant/core"
	"ant/db"
	"ant/fetcher/base"
	"ant/fetcher/networkfetcher"
)

type Manager struct {
	fetcherCtx     *base.FetcherCtx
	netWorkFetcher *networkfetcher.NetworkDataFetcher
}

func (f *Manager) Start() error {
	f.netWorkFetcher.Start()
	return nil
}

func NewFetcherManager(ctx core.AntContext, storage db.Interface) *Manager {
	fetcherCtx := new(base.FetcherCtx)
	fetcherCtx.AntCtx = ctx
	fetcherCtx.Storage = storage

	manager := new(Manager)
	manager.fetcherCtx = fetcherCtx
	manager.netWorkFetcher = networkfetcher.NewInstance(fetcherCtx)
	return manager
}
