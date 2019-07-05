package main

// gRPC API for Event Store

import (
	"context"
	"log"
	"github.com/krijnrien/microguild/pkg/messages"
	"github.com/krijnrien/microguild/pkg/micro_db"
	"github.com/krijnrien/microguild/pkg/nats"
	"net"
	"github.com/nats-io/go-nats-streaming"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc"
	"flag"
)

const (
	clusterID   = "test-cluster"
	clientID    = "event-store-api"
	natsURL     = "192.168.99.100:4222"
	defaultPort = ":50051"
)

var (
	port = flag.String("port", defaultPort, "grpc port number to listen.")
)

type server struct {
	*nats.StreamingComponent
}

func main() {
	flag.Parse()

	lis, err := net.Listen("tcp", *port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Println("Listening on: %v", *port)

	// Register new component within the NATS system.
	comp := nats.NewStreamingComponent(clientID)

	// Connect to NATS
	err = comp.ConnectToNATSStreaming(
		clusterID,
		stan.NatsURL(natsURL),
	)
	if err != nil {
		log.Fatal(err)
	}

	// Creates a new gRPC server
	// Create the TLS credentials
	creds, err := credentials.NewServerTLSFromFile("/usr/bin/configs/cert/servercert.pem", "/usr/bin/configs/cert/serverkey.pem")
	if err != nil {
		log.Fatalf("could not load TLS keys: %s", err)
	}
	// Create an array of gRPC options with the credentials
	opts := []grpc.ServerOption{grpc.Creds(creds)}
	// create a gRPC server object
	s := grpc.NewServer(opts...)
	//s := grpc.NewServer()

	messages.RegisterEventStoreServer(s, &server{StreamingComponent: comp})
	err2 := s.Serve(lis)
	if err2 != nil {
		log.Fatal(err)
	}
}

// CreateEvent RPC creates a new Event into EventStore
// and publish an event to NATS Streaming
func (s *server) CreateEvent(ctx context.Context, in *messages.Event) (*messages.Response, error) {
	// Persist data into EventStore database
	//command := store.EventStore{}
	//// Persist events as immutable logs into CockroachDB
	//err := command.CreateEvent(in)

	conf := micro_db.MySQLConfig{
		Username: "",
		Password: "",
		Host:     "localhost",
		Port:     3306,
	}

	DB, configErr := conf.NewMySQLDB()
	if configErr != nil {
		log.Fatalln(configErr)
	}

	err2 := DB.Event.CreateEvent(in)
	if err2 != nil {
		return nil, err2
	}
	// Publish event on NATS Streaming Server
	go publishEvent(s.StreamingComponent, in)
	return &messages.Response{IsSuccess: true}, nil
}

// GetEvents RPC gets events from EventStore by given AggregateId
func (s *server) GetEvents(ctx context.Context, in *messages.EventFilter) (*messages.EventResponse, error) {
	// TODO
	//eventStore := store.EventStore{}
	//events := eventStore.GetEvents(in)
	//return &messages.EventResponse{Events: events}, nil
	return nil, nil
}

// publishEvent publishes an event via NATS Streaming server
func publishEvent(component *nats.StreamingComponent, event *messages.Event) {
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
