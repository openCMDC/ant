package db

import (
	"ant/core"
	"fmt"
)

type Interface interface {
	StartConsume() error
	InsertRow(row *core.Row)
}

type DefaultDB struct {
	antCtx *core.AntContext
	cache  chan *core.Row
}

func (t *DefaultDB) StartConsume() error {
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

func (t DefaultDB) InsertRow(row *core.Row) {
	t.cache <- row
}

func NewTestStorage(ctx *core.AntContext) Interface {
	return &DefaultDB{
		antCtx: ctx,
		cache:  make(chan *core.Row, 4096),
	}
}
