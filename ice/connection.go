package ice

type ConnectionInfo struct {
	Incoming     bool
	AdapterName  string
	ConnectionId string
	RcvSize      int
	SndSize      int
}
type ConnectionCallback interface {
	Heartbeat(con Connection)
	Closed(con Connection)
}

type Connection interface {
	Close(force bool) error
	CreateProxy(id Identity) ObjectPrx
	SetAdapter(adapter ObjectAdapter)
	GetAdapter() ObjectAdapter
	GetEndpoint() Endpoint
	FlushBatchRequests()
	SetCallback(callback ConnectionCallback)
	// SetACM()
	// GetACM() ACM
	Type() string
	Timeout() int
	ToString() string
	GetInfo() ConnectionInfo
	SetBufferSize(rcvSize int, sndSize int)
}
