package main

//import (
	//"flag"
	//"fmt"
	//"github.com/streadway/amqp"
	//"log"
	//"os"
//)

//var (
	//uri          = flag.String("uri", "amqp://admin:Bsn7ixZTA6CD@localhost:5672/", "AMQP URI")
	//exchangeName = flag.String("exchange", "test-exchange", "Durable AMQP exchange name")
	//exchangeType = flag.String("exchange-type", "direct", "Exchange type - direct|fanout|topic|x-custom")
	//routingKey   = flag.String("key", "test-key", "AMQP routing key")
	//body         = flag.String("body", "foobar", "Body of message")
	//reliable     = flag.Bool("reliable", true, "Wait for the publisher confirmation before exiting")
//)

//func init() {
	//flag.Parse()
//}

//func failOnError(err error, msg string) {
	//if err != nil {
		//fmt.Printf(msg)
		//os.Exit(1)
	//}
//}

//func main() {
	//log.Println("START")
	//fmt.Print("TEST\n")
	//fmt.Printf("URI: %q\n", *uri)

	//connection, err := amqp.Dial(*uri)
	//fmt.Printf("err %+v\n", err)
	//failOnError(err, fmt.Sprint("Dial: %s\n", err))
	//defer connection.Close()
	//fmt.Println("Got connection")

	//channel, err := connection.Channel()
	//failOnError(err, fmt.Sprint("Get channel: %s\n", err))
	//fmt.Println("Got chanel")

	//queue, err := channel.QueueDeclare(
		//"hello", // name
		//false,   // durable
		//false,   // delete when usused
		//false,   // exclusive
		//false,   // noWait
		//nil,     // arguments
	//)

	//fmt.Println("Got queue")

	//failOnError(err, "Error on queue creating ")

	//err = channel.Publish(
		//"",         // exchange
		//queue.Name, // routing key
		//false,      // mandatory
		//false,      // immediate
		//amqp.Publishing{
			//ContentType: "text/plain",
			//Body:        []byte(*body),
		//})
	//fmt.Println("Published")
//}
