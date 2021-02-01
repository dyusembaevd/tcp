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
	conn, err := net.Dial("tcp", t.hostname+":"+t.port)
	// conn, err := net.DialTimeout("tcp", t.hostname+":"+t.port, 1*time.Minute)
	t.conn = conn
	return err
}

func (t *TCPConfig) ReadTCPMessage() []byte {
	conn := t.conn
	timer := time.NewTimer(1 * time.Minute)
	out := make(chan string, 1)
	buffer := make([]byte, 1024)
	go func(foo chan string) {
		conn.Read(buffer)
		foo <- "Done"
	}(out)
	select {
	case <-timer.C:
		return []byte{}
	case <-out:
		timer.Stop()
		return buffer
	}
}

func (t *TCPConfig) WriteTCPMessage(message []byte) error {
	_, err := t.conn.Write([]byte(message))
	return err
}
