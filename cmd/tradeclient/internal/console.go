package internal

import (
	"fmt"
	"time"
	"bufio"
	"os"

	"github.com/quickfixgo/enum"
	"github.com/quickfixgo/field"
	"github.com/quickfixgo/quickfix"
	"github.com/shopspring/decimal"

	"strings"

	fix50nos "github.com/quickfixgo/fix50/newordersingle"

)
func queryString(fieldName string) string {
	fmt.Printf("%v: ", fieldName)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return scanner.Text()
}

func queryDecimal(fieldName string) decimal.Decimal {
	val, err := decimal.NewFromString(queryString(fieldName))
	if err != nil {
		panic(err)
	}

	return val
}

func queryClOrdID(clordid string) field.ClOrdIDField {
	fmt.Println("ClOrdID: "+clordid)
	return field.NewClOrdID(clordid)
}

func querySide(side string) field.SideField {
	fmt.Println("Side: "+side)
	return field.NewSide(enum.Side(side))
}

func queryOrdType(f *field.OrdTypeField, ordtype string) field.OrdTypeField {
	fmt.Println("OrdType: "+ordtype)
	f.FIXString = quickfix.FIXString(ordtype)
	return *f
}

func querySymbol(symbol string) field.SymbolField {
	fmt.Println("Symbol: "+symbol)
	return field.NewSymbol(symbol)
}

func queryOrderQty(orderqty string) field.OrderQtyField {
	fmt.Println("OrderQty: "+orderqty)
	orderqty2, _ := decimal.NewFromString(orderqty)
	return field.NewOrderQty(orderqty2, 2)
}

func queryTimeInForce(timeinforce string) field.TimeInForceField {
	fmt.Println("TimeInForce: "+timeinforce)
	return field.NewTimeInForce(enum.TimeInForce(timeinforce))
}

func queryPrice(price string) field.PriceField {
	fmt.Println("Price: "+price)
	price2, _ := decimal.NewFromString(price)
	return field.NewPrice(price2, 2)
}

//NOT DEFINED FOR STOP PRICES YET 
//func queryStopPx() field.StopPxField {
//	return field.NewStopPx(queryDecimal("Stop Price"), 2)
//}

func querySenderCompID() field.SenderCompIDField {
	return field.NewSenderCompID(queryString("SenderCompID"))
}

func queryTargetCompID() field.TargetCompIDField {
	return field.NewTargetCompID(queryString("TargetCompID"))
}

func queryTargetSubID() field.TargetSubIDField {
	return field.NewTargetSubID(queryString("TargetSubID"))
}

func queryConfirm(prompt string) bool {
	fmt.Println()
	fmt.Printf("%v?: ", prompt)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	return strings.ToUpper(scanner.Text()) == "Y"
}

type header interface {
	Set(f quickfix.FieldWriter) *quickfix.FieldMap
}

func queryHeader(h header) {
	h.Set(querySenderCompID())
	h.Set(queryTargetCompID())
	if ok := queryConfirm("Use a TargetSubID"); !ok {
		return
	}

	h.Set(queryTargetSubID())
}

func queryNewOrderSingle50(neworder string) (msg *quickfix.Message) {
	var ordType field.OrdTypeField

	neworder2 := strings.Split(neworder, " ")
	clordid := neworder2[0]
	clordid = clordid[1:len(clordid)]
	price := neworder2[1]
	symbol := neworder2[2]
	orderqty := neworder2[3]
	side := neworder2[4]
	ordtype := neworder2[5]
	timeinforce := neworder2[6]
	//senderid := neworder2[7]
	//targetid := neworder2[8]
	//targetsubid := neworder2[9]
	//targetsubid = targetsubid[0:len(targetsubid)-1]

	order := fix50nos.New(queryClOrdID(clordid), querySide(side), field.NewTransactTime(time.Now()), queryOrdType(&ordType, ordtype))
	order.SetHandlInst("1")
	order.Set(querySymbol(symbol))
	order.Set(queryOrderQty(orderqty))
	order.Set(queryTimeInForce(timeinforce))

	switch ordType.Value() {
	case enum.OrdType_LIMIT, enum.OrdType_STOP_LIMIT:
		order.Set(queryPrice(price))
	}

	//switch ordType.Value() {
	//case enum.OrdType_STOP, enum.OrdType_STOP_LIMIT:
	//	order.Set(queryStopPx())
	//}

	msg = order.ToMessage()
	queryHeader(&msg.Header)

	return
}

func QueryEnterOrder(neworder string) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
		}
	}()

	order := queryNewOrderSingle50(neworder)
	return quickfix.Send(order)
}