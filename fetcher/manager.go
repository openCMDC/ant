package fetcher

import (
	"ant/core"
	"ant/fetcher/base"
	"ant/fetcher/networkfetcher"
)

type Manager struct {
	fetcherBackend *base.FetcherBackend
	netWorkFetcher *networkfetcher.NetworkFetcherRuntime
}

func (m *Manager) Start() error {
	return m.netWorkFetcher.Start()
}

func (m *Manager) RegisterFechedRowProcessor(processor base.RowProcessor) {
	m.fetcherBackend.RegisterFechedRowProcessor(processor)
}

func (m *Manager) initFetchers() error {
	runtime, err := networkfetcher.NewInstance(m.fetcherBackend)
	if err != nil {
		return err
	}
	runtime.InitCapture()
	m.netWorkFetcher = runtime
	return nil
}

func NewFetcherManager(remote core.Remote) (*Manager, error) {
	manager := new(Manager)
	manager.fetcherBackend = base.NewFetcherBackend()
	err := manager.initFetchers()
	if err != nil {
		return nil, err
	}
	//todo init logFetcher and systemFetcher
	return manager, nil
}
