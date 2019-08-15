package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"flag"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"runtime"
	"sort"
	"time"

	"github.com/gogo/protobuf/proto"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/krijnrien/guilddigger/gw2api"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	defaultDBHost = "cassandra"

	channel = "building-items"

	// NATS
	clusterID = "stan"
	//TODO maybe get the POD ID from environment variable??
	clientID  = "order_item-details-1"
	natsURL   = "nats://nats:4222"

	grpcUri = "present-item-details:50051"
)

var (
	// Initializing gw2 api wrapper
	api = &gw2api.GW2Api{}

	// Initializing logging library
	log = logrus.New()

	clientLogger = log.WithFields(logrus.Fields{"channel": channel})

	// Flags
	//scheme      = flag.String("scheme", defaultScheme, "gRPC scheme used")
	//serviceName = flag.String("service-name", defaultServiceName, "gRPC service name")
	dbHost = flag.String("dbHost", defaultDBHost, "IP of the cassandra database")
)

type server struct {
	*StreamingComponent
}

func main() {
	log.SetLevel(logrus.DebugLevel)
	flag.Parse()

	log.Info("App started...3")

	// Register new component within the NATS system.
	comp := NewStreamingComponent(clientID)

	// Connect to NATS
	err := comp.ConnectToNATSStreaming(clusterID, natsURL)
	if err != nil {
		log.Fatal(err)
	}

	apple := &server{StreamingComponent: comp}

	clientLogger.Info("Passed connecting to NATS: ", apple)

	apple.updateMissingItems()
	ticker := time.NewTicker(1 * time.Hour)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				apple.updateMissingItems()
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
	runtime.Goexit()
}

func (s *server) updateMissingItems() {

	// Create the client TLS transportCredentials
	transportCredentials, err := credentials.NewClientTLSFromFile("configs/cert/servercert.pem", "")
	if err != nil {
		clientLogger.Fatal("could not load tls cert: %s", err)
	}
	var opts = []grpc.DialOption{
		grpc.WithTransportCredentials(transportCredentials),
		grpc.WithBalancerName("round_robin"),
	}

	conn, err := grpc.Dial(grpcUri, opts...)
	if err != nil {
		log.Fatalf("Unable to connect: %v", err)
	}
	defer conn.Close()

	//creating the proxyURL
	proxyStr := "http://cuttle:3128"
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
	client2 := http.Client{
		Transport: &http.Transport{
			//TODO need new certificate pointing to kubernetes instead of old docker IP. see original cuttle readme why
			TLSClientConfig: &tls.Config{
				RootCAs: caCertPool,
			},
			Proxy: http.ProxyURL(proxyURL),
		},
		Timeout: time.Duration(15 * time.Second),
	}
	api = gw2api.NewProxyGW2Api(&client2)
	fIds, err := api.Items()
	if err != nil {
		clientLogger.Error("couldn't retrieve list of item id's from gw2 api", err)
	}
	clientLogger.WithFields(logrus.Fields{"length": len(fIds),}).Info("fetched list of item id's from gw2 api")



	client := NewItemPresentClient(conn)
	resp, err := client.GetItemIds(context.Background(), &empty.Empty{})
	if err != nil {
		log.Error("error from RPC server, ", err)
	}

	done := make(chan bool)
	//var itemIds []ItemIds

	func() {
		for {
			resp, err := resp.Recv()
			if err == io.EOF {
				close(done)
				return
			}
			if err != nil {
				log.Fatalf("can not receive %v", err)
			}
			i := sort.Search(len(fIds), func(i int) bool { return fIds[i] >= resp.ItemId })
			if i < len(fIds) && fIds[i] == resp.ItemId {
				fIds = append(fIds[:i], fIds[i+1:]...)
			}
		}
	}()

	<-done

	var mIds []*ItemIds
	for i := 0; i < len(fIds); i++ {
		mIds = append(mIds, &ItemIds{ItemId: fIds[i]})
	}

	for i := 0; i < len(mIds); i++ {
		go createFetchCommand(s.StreamingComponent, mIds[i])
	}
}

func createFetchCommand(component *StreamingComponent, itemIds *ItemIds) {
	itemMsg, err := proto.Marshal(itemIds)
	if err != nil {
		log.Fatal(err)
	}

	clientLogger.WithFields(logrus.Fields{
		"item_id": itemIds,
	}).Debug("fetch command created")

	sc := component.NATS()
	// Publish message on subject (channel)
	err2 := sc.Publish(channel, []byte(itemMsg))
	if err != nil {
		log.Fatal(err2)
	}
	log.Debug("Published message on channel: " + channel)
}
