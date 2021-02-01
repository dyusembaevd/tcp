package tcp

import (
	"fmt"
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
	timer := time.NewTimer(1 * time.Minute)
	buffer := make([]byte, 1024)
	go func(buf *[]byte) {
		fmt.Println("Start reading in goroutine")
		t.conn.Read(*buf)
		fmt.Println("Got message in goroutine")
	}(&buffer)

LOOP:
	select {
	case <-timer.C:
		fmt.Println("message not found :(")
		return []byte{}
	default:
		if len(string(buffer)) != 0 {
			timer.Stop()
			break LOOP
		}
	}
	return buffer
}

func (t *TCPConfig) WriteTCPMessage(message []byte) error {
	_, err := t.conn.Write([]byte(message))
	return err
}
