package main

import (
	"time"

	"github.com/gocql/gocql"
	log "github.com/sirupsen/logrus"
)

var session *gocql.Session

func init() {

	// Providing  cassandra cluster instance
	cluster := gocql.NewCluster(*dbHost)

	//// The authenticator is needed if password authentication is enabled
	//// for your Cassandra installation. If not, this can be commented.
	//cluster.Authenticator = gocql.PasswordAuthenticator{
	//	Username: "some_username",
	//	Password: "some_password",
	//}

	// gocql requires the keyspace to be provided before the session is created
	cluster.Keyspace = "itemdetails"
	var err error
	session, err = cluster.CreateSession()
	if err != nil {
		// Logging & entering Panic state.
		log.WithFields(log.Fields{
			"Error": err.Error(),
		}).Panic("Failed to create Cassandra DB getSession!")
	}

	cluster.ProtoVersion = 4
	cluster.Timeout = time.Second * 5

	// Get cassandra meta data from keyspace
	keySpaceMeta, _ := session.KeyspaceMetadata("itemdetails")

	// Check if item table exists, if not create table
	if _, exists := keySpaceMeta.Tables["item"]; exists != true {
		// Create item table
		err := session.Query("CREATE TABLE IF NOT EXISTS event (id text PRIMARY KEY, eventtype text, aggregateid text, aggregatetype text, eventdata text, channel text);").Exec()
		if err != nil {
			log.Error("Error while creating tables, ", err)
		}
	}
}

func GetAllItemDetails(itemsServer ItemPresent_GetItemsServer) {
	m := map[string]interface{}{}
	i := session.Query("SELECT item_id FROM item").Iter()
	for i.MapScan(m) {
		itemsServer.Send(&Item{
			ID:          uint32(m["item_id"].(int)),
			Name:        m["item_name"].(string),
			Icon:        m["icon"].(string),
			Description: m["description"].(string),
			Type:        m["item_type"].(string),
			Rarity:      m["rarity"].(string),
			Level:       uint32(m["item_level"].(int32)),
			VendorValue: uint32(m["vendorvalue"].(int32)),
		})
	}
}

func StreanmAllItemDetails() []Item {
	var items []Item
	m := map[string]interface{}{}
	i := session.Query("SELECT item_id FROM item").Iter()
	for i.MapScan(m) {
		items = append(items, Item{
			ID:          uint32(m["item_id"].(int)),
			Name:        m["item_name"].(string),
			Icon:        m["icon"].(string),
			Description: m["description"].(string),
			Type:        m["item_type"].(string),
			Rarity:      m["rarity"].(string),
			Level:       uint32(m["item_level"].(int32)),
			VendorValue: uint32(m["vendorvalue"].(int32)),
		})
		m = map[string]interface{}{}
	}
	return items
}

func GetAllItemIds() []ItemIds {
	var items []ItemIds
	m := map[string]interface{}{}
	i := session.Query("SELECT item_id FROM item").Iter()
	for i.MapScan(m) {
		items = append(items, ItemIds{
			ItemId: uint32(m["item_id"].(int)),
		})
		m = map[string]interface{}{}
	}
	return items
}

func StreamAllItemIds(itemsServer ItemPresent_GetItemIdsServer) {
	m := map[string]interface{}{}
	i := session.Query("SELECT item_id FROM item").Iter()
	for i.MapScan(m) {
		itemsServer.Send(&ItemIds{
			ItemId: uint32(m["item_id"].(int)),
		})
		m = map[string]interface{}{}
	}
}