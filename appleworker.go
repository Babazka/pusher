// Package main provides ...
package main

import (
	"flag"
	"fmt"
)

var (
	uri          = flag.String("uri", "amqp://admin:Bsn7ixZTA6CD@localhost:5672/", "AMQP URI")
	exchangeName = flag.String("exchange", "test-exchange", "Durable AMQP exchange name")
	exchangeType = flag.String("exchange-type", "direct", "Exchange type - direct|fanout|topic|x-custom")
	routingKey   = flag.String("key", "test-key", "AMQP routing key")
	body         = flag.String("body", "foobar", "Body of message")
	reliable     = flag.Bool("reliable", true, "Wait for the publisher confirmation before exiting")
)

func init() {
	flag.Parse()
}




func main() {
	fmt.Printf("exchangeType %+v\n", *exchangeType)
	send()

}
