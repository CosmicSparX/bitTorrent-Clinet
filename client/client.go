package client

import (
	"github.com/CosmicSparX/bitTorrent-Clinet/peers"
	"net"
	"time"
)

type Client struct {
	Conn   net.Conn
	Choked bool
}

func New(peer peers.Peer, peerID, infoHash [20]byte) (*Client, error) {
	conn, err := net.DialTimeout("tcp", peer.String(), 3*time.Second)
	if err != nil {
		return nil, err
	}

}
