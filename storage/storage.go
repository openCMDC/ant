package storage

import (
	"ant/core"
)

type Interface interface {
	StartConsume() error
	StoreRow(row *core.Row)
}

type TestStorage struct {
	antCtx core.AntContext
	cache  chan *core.Row
}

func (t *TestStorage) StartConsume() error {
	go func() {
		for r := range t.cache {
			//fmt.Println(r.String())
			r.String()
		}
	}()
	return nil
}

func (t TestStorage) StoreRow(row *core.Row) {
	t.cache <- row
}

func NewTestStorage(ctx core.AntContext) Interface {
	return &TestStorage{
		antCtx: ctx,
		cache:  make(chan *core.Row, 4096),
	}
}
