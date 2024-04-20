package main

import "net"

type Peer struct {
	conn  net.Conn
	msgCh chan []byte
}

func NewPeer(conn net.Conn, msgCh chan []byte) *Peer {
	return &Peer{
		conn:  conn,
		msgCh: msgCh}
}

func (p *Peer) readLoop() error {
	buffer := make([]byte, 1024)

	for {
		n, err := p.conn.Read(buffer)
		if err != nil {
			return err
		}

		msgBuffer := make([]byte, n)
		copy(msgBuffer, buffer[:n])
		p.msgCh <- msgBuffer
	}
}
