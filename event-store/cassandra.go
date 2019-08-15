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

func CreateEvent(event *Event) {
	err := session.Query("INSERT INTO item(item_id, item_name, icon, description, item_type, rarity, item_level, vendorvalue) VALUES(?, ?, ?, ?, ?, ?, ?, ?)", event.EventId, event.EventType, event.AggregateId, event.AggregateType, event.EventData, event.Channel).Exec()
	if err != nil {
		log.Error("Error while inserting item", err)
	}
}
