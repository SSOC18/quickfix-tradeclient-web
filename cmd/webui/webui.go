package main

import (
	"fmt"
	"log"
	"html/template"
	"net/http"
	"github.com/streadway/amqp"
	"path"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

type OrderDetails struct {
	Action string
	Version string
	ClOrdId string
	Price   string
	Symbol string
	OrderQty string
	Side string
	OrdType string
	TimeInForce string
	SenderCompID string
	TargetCompID string
	TargetSubID string
}

func main(){
	dir:="cmd/webui"
	tmpl := template.Must(template.ParseFiles(path.Join(dir,"form.html")))
	
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}

  	details := OrderDetails{
  		Action: r.FormValue("action"),
  		Version: r.FormValue("version"),
		ClOrdId: r.FormValue("clordid"),
		Price:   r.FormValue("price"),
		Symbol: r.FormValue("symbol"),
		OrderQty: r.FormValue("ordqty"),
		Side: r.FormValue("side"),
		OrdType: r.FormValue("ordtype"),
		TimeInForce: r.FormValue("timeinforce"),
		SenderCompID: r.FormValue("senderid"),
		TargetCompID: r.FormValue("targetid"),
		TargetSubID: r.FormValue("targetsubid"),
	}

	// do something with details
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

	body := details
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(fmt.Sprintf("%v", body)), //[]byte(details.Price)
		})
	log.Printf(" [x] Sent %s", body)
	failOnError(err, "Failed to publish a message")
	
	tmpl.Execute(w, struct{ Success bool }{true})
	
	})

	log.Fatal(http.ListenAndServe(":5004", nil))
} 
	