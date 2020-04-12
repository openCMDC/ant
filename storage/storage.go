package storage

import (
	"ant/core"
	"fmt"
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
		count := 0
		for r := range t.cache {
			fmt.Println(r.String())
			count++
			fmt.Println(count)
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
