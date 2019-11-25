// Package network is for creating internetworks
package network

import (
	"time"

	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/server"
	"github.com/micro/go-micro/transport"
	"github.com/micro/go-micro/tunnel"
)

var (
	// DefaultName is default network name
	DefaultName = "go.micro"
	// DefaultAddress is default network address
	DefaultAddress = ":0"
	// ResolveTime defines time interval to periodically resolve network nodes
	ResolveTime = 1 * time.Minute
	// AnnounceTime defines time interval to periodically announce node neighbours
	AnnounceTime = 30 * time.Second
	// PruneTime defines time interval to periodically check nodes that need to be pruned
	// due to their not announcing their presence within this time interval
	PruneTime = 90 * time.Second
)

// Node is network node
type Node interface {
	// Id is node id
	Id() string
	// Address is node bind address
	Address() string
	// Peers returns node peers
	Peers() []Node
	// Network is the network node is in
	Network() Network
}

// Network is micro network
type Network interface {
	// Node is network node
	Node
	// Initialise options
	Init(...Option) error
	// Options returns the network options
	Options() Options
	// Name of the network
	Name() string
	// Connect starts the resolver and tunnel server
	Connect() error
	// Close stops the tunnel and resolving
	Close() error
	// Client is micro client
	Client() client.Client
	// Server is micro server
	Server() server.Server
}

// message is network message
type message struct {
	// msg is transport message
	msg *transport.Message
	// session is tunnel session
	session tunnel.Session
}

// NewNetwork returns a new network interface
func NewNetwork(opts ...Option) Network {
	return newNetwork(opts...)
}
