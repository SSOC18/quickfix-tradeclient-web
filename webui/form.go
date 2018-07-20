package main

import (
	"html/template"
	"net/http"
	"fmt"
)

type OrderDetails struct {
	ClOrdId string
	Price   string
	Symbol string
	Quantity string
}

func main() {
	tmpl := template.Must(template.ParseFiles("forms.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}

  	details := OrderDetails{
		ClOrdId: r.FormValue("clordid"),
		Price:   r.FormValue("price"),
		Symbol: r.FormValue("symbol"),
		Quantity: r.FormValue("quantity"),
	}

		// do something with details
		fmt.Println(details)
		_ = details

		tmpl.Execute(w, struct{ Success bool }{true})
	})

	http.ListenAndServe(":8080", nil)
}
