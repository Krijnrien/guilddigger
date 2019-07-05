package main

import (
	"encoding/json"
	"fmt"
	"log"
	"runtime"
	"strconv"
	"time"

	"github.com/krijnrien/microguild/pkg/gw2api"
	"github.com/krijnrien/microguild/pkg/messages"

	"github.com/pkg/errors"
	"github.com/satori/go.uuid"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	event     = "commerce-prices-request"
	channel   = "commerce-prices-request"
	aggregate = "prices"

	grpcUri = "192.168.56.1:50051"

	batchSize = 50
)

var (
	api = gw2api.NewGW2Api()
)

func main() {
	fetchCommerceListings()
	ticker := time.NewTicker(1 * time.Minute)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				fetchCommerceListings()
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
	runtime.Goexit()
}

func fetchCommerceListings() {
	commercePricePageHeaders, err := api.CommercePricePageHeaders(1, batchSize)
	if err != nil {
		log.Fatal(err)
	}

	pageTotalHeader := commercePricePageHeaders.Header.Get("x-page-total")
	pageTotal, err := strconv.Atoi(pageTotalHeader)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < pageTotal; i++ {
		time.Sleep(2 * time.Second)
		createPricesRequest(uint32(i))
	}
}

func createPricesRequest(pageNr uint32) {

	buysRequest := messages.PricesRequest{
		Page:     pageNr,
		PageSize: batchSize,
	}
	err := createPricesRequestRPC(buysRequest)
	if err != nil {
		log.Fatal(err)
	}
}

// createOrderRPC calls the CreateEvent RPC
func createPricesRequestRPC(item messages.PricesRequest) error {
	conn, err := grpc.Dial(grpcUri, grpc.WithInsecure())
	if err != nil {
		log.Fatal("unable to connect: ", err)
	}
	defer conn.Close()
	client := messages.NewEventStoreClient(conn)
	fmt.Println(item)
	orderJSON, _ := json.Marshal(item)

	event := &messages.Event{
		EventId:       uuid.NewV4().String(),
		EventType:     event,
		AggregateId:   uuid.NewV4().String(),
		AggregateType: aggregate,
		EventData:     string(orderJSON),
		Channel:       channel,
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
