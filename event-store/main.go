package main

import (
	"context"
	"flag"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	defaultDBHost = "cassandra"
	clusterID     = "test-cluster"
	clientID      = "event-store-api"
	natsURL       = "nats://example-nats:4222"
	stanURL       = "example-stan"

	defaultPort = ":50051"
)

var (
	port   = flag.String("port", defaultPort, "grpc port number to listen.")
	dbHost = flag.String("dbHost", defaultDBHost, "IP of the cassandra database")
)

type server struct {
	*StreamingComponent
}

func main() {
	log.Println("starting...")
	flag.Parse()

	lis, err := net.Listen("tcp", *port)
	if err != nil {
		log.Fatalf("failed to listen on port %s: %v", *port, err)
	}

	// Register new component within the NATS system.
	comp := NewStreamingComponent(clientID)

	// Connect to NATS
	err = comp.ConnectToNATSStreaming(clusterID, natsURL, )
	if err != nil {
		log.Fatal(err)
	}

	// Creates a new gRPC server
	// Create the TLS credentials
	creds, err := credentials.NewServerTLSFromFile("configs/cert/servercert.pem", "configs/cert/serverkey.pem")
	if err != nil {
		log.Fatalf("could not load TLS keys: %s", err)
	}
	// Create an array of gRPC options with the credentials
	opts := []grpc.ServerOption{grpc.Creds(creds)}
	// create a gRPC server object
	s := grpc.NewServer(opts...)
	//s := grpc.NewServer()

	RegisterEventStoreServer(s, &server{StreamingComponent: comp})
	err2 := s.Serve(lis)
	if err2 != nil {
		log.Fatal(err)
	}
}

// CreateEvent RPC creates a new Event into EventStore
// and publish an event to NATS Streaming
func (s *server) CreateEvent(ctx context.Context, in *Event) (*Response, error) {
	// Persist data into EventStore database
	//command := store.EventStore{}
	//// Persist events as immutable logs into CockroachDB
	//err := command.CreateEvent(in)

	CreateEvent(in)

	// Publish event on NATS Streaming Server
	go publishEvent(s.StreamingComponent, in)
	return &Response{IsSuccess: true}, nil
}

// GetEvents RPC gets events from EventStore by given AggregateId
func (s *server) GetEvents(ctx context.Context, in *EventFilter) (*EventResponse, error) {
	// TODO
	//eventStore := store.EventStore{}
	//events := eventStore.GetEvents(in)
	//return &EventResponse{Events: events}, nil
	return nil, nil
}

// publishEvent publishes an event via NATS Streaming server
func publishEvent(component *StreamingComponent, event *Event) {
	sc := component.NATS()
	channel := event.Channel
	eventMsg := []byte(event.EventData)
	// Publish message on subject (channel)
	err := sc.Publish(channel, eventMsg)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Published message on channel: " + channel)
}
