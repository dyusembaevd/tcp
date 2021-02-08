package tcp

import (
	"net"
	"time"
)

// TCPConfig ...
type TCPConfig struct {
	Hostname string
	Port     string
	Conn     net.Conn
}

func NewConfig() *TCPConfig {
	return &TCPConfig{}
}

func (t *TCPConfig) Connect(Hostname, Port string) error {
	t.Hostname = Hostname
	t.Port = Port
	err := t.connection()
	return err
}

func (t *TCPConfig) connection() error {
	Conn, err := net.Dial("tcp", t.Hostname+":"+t.Port)
	t.Conn = Conn
	return err
}

func (t *TCPConfig) ReadTCPMessage() []byte {
	timer := time.NewTimer(1 * time.Minute)
	buffer := make([]byte, 1024)
	go func() {
		t.Conn.Read(buffer)
	}()

LOOP:
	select {
	case <-timer.C:
		return []byte{}
	default:
		if buffer[0] != 0 {
			timer.Stop()
			break LOOP
		} else {
			goto LOOP
		}
	}
	return buffer
}

func (t *TCPConfig) WriteTCPMessage(message []byte) error {
	_, err := t.Conn.Write([]byte(message))
	return err
}
