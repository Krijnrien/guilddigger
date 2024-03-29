package main

import (
	"fmt"
	"github.com/nats-io/go-nats"
	"sync"

	"github.com/nats-io/go-nats-streaming"
	"github.com/nats-io/nuid"
)

// StreamingComponent is contains reusable logic related to handling
// of the connection to NATS Streaming in the system.
type StreamingComponent struct {
	// cmu is the lock from the component.
	cmu sync.Mutex

	// id is a unique identifier used for this component.
	id string

	// nc is the connection to NATS Streaming.
	nc stan.Conn

	// kind is the type of component.
	kind string
}

// NewStreamingComponent creates a StreamingComponent
func NewStreamingComponent(kind string) *StreamingComponent {
	id := nuid.Next()
	return &StreamingComponent{
		id:   id,
		kind: kind,
	}
}

// ConnectToNATSStreaming connects to NATS Streaming
func (c *StreamingComponent) ConnectToNATSStreaming(clusterID string, natsURL string, options ...stan.Option) error {
	c.cmu.Lock()
	defer c.cmu.Unlock()

	nc2, err := nats.Connect(natsURL)
	if err != nil {
		return err
	}
	nc, err := stan.Connect(clusterID, c.id, stan.NatsConn(nc2))
	if err != nil {
		return err
	}

	c.nc = nc
	return err
}

// NATS returns the current NATS connection.
func (c *StreamingComponent) NATS() stan.Conn {
	c.cmu.Lock()
	defer c.cmu.Unlock()
	return c.nc
}

// ID returns the ID from the component.
func (c *StreamingComponent) ID() string {
	c.cmu.Lock()
	defer c.cmu.Unlock()
	return c.id
}

// Name is the label used to identify the NATS connection.
func (c *StreamingComponent) Name() string {
	c.cmu.Lock()
	defer c.cmu.Unlock()
	return fmt.Sprintf("%s:%s", c.kind, c.id)
}

// Shutdown makes the component go away.
func (c *StreamingComponent) Shutdown() error {
	err := c.NATS().Close()
	if err != nil {
		return err
	}
	return nil
}
