package tcp

import (
	"net"
	"time"
)

// TCPConfig ...
type TCPConfig struct {
	hostname string
	port     string
	conn     net.Conn
}

func NewConfig() *TCPConfig {
	return &TCPConfig{}
}

func (t *TCPConfig) Connect(hostname, port string) error {
	t.hostname = hostname
	t.port = port
	err := t.connection()
	return err
}

func (t *TCPConfig) connection() error {
	conn, err := net.DialTimeout("tcp", t.hostname+":"+t.port, time.Minute)
	t.conn = conn
	return err
}

func (t *TCPConfig) ReadTCPMessage() []byte {
	conn := t.conn
	buffer := make([]byte, 1024)
	conn.Read(buffer)
	return buffer
}

func (t *TCPConfig) WriteTCPMessage(message []byte) error {
	_, err := t.conn.Write([]byte(message))
	return err
}
