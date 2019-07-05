package main

import (
	"encoding/json"
	"fmt"
	"github.com/krijnrien/microguild/pkg/gw2api"
	"github.com/krijnrien/microguild/pkg/messages"
	"github.com/krijnrien/microguild/pkg/nats"
	"github.com/nats-io/go-nats-streaming"
	"github.com/pkg/errors"
	"github.com/satori/go.uuid"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"runtime"
)

const (
	natsURL   = "192.168.99.100:4222"
	event     = "commerce-prices-request"
	channel   = "commerce-prices-request"
	aggregate = "prices"

	clusterID  = "test-cluster"
	clientID   = "fetcher-commerce-prices-1"
	durableID  = "fetcher-commerce-prices-durable"
	queueGroup = "fetcher-commerce-prices-group"

	grpcUri = "192.168.56.1:50051"
)

var (
	api = gw2api.NewGW2Api()
)

func main() {
	// Register new NATS component within the system.
	comp := nats.NewStreamingComponent(clientID)

	// Connect to NATS Streaming server
	err := comp.ConnectToNATSStreaming(
		clusterID,
		stan.NatsURL(natsURL),
	)
	if err != nil {
		log.Fatal(err)
	}
	// Get the NATS Streaming Connection
	sc := comp.NATS()
	//TODO check what to do with return type
	_, err2 := sc.QueueSubscribe(channel, queueGroup, func(msg *stan.Msg) {
		pricesRequest := messages.PricesRequest{}
		err := json.Unmarshal(msg.Data, &pricesRequest)
		if err == nil {
			// Handle the message
			log.Printf("Subscribed message from clientID - %s: %+v\n", clientID, pricesRequest)
			fetchRequestedItems(pricesRequest)
		}
	}, stan.DurableName(durableID),
	)
	if err2 != nil {
		log.Fatal(err)
	}
	runtime.Goexit()
}

func fetchRequestedItems(pricesRequest messages.PricesRequest) {
	prices, _ := api.CommercePricePages(pricesRequest.Page, pricesRequest.PageSize)
	fmt.Println("", &prices)

	for _, itemPB := range prices {
		createOrder(itemPB)
	}
}

func createOrder(prices messages.Prices) {
	err := createOrderRPC(prices)
	if err != nil {
		log.Fatal(err)
	}
}

// createOrderRPC calls the CreateEvent RPC
func createOrderRPC(prices messages.Prices) error {
	conn, err := grpc.Dial(grpcUri, grpc.WithInsecure())
	if err != nil {
		log.Fatal("unable to connect: ", err)
	}
	defer conn.Close()
	client := messages.NewEventStoreClient(conn)
	orderJSON, _ := json.Marshal(prices)

	event := &messages.Event{
		EventId:       uuid.NewV4().String(),
		EventType:     event,
		AggregateId:   prices.AggregateId,
		AggregateType: aggregate,
		EventData:     string(orderJSON),
		Channel:       event,
	}
	resp, err := client.CreateEvent(context.Background(), event)
	if err != nil {
		fmt.Println(err)
		return errors.Wrap(err, "Error from RPC server")
	}
	if resp.IsSuccess {
		return nil
	} else {
		return errors.Wrap(err, "Error from RPC server")
	}
}
