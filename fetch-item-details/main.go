package main

import (
	"flag"
	"time"
	"net/url"
	"runtime"
	"io/ioutil"
	"net/http"
	"crypto/tls"
	"crypto/x509"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"github.com/krijnrien/guilddigger/gw2api"
	"github.com/krijnrien/guilddigger/nats"
	"github.com/nats-io/go-nats-streaming"
	_ "github.com/krijnrien/guilddigger/rpc_resolver"
	"github.com/gogo/protobuf/proto"
)

const (
	// gRPC
	defaultServer     = "192.168.56.1:50051"
	defaultServerAddr = "192.168.56.1:50051" // contains more ip's, see flag

	// NATS server
	clusterID  = "test-cluster"
	natsURL    = "192.168.99.100:4222"
	channel    = "building-items"
	queueGroup = "order-query-store-group"
	durableID  = "store-durable"

	// NATS client
	clientID = "order-query-store1"

	exampleScheme      = "example"
	exampleServiceName = "www.example.com"
)

var addrs = []string{"192.168.99.100:50051"}

var (
	// Initializing gw2 api wrapper
	api = &gw2api.GW2Api{}

	// Initializing logging library
	log          = logrus.New()
	clientLogger = log.WithFields(logrus.Fields{"clientID": clientID, "channel": channel})

	// Initialize GRPC dial options
	opts []grpc.DialOption

	server    = flag.String("server", defaultServer, "Name or IP of the target server, including port number.")
	enableLB  = flag.Bool("enable-load-balancing", false, "Set to true to enable client-side load balancing")
	serverIPs = flag.String("server-ipv4", defaultServerAddr, "If load balancing is enabled, this is a list of comma-separated server addresses used by the GRPC name resolver")
)

func main() {
	// Set log level in logrus
	log.SetLevel(logrus.InfoLevel)
	clientLogger.Info("starting...")
	flag.Parse()

	// Create the client TLS transportCredentials
	transportCredentials, err := credentials.NewClientTLSFromFile("configs/cert/servercert.pem", "")
	if err != nil {
		clientLogger.Fatal("could not load tls cert: %s", err)
	}
	opts = []grpc.DialOption{
		grpc.WithTransportCredentials(transportCredentials),
		grpc.WithBalancerName("round_robin"),
	}

	//creating the proxyURL
	proxyStr := "http://192.168.99.100:3128"
	proxyURL, err := url.Parse(proxyStr)
	if err != nil {
		log.Println(err)
	}

	caCert, err := ioutil.ReadFile("configs/cert/cacert.pem")
	if err != nil {
		log.Fatal(err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	//adding the Transport object to the http Client
	client := http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs: caCertPool,
			},
			Proxy: http.ProxyURL(proxyURL),
		},
		Timeout: time.Duration(15 * time.Second),
	}
	api = gw2api.NewProxyGW2Api(&client)

	from := make(chan *ItemIds)
	to := make(chan []*ItemIds)

	var wait = time.After(time.Second * 5)
	var buffer = make([]*ItemIds, 0)

	// Register new NATS component within the system.
	comp := nats.NewStreamingComponent(clientID)

	// Connect to NATS Streaming server
	natsStreamingErr := comp.ConnectToNATSStreaming(
		clusterID,
		stan.NatsURL(natsURL),
	)
	if natsStreamingErr != nil {
		// Logging error with cluster details.
		clientLogger.WithFields(logrus.Fields{
			"clusterID":  clusterID,
			"nastURL":    natsURL,
			"channel":    channel,
			"queueGroup": queueGroup,
			"durableID":  durableID,
		}).Fatal("Could not connect to NATS")
	}

	// Get the NATS Streaming Connection
	sc := comp.NATS()

	// subscribing to channel via pool
	_, errSubscribe := sc.QueueSubscribe(channel, queueGroup, func(msg *stan.Msg) {
		var itemIds ItemIds
		err := proto.Unmarshal(msg.Data, &itemIds)
		if err != nil {
			clientLogger.WithFields(logrus.Fields{
				"message": msg.Data,
				"struct":  "ItemidsRequest",
			}).Error("Could not unmarshal message from queue into struct")
		} else {
			clientLogger.WithFields(logrus.Fields{
				"itemIds": itemIds,
			}).Debug("Received subscribed message")
			from <- &itemIds
			//TODO
			//fetchRequestedItems(itemIds)
		}
	}, stan.DurableName(durableID))
	if errSubscribe != nil {
		clientLogger.Fatal(errSubscribe)
	}

	go func() {
		for {
			fetchRequestedItems(<-to)
		}
	}()

	for {
		select {
		case <-wait:
			if len(buffer) > 0 {
				to <- buffer
				buffer = make([]*ItemIds, 0)
			}
			wait = time.After(time.Second * 5)
		case i := <-from:
			buffer = append(buffer, i)
			if len(buffer) == 15 {
				to <- buffer
				buffer = buffer[:0]
				wait = time.After(time.Second * 5)
			}
		}
	}

	runtime.Goexit()
}

func fetchRequestedItems(items []*ItemIds) {
	//
	ids := make([]uint32, len(items))
	for i, item := range items {
		ids[i] = item.ItemId
	}

	itemsPB, err := api.ItemIds("", ids...)
	if err != nil {
		clientLogger.Error(err)
	} else {
		log.Println(itemsPB)
		//createOrder(itemsPB)
	}
}

func createOrder(item []Item) {
	err := createOrderRPC(item)
	if err != nil {
		clientLogger.Fatal(err)
	}
}

// createOrderRPC calls the CreateEvent RPC
func createOrderRPC(item []Item) error {
	//conn, err := grpc.Dial(
	//	fmt.Sprintf("%s:///%s", exampleScheme, exampleServiceName), opts...
	//)
	////conn, err := grpc.Dial(*server, opts...)
	//if err != nil {
	//	log.Fatalf("Unable to connect: %v", err)
	//}
	//if err != nil {
	//	clientLogger.Fatal(err)
	//}
	//defer conn.Close()
	//client := NewEventStoreClient(conn)
	//orderJSON, err := json.Marshal(item)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//event := &Event{
	//	EventId:       uuid.NewV4().String(),
	//	EventType:     event,
	//	AggregateId:   item,
	//	AggregateType: aggregate,
	//	EventData:     string(orderJSON),
	//	Channel:       event,
	//}
	//resp, err := client.CreateEvent(context.Background(), event)
	//if err != nil {
	//	return errors.Wrap(err, "Error from RPC server")
	//}
	//if resp.IsSuccess {
	//	return nil
	//} else {
	//	return errors.Wrap(err, "Error from RPC server")
	//}
	return nil
}
