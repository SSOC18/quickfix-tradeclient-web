package main

import (
	"fmt"
	"log"
    "path"
	"html/template"
	"net/http"
	"github.com/streadway/amqp"
	"io/ioutil"
///	"database/sql"
///	"github.com/lib/pq"
///	"github.com/shopspring/decimal"

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

type Page struct {
	Title string
 	Body  []byte
}
  
func loadPage(title string) (*Page, error) {
    dir := "cmd/webui"
	filename := title + ".txt"
    body, err := ioutil.ReadFile(path.Join(dir, filename))
  	if err != nil {
  		return nil, err
  	}

  	return &Page{Title: title, Body: body}, nil
}

func main(){
    dir := "cmd/webui"
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

	fmt.Println(q.Name)

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

	//fmt.Println("Connecting to cryppro_v0 database")
    //connStr := "user=mickael dbname=cryppro_v0 password=r5vPg3Q8 host=localhost port=postgresql"
    //db, err1 := sql.Open("postgres", connStr)
	//if err1 != nil {
	//	log.Fatal(err1)
	//}
    //_ = pq.Efatal
    //var minimum_ask string
    //var maximum_bid string
    //err2 := db.QueryRow("SELECT minimum_ask, maximum_bid FROM btcprice ORDER BY index DESC LIMIT 1;").Scan(&minimum_ask, &maximum_bid)
    //fmt.Println("Mimimum Ask    |   Maximum Bid")
    //fmt.Println(minimum_ask, maximum_bid)
    //if err2 != nil {
	//	log.Fatal(err2)
	//}
    
    //minask, _ := decimal.NewFromString(minimum_ask)
    //maxbid, _ := decimal.NewFromString(maximum_bid)
    //best_bidask1 := decimal.Avg(minask, maxbid)
    //best_bidask := best_bidask1.String()
    

    ///price := []byte(best_bidask)
    ///err := ioutil.WriteFile(path.Join(dir,"btcusd.txt"), price, 0644)
    ///if err != nil {
  	///		fmt.Println("Can't write new price to file text")
  	///		return
  	///	}

    ///tmpl1 := template.Must(template.ParseFiles(path.Join(dir,"view.html")))
	///http.HandleFunc("/view/", func(w http.ResponseWriter, r *http.Request) {
  	///	p, err := loadPage("btcusd")
  	///	if err != nil {
  	///		fmt.Println("Can't access btcusd.txt for price")
  	///		return
  	///	}

  	///	err4 := tmpl1.ExecuteTemplate(w, "view.html", p)
  	///	if err4 != nil {
  	///		http.Error(w, err4.Error(), http.StatusInternalServerError)
  	///	}

  	///	if r.Method != http.MethodPost {
	///		tmpl1.Execute(w, nil)
	///		return
	///	}

	///})

	log.Fatal(http.ListenAndServe(":5004", nil))
} 