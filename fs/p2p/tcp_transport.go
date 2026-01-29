package p2p

import (
	"fmt"
	"net"
	"sync"
)

type TCPPeer struct {
	conn     net.Conn
	outbound bool
}

type TCPTransportOpts struct {
	ListenAddr    string
	HandshakeFunc HandShakeFunc
	Decoder       Decoder
}
type TCPTransport struct {
	TCPTransportOpts
	Listener net.Listener
	mu       sync.RWMutex
	peers    map[net.Addr]Peer
}

func NewTCPPeer(conn net.Conn, outbound bool) *TCPPeer {
	return &TCPPeer{
		conn:     conn,
		outbound: outbound,
	}
}

func NewTCPTransport(opts TCPTransportOpts) *TCPTransport {
	return &TCPTransport{
		TCPTransportOpts: opts,
	}
}

func (t *TCPTransport) ListenAndAccept() error {
	var err error
	t.Listener, err = net.Listen("tcp", t.ListenAddr)
	if err != nil {
		return err
	}
	go t.startAcceptLoop()
	return nil
}

func (t *TCPTransport) startAcceptLoop() {
	for {
		conn, err := t.Listener.Accept()
		if err != nil {
			fmt.Printf("TCP Accept err: %s", err)
		}

		go t.handleConn(conn)
	}
}

func (t *TCPTransport) handleConn(conn net.Conn) {
	peer := NewTCPPeer(conn, true)
	if err := t.HandshakeFunc(peer); err != nil {
		conn.Close()
		fmt.Printf("tcp hadshake error :%s\n", err)

		return
	}
	// msg := &Message{}
	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Printf("TCP error %s\n", err)
		}
		// if err := t.Decoder.Decode(conn, msg); err != nil {
		// 	fmt.Printf("TCP error %s\n", err)
		// 	continue
		// }
		fmt.Printf("Message : %+v\n", buf[:n])
	}
}
