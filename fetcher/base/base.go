package base

import (
	"ant/core"
)

type DataFetcher interface {
	Name() string
	Start() error
	Pause() error
	Stop() error
}

type FetcherBackend struct {
	processors []RowProcessor
}

func (f *FetcherBackend) RegisterRowProcessor(processor RowProcessor) {
	f.processors = append(f.processors, processor)
}

func (f *FetcherBackend) ReportRow(row *core.Row) {
	if len(f.processors) == 0 {
		return
	}
	for _, processor := range f.processors {
		processor.OnRow(row)
	}
}
func NewFetcherBackend() *FetcherBackend {
	f := new(FetcherBackend)
	f.processors = make([]RowProcessor, 0, 10)
	return f
}

type RowProcessor interface {
	OnRow(row *core.Row)
}
