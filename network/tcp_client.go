package network

import (
	"fmt"
	"net"
)

type TcpClient struct {
	conn     TcpConnection
	conn_mgr *TcpConnectionMgr
}

func (this *TcpClient) Init(conn_mgr *TcpConnectionMgr) {
	this.conn_mgr = conn_mgr
}

func (this *TcpClient) Connect(addr string) bool {
	var err error
	var conn net.Conn
	conn, err = net.Dial("tcp", addr)
	if err != nil {
		fmt.Printf("connect failed\n")
		return false
	}

	this.conn.Init(conn)
	this.conn.Start()

	return true
}

func (this *TcpClient) Send(data []byte) error {
	return this.conn.Send(data)
}

func (this *TcpClient) Close() {
	this.conn.Close()
}
