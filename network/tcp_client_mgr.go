package network

type TcpClientMgr struct {
	conn_mgr *TcpConnectionMgr
	clients  map[*TcpClient]*TcpClient
}

func (this *TcpClientMgr) Init(max_count uint32) bool {
	this.clients = make(map[*TcpClient]*TcpClient, max_count)
	this.conn_mgr = &TcpConnectionMgr{}
	return this.conn_mgr.Init(TCP_CONNECTION_TYPE_ACTIVE, max_count)
}

func (this *TcpClientMgr) NewClient(processor IConnProcessor) *TcpClient {
	conn := this.conn_mgr.NewConn3(processor)
	if conn == nil {
		return nil
	}
	c := &TcpClient{}
	this.clients[c] = c
	c.conn = conn
	return c
}

func (this *TcpClientMgr) FreeClient(c *TcpClient) bool {
	if _, o := this.clients[c]; !o {
		return false
	}
	delete(this.clients, c)
	this.conn_mgr.FreeConn(c.conn.id)
	return true
}
