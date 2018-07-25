package main

import (
	"log"
	"fmt"
	"flag"
	"os"
	"path"

	"github.com/quickfixgo/examples/cmd/tradeclient/internal"
	"github.com/quickfixgo/quickfix"
	"github.com/streadway/amqp"
)

//TradeClient implements the quickfix.Application interface
type TradeClient struct {
}

//OnCreate implemented as part of Application interface
func (e TradeClient) OnCreate(sessionID quickfix.SessionID) {
	return
}

//OnLogon implemented as part of Application interface
func (e TradeClient) OnLogon(sessionID quickfix.SessionID) {
	return
}

//OnLogout implemented as part of Application interface
func (e TradeClient) OnLogout(sessionID quickfix.SessionID) {
	return
}

//FromAdmin implemented as part of Application interface
func (e TradeClient) FromAdmin(msg *quickfix.Message, sessionID quickfix.SessionID) (reject quickfix.MessageRejectError) {
	return
}

//ToAdmin implemented as part of Application interface
func (e TradeClient) ToAdmin(msg *quickfix.Message, sessionID quickfix.SessionID) {
	return
}

//ToApp implemented as part of Application interface
func (e TradeClient) ToApp(msg *quickfix.Message, sessionID quickfix.SessionID) (err error) {
	fmt.Printf("Sending %s\n", msg)
	return
}

//FromApp implemented as part of Application interface. This is the callback for all Application level messages from the counter party.
func (e TradeClient) FromApp(msg *quickfix.Message, sessionID quickfix.SessionID) (reject quickfix.MessageRejectError) {
	fmt.Printf("FromApp: %s\n", msg.String())
	return
}
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"orders", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func(){
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)

			flag.Parse()

			cfgFileName := path.Join("config", "tradeclient.cfg")
			if flag.NArg() > 0 {
				cfgFileName = flag.Arg(0)
			}

			cfg, err := os.Open(cfgFileName)
			if err != nil {
				fmt.Printf("Error opening %v, %v\n", cfgFileName, err)
				return
			}

			appSettings, err := quickfix.ParseSettings(cfg)
			if err != nil {
				fmt.Println("Error reading cfg,", err)
				return
			}

			app := TradeClient{}
			fileLogFactory, err := quickfix.NewFileLogFactory(appSettings)

			if err != nil {
				fmt.Println("Error creating file log factory,", err)
				return
			}

			initiator, err := quickfix.NewInitiator(app, quickfix.NewMemoryStoreFactory(), appSettings, fileLogFactory)
			if err != nil {
				fmt.Printf("Unable to create Initiator: %s\n", err)
				return
			}

			initiator.Start()
			err = internal.QueryEnterOrder(string(d.Body))
			if err != nil {
				fmt.Printf("%v\n", err)
				}

			initiator.Stop()
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}