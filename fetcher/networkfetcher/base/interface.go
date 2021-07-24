package base

import (
	"ant/fetcher/base"
	"fmt"
)

type Interface interface {
	Protocol() string
	IsValidBeginningBytes(bytes []byte, streamType StreamType) ([]byte, bool)
	ParseClient2ServerStream(decider *Decider) error
	ParseServer2ClientStream(decider *Decider) error
}

type DecoderFactoryFunc func(backend *base.FetcherBackend) Interface

type MismatchedDecoderError struct {
	Protocol string
	Cause    error
	Data     []byte
}

func (m *MismatchedDecoderError) Error() string {
	return fmt.Sprintf("decode data {%s} to {%s} failed of {%s}", m.Data, m.Protocol, m.Cause)
}

func NewMismatchedDecoderError(protocol string, cause error, data []byte) error {
	err := new(MismatchedDecoderError)
	err.Protocol = protocol
	if cause != nil {
		err.Cause = cause
	} else {
		err.Cause = fmt.Errorf("Unknown")
	}

	if data != nil {
		err.Data = data
	} else {
		err.Data = data
	}
	return err
}
