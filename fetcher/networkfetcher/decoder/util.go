package decoder

import "ant/fetcher/networkfetcher/base"

type namedReader struct {
	name   string
	reader *base.ConcurrentReader
}

func NewNamedReader(name string, reader *base.ConcurrentReader) *namedReader {
	return &namedReader{
		name:   name,
		reader: reader,
	}
}

func (h *namedReader) Read(p []byte) (int, error) {
	return h.reader.Read(p, h.name)
}
