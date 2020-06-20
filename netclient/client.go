package netclient

import (
	"bufio"
	"fmt"
	"net"
)

type Command struct {
}
type Client interface {
	Read() ([]byte, error)
	Send(b []byte) error
}

type SocketClient struct {
	conn   *net.TCPConn
	reader *bufio.Reader
	writer *bufio.Writer
}

func NewSocketClient(ipAddr, port string) (*SocketClient, error) {
	tcpAddr, err := net.ResolveTCPAddr("tcp", ipAddr+":"+port)
	if err != nil {
		return nil, fmt.Errorf("Error in creating connection. Error : %+v", err)
	}
	c := &SocketClient{}
	c.conn, err = net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		return nil, fmt.Errorf("Error in connection to server. Error : %+v", err)
	}
	c.reader = bufio.NewReader(c.conn)
	c.writer = bufio.NewWriter(c.conn)
	return c, nil
}

func (sc *SocketClient) Read() ([]byte, error) {
	data := make([]byte, 1024)
	_, err := sc.reader.Read(data)
	if err != nil {
		return nil, fmt.Errorf("Error in reading from server. Error : %+v", err)
	}
	return data, nil
}

func (sc *SocketClient) Send(b []byte) error {
	_, err := sc.writer.Write(b)
	if err != nil {
		return fmt.Errorf("Error in writing to server. Error : %+v", err)
	}
	return nil
}
