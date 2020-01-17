# go-distributed-sys

Check out the article [Building Microservices with Event Sourcing/CQRS in Go using gRPC, NATS Streaming and CockroachDB](https://medium.com/@shijuvar/building-microservices-with-event-sourcing-cqrs-in-go-using-grpc-nats-streaming-and-cockroachdb-983f650452aa)
## Technologies Used: 
* Go
* NATS Streaming
* gRPC
* ~~CockroachDB~~ Postgres


## Start nats-streaming
Run the command below from the nats-streaming directory:

`nats-streaming-server.exe --store file --dir ./data --max_msgs 0 --max_bytes 0`

## Start Postgres in Docker

### Start the Postgres server in Docker
`docker run -p 26257:5432 --name some-postgres -e POSTGRES_USER=shijuvar -e POSTGRES_DB=ordersdb -d postgres`

### Connect Postgres client command line: psql
`docker exec -it some-postgres psql ordersdb shijuvar`


## Start OrderService

### cd to the orderservice folder

`go get -v`

`go run main.go`

## Start EventStore

### cd to eventstore folder

`go get -v`

`go run main.go`

## Start PaymentService

### cd to paymentservice folder

`go get -v`

`go run main.go`

## Start orderquery-store1

### cd to orderquery-store1 folder

`go get -v`

`go run main.go`

## Start orderquery-store2

### cd to orderquery-store2 folder

`go get -v`

`go run main.go`

## Start restaurantservice

### cd to restaurantservice folder

`go get -v`

`go run main.go`

## Use Postman to send a OrderCreateCommand

### POST url: http://localhost:3000/api/orders

### Body like:

`{"customer_id" : "Google"}`

## Use curl to send a OrderCreateCommand

`curl -X POST -d '{"customer_id":"Google"}' http://localhost:3000/api/orders`

## Check the events table in Postgres

## Basic Workflow in the example:
1. A client app post an Order to an HTTP API.
2. An HTTP API (**orderservice**) receives the order, then executes a command onto Event Store, which is an immutable log of events, to create an event via its gRPC API (**eventstore**). 
3. The Event Store API executes the command and then publishes an event "order-created" to NATS Streaming server to let other services know that an event is created.
4. The Payment service (**paymentservice**) subscribes the event “order-created”, then make the payment, and then create an another event “order-payment-debited” via Event Store API. 
5. The Query syncing workers (**orderquery-store1 and orderquery-store2 as queue subscribers**) are also subscribing the event “order-created” that synchronise the data models to provide state of the aggregates for query views.
6. The Event Store API executes a command onto Event Store to create an event “order-payment-debited” and publishes an event to NATS Streaming server to let other services know that the payment has been debited.
7. The restaurant service (**restaurantservice**) finally approves the order.
8. A Saga coordinator manages the distributed transactions and makes void transactions on failures (to be implemented). 

