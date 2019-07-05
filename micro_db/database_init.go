package micro_db

//
//import (
//	"log"
//	"database/sql"
//	_ "github.com/lib/pq"
//)
//
//var DB *sql.DB
//
//func init() {
//	var err error
//	// Connect to the "ordersdb" database
//	DB, err = sql.Open("postgres", "postgresql://root@192.168.99.100:26257/ordersdb?sslmode=disable")
//	if err != nil {
//		log.Fatal("error connecting to the database: ", err)
//	}
//	createTables()
//}
//func createTables() {
//
//	if _, err := DB.Exec(
//		"CREATE DATABASE IF NOT EXISTS ordersdb;"); err != nil {
//		log.Fatal(err)
//	}
//
//	// Create the "events" table.
//	if _, err := DB.Exec(
//		"CREATE TABLE IF NOT EXISTS events (id string PRIMARY KEY, eventtype string, aggregateid string, aggregatetype string, eventdata string, channel string)"); err != nil {
//		log.Fatal(err)
//	}
//
//	// Create the "orders" table.
//	if _, err := DB.Exec(
//		"CREATE TABLE IF NOT EXISTS items (item_id int4 PRIMARY KEY, item_name string, icon string, description string, item_type string, rarity string, item_level int4, vendorvalue int4, defaultskin int4)"); err != nil {
//		log.Fatal(err)
//	}
//
//	// Create the "orderitems" table.
//	if _, err := DB.Exec(
//		"CREATE TABLE IF NOT EXISTS orderitems (id serial PRIMARY KEY, orderid string, customerid string, code string, name string, unitprice float, quantity int)"); err != nil {
//		log.Fatal(err)
//	}
//
//}
