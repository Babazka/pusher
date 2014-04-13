package main

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s : %s", msg, err)
	}
}

func listen(ch chan string) {
	log.Println("START")
	uri := "amqp://admin:Bsn7ixZTA6CD@localhost:5672/"

	connection, err := amqp.Dial(uri)
	failOnError(err, "Cannot connect")

	channel, err := connection.Channel()
	failOnError(err, "Cannot get channel")

	res, err := channel.Consume("hello", "", true, false, false, false, nil)
	failOnError(err, "Cannot consume")

	for {
		msg := <-res
		fmt.Println("Got message first")
		fmt.Printf("msg %+v\n", string(msg.Body))
		ch <- string(msg.Body)
	}
}

//func main() {
    //ch:=make(chan string)
    //go listen(ch)
    //for {
        //msg:= <-ch
		//fmt.Println("Got message too !")
		//fmt.Printf("msg %+v\n", msg)

    //}
//}

