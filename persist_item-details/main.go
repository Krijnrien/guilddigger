package main

import (
	"fmt"
	"log"
	"runtime"
	"encoding/json"

	"github.com/nats-io/go-nats-streaming"

	"github.com/krijnrien/guilddigger/micro_db"
	"github.com/krijnrien/guilddigger/nats"
)

const (
	clusterID  = "test-cluster"
	clientID   = "order-query-store1"
	channel    = "itemdetails-fetched"
	durableID  = "store-durable"
	queueGroup = "order-query-store-group"
	natsURL    = "192.168.99.100:4222"
)

func main() {
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

	_, err2 := sc.QueueSubscribe(channel, queueGroup, func(msg *stan.Msg) {
		log.Printf("[Orer] %s", string(msg.Data))
		var order ItemBatch
		err := json.Unmarshal(msg.Data, &order)

		if err == nil {
			var item []Item
			err := json.Unmarshal(order.ItemData, &item)
			if err == nil {
				// Handle the message
				log.Printf("Subscribed message from clientID - %s: %+v\n", clientID, item)

				err := DB.BulkCreateItem(item)
				if err != nil {
					log.Printf("Error while replicating the query model %+v", err)
				}
			} else {
				fmt.Println(err)
			}
		} else {
			fmt.Println(err)
		}
	}, stan.DurableName(durableID),
	)
	if err2 != nil {
		log.Fatal(err)
	}
	// prevents exiting, note application can now only exit with exit code 2 (either all goroutines completed, or a goroutine panic'd)
	runtime.Goexit()
}
