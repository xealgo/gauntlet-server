package network

import (
	"net"
	"time"
)

const (
	ClientStateConnecting = iota + 1
	ClientStateActive     = iota
	ClientStateIdle       = iota
	ClientStateDisconnect = iota
)

// Client represents each connected player.
type Client struct {
	UUID     string
	Address  *net.UDPAddr
	LastPing time.Time
	State    uint
}
