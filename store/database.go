package store

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB

func init() {
	log.Println("initialize the database, create tables")
	var err error
	// Connect to the "ordersdb" database
	db, err = sql.Open("postgres", "postgresql://shijuvar@localhost:26257/ordersdb?sslmode=disable")
	if err != nil {
		log.Fatal("error connecting to the database: ", err)
	}
	createTables()
}
func createTables() {

	// Create the "events" table.
	if _, err := db.Exec(
		"CREATE TABLE IF NOT EXISTS events (id varchar(255) PRIMARY KEY, eventtype varchar(255), aggregateid varchar(255), aggregatetype varchar(255), eventdata varchar(255), channel varchar(255))"); err != nil {
		log.Fatalf("Failed to create table events: %v", err)
	}

	// Create the "orders" table.
	if _, err := db.Exec(
		"CREATE TABLE IF NOT EXISTS orders (id varchar(255) PRIMARY KEY, customerid varchar(255), status varchar(255), createdon int, restaurantid varchar(255), amount float)"); err != nil {
		log.Fatalf("Failed to create table orders %v", err)
	}

	// Create the "orderitems" table.
	if _, err := db.Exec(
		"CREATE TABLE IF NOT EXISTS orderitems (id serial PRIMARY KEY, orderid varchar(255), customerid varchar(255), code varchar(255), name varchar(255), unitprice float, quantity int)"); err != nil {
		log.Fatalf("Failed to create table orderitems: %v", err)
	}

}
