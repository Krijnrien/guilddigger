package main

import (
	"fmt"
	"sort"
	"flag"
	"time"
	"runtime"
	"encoding/json"
	"net/url"
	"io/ioutil"
	"crypto/x509"
	"net/http"
	"crypto/tls"

	"github.com/satori/go.uuid"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/krijnrien/guilddigger/gw2api"
	"github.com/krijnrien/guilddigger/micro_db"
	_ "github.com/krijnrien/guilddigger/rpc_resolver"
)

const (
	event     = "itemdetails-request"
	aggregate = "order"

	batchSize = 200
)

const (
	defaultServer     = "world"
	defaultServerAddr = "192.168.56.1:50051, 192.168.56.1:50052"

	exampleScheme      = "example"
	exampleServiceName = "www.example.com"
)

var (
	// Initializing gw2 api wrapper
	api = &gw2api.GW2Api{}

	// Initializing logging library
	log          = logrus.New()
	clientLogger = log.WithFields(logrus.Fields{"event": event})

	// Initialize GRPC dial options
	opts []grpc.DialOption

	// Flags
	server    = flag.String("server", defaultServer, "Name or IP of the target server, including port number.")
	enableLB  = flag.Bool("enable-load-balancing", false, "Set to true to enable client-side load balancing")
	serverIPs = flag.String("server-ipv4", defaultServerAddr, "If load balancing is enabled, this is a list of comma-separated server addresses used by the GRPC name resolver")
)

func main() {
	flag.Parse()

	// Create the client TLS transportCredentials
	transportCredentials, err := credentials.NewClientTLSFromFile("/usr/bin/configs/cert/servercert.pem", "")
	if err != nil {
		clientLogger.Fatal("could not load tls cert: %s", err)
	}
	opts = []grpc.DialOption{
		grpc.WithTransportCredentials(transportCredentials),
		grpc.WithBalancerName("round_robin"),
	}

	updateMissingItems()
	ticker := time.NewTicker(1 * time.Hour)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				updateMissingItems()
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
	runtime.Goexit()
}

func updateMissingItems() {
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

	conf := micro_db.MySQLConfig{
		Username: "",
		Password: "",
		Host:     "localhost",
		Port:     3306,
	}
	DB, configErr := conf.NewMySQLDB()
	if configErr != nil {
		clientLogger.Fatal("couldn't establish DB connection", configErr)
	}

	itemResponsesExist, err := DB.Item.ListItemIds()
	if err != nil {
		clientLogger.Error("couldn't retrieve already stored items, fetching all existing items", err)
	}

	clientLogger.WithFields(logrus.Fields{
		"length": len(itemResponsesExist),
	}).Info("retrieved list of item id's from database")

	IdsListFetched, err := api.Items()
	if err != nil {
		clientLogger.Error("couldn't retrieve list of item id's from gw2 api", err)
	}
	clientLogger.WithFields(logrus.Fields{
		"length": len(IdsListFetched),
	}).Info("fetched list of item id's from gw2 api")

	// Filter out all item ID's already exist in the database
	for _, itemResponse := range itemResponsesExist {
		i := sort.Search(len(IdsListFetched), func(i int) bool { return IdsListFetched[i] >= itemResponse.ItemId })
		if i < len(IdsListFetched) && IdsListFetched[i] == itemResponse.ItemId {
			IdsListFetched = append(IdsListFetched[:i], IdsListFetched[i+1:]...)
		}
	}

	// Create fetch commands
	// List of item ID's to fetch is divided by the set var batch size
	for i := 0; i < len(IdsListFetched); i += batchSize {
		batch := IdsListFetched[i:Min(i+batchSize, len(IdsListFetched))]
		createItemDetailsFetchCommand(batch)
	}
}

func createItemDetailsFetchCommand(itemIdsList []uint32) {
	clientLogger.WithFields(logrus.Fields{
		"batch_size": batchSize,
	}).Debug("fetch command created")

	itemIds := &ItemIdsRequest{
		BatchUUID: uuid.NewV4().String(),
		ItemIds:   itemIdsList,
	}

	err := createItemDetailsFetchCommandRPC(itemIds)
	if err != nil {
		log.Fatal(err)
	}
}

func createItemDetailsFetchCommandRPC(itemIds *ItemIdsRequest) error {
	conn, err := grpc.Dial(
		fmt.Sprintf("%s:///%s", exampleScheme, exampleServiceName), opts...
	)
	if err != nil {
		log.Fatalf("Unable to connect: %v", err)
	}
	defer conn.Close()

	client := NewEventStoreClient(conn)
	paymentJSON, _ := json.Marshal(itemIds)

	event := &Event{
		EventId:       uuid.NewV4().String(),
		EventType:     event,
		AggregateId:   itemIds.BatchUUID,
		AggregateType: aggregate,
		EventData:     string(paymentJSON),
		Channel:       event,
	}

	resp, err := client.CreateEvent(context.Background(), event)
	if err != nil {
		return errors.Wrap(err, "error from RPC server")
	}
	if resp.IsSuccess {
		return nil
	} else {
		return errors.Wrap(err, "error from RPC server")
	}
}

func Min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
