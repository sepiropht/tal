package p2p

import (
  "net"
  "sync"
  "fmt"
)

type TCPTransport struct {
  listenAddress string
  listener      net.Listener

  mu sync.RWMutex
  peers map[net.Addr]Peer
}

func NewTCPTransport(listenAddr string) *TCPTransport {
  return &TCPTransport {
    listenAddress: listenAddr,
  }
}

func (t *TCPTransport) listenAndAccept() error {
  var err error 
  t.listener, err = net.Listen("tcp", t.listenAddress)
  if err != nil {
    return err
  }
  go t.startAcceptLoop()
  return nil
} 


func (t *TCPTransport) startAcceptLoop()  {
  for {
    conn, err := t.listener.Accept()
    if err != nil {
      fmt.Printf("TCP accpet error: %s\n", err)
    }
    go t.handleConn(conn)
  }
}

func (t *TCPTransport) handleConn(conn net.Conn) {
  fmt.Printf("new incomming connection %+v\n", conn)
}

