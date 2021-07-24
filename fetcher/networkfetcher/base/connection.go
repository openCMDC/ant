package base

import (
	"ant/fetcher/base"
	log "github.com/sirupsen/logrus"
	"net"
)

type StreamType int8

const (
	_ StreamType = iota
	Client2ServerStream
	Server2ClientStream
)

func (t StreamType) String() string {
	if t == Client2ServerStream {
		return "Client2ServerStream"
	} else if t == Server2ClientStream {
		return "Server2ClientStream"
	} else {
		return "unKnown stream type"
	}
}

type TCPConn struct {
	clientAddr          *net.TCPAddr
	serverAddr          *net.TCPAddr
	client2ServerStream *ConcurrentReader
	server2ClientStream *ConcurrentReader
	started             bool
	ctx                 *base.FetcherBackend
}

func (conn *TCPConn) GetClientAddr() *net.TCPAddr {
	return conn.clientAddr
}

func (conn *TCPConn) GetServerAddr() *net.TCPAddr {
	return conn.serverAddr
}

func (conn *TCPConn) C2SStream() *ConcurrentReader {
	return conn.client2ServerStream
}

func (conn *TCPConn) S2CStream() *ConcurrentReader {
	return conn.server2ClientStream
}

func (conn *TCPConn) CheckC2SStreamExist() bool {
	return conn.client2ServerStream != nil
}

func (conn *TCPConn) CheckS2CStreamExist() bool {
	return conn.server2ClientStream != nil
}

func (conn *TCPConn) SetC2SStream(stream *StreamReader) {
	if conn.client2ServerStream.reader != nil {
		log.Warn("replace client to server stream which already existed is not allowed")
		return
	}
	conn.client2ServerStream.reader = stream
	conn.client2ServerStream.SetReaderReady()
}

func (conn *TCPConn) IsConnIncomplete() bool {
	return conn.server2ClientStream.reader == nil || conn.client2ServerStream.reader == nil
}

func (conn *TCPConn) HasClosedStream() bool {
	return (conn.server2ClientStream.reader != nil && conn.server2ClientStream.reader.closed) ||
		(conn.client2ServerStream.reader != nil && conn.client2ServerStream.reader.closed)
}

func (conn *TCPConn) CheckBothStreamClosed() bool {
	return conn.server2ClientStream.reader != nil && conn.server2ClientStream.reader.closed &&
		conn.client2ServerStream.reader != nil && conn.client2ServerStream.reader.closed
}

func (conn *TCPConn) SetS2CStream(stream *StreamReader) {
	if conn.server2ClientStream.reader != nil {
		log.Warn("replace server to clietn stream which already existed is not allowed")
	}
	conn.server2ClientStream.reader = stream
	conn.server2ClientStream.SetReaderReady()
	return
}

func NewTCPConn(clientAddr, serverAddr *net.TCPAddr, c2sStream, s2cStream *StreamReader) *TCPConn {
	return &TCPConn{clientAddr: clientAddr,
		serverAddr:          serverAddr,
		client2ServerStream: NewConcurrentReader(c2sStream),
		server2ClientStream: NewConcurrentReader(s2cStream),
		started:             false}
}
