package main

import (
	"github.com/gocql/gocql"
	"time"
	"github.com/sirupsen/logrus"
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
		log.WithFields(logrus.Fields{
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
		err := session.Query("CREATE TABLE item (item_id int PRIMARY KEY, item_name text, icon text, description text, item_type text, rarity text, item_level int, vendorvalue int);").Exec()
		if err != nil {
			log.Error("Error while creating tables, ", err)
		}
	}
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
