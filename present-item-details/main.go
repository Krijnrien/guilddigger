package main

import (
	"flag"
	"log"
	"net"

	"github.com/nats-io/go-nats-streaming"
	"google.golang.org/grpc"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/krijnrien/guilddigger/nats"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/credentials"
)

const (
	defaultDBHost = "192.168.99.100"
	defaultPort   = ":50051"

	// NATS
	clusterID = "stan"
	natsURL   = "nats://nats:4222"
	clientID  = "event-store-api"
)

var (
	port   = flag.String("port", defaultPort, "grpc port number to listen.")
	dbHost = flag.String("dbHost", defaultDBHost, "IP of the cassandra database")
)

type server struct {
	*nats.StreamingComponent
}

func main() {
	logrus.SetLevel(logrus.DebugLevel)
	flag.Parse()

	log.Println("Listening on port ", *port)

	lis, err := net.Listen("tcp", *port)
	if err != nil {
		log.Fatalf("failed to listen on port %s: %v", *port, err)
	}

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
	// Create the TLS c
	c, err := credentials.NewServerTLSFromFile("configs/cert/servercert.pem", "configs/cert/serverkey.pem")
	if err != nil {
		log.Fatalf("could not load TLS keys: %s", err)
	}
	// Create an array of gRPC options with the c
	opts := []grpc.ServerOption{grpc.Creds(c)}
	// create a gRPC server object
	s := grpc.NewServer(opts...)

	RegisterItemPresentServer(s, &server{StreamingComponent: comp})
	err2 := s.Serve(lis)
	if err2 != nil {
		log.Fatal(err)
	}
}

// GetItems RPC gets items from present_item-details
func (s *server) GetItems(empty *empty.Empty, itemsServer ItemPresent_GetItemsServer) (error) {
	GetAllItemDetails(itemsServer)
	return nil
}

// GetItemIds RPC gets item ids from present_item-details
func (s *server) GetItemIds(empty *empty.Empty, itemsServer ItemPresent_GetItemIdsServer) (error) {
	StreamAllItemIds(itemsServer)
	return nil
}
